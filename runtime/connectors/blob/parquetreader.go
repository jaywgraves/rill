package blob

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v11/arrow"
	"github.com/apache/arrow/go/v11/arrow/array"
	"github.com/apache/arrow/go/v11/arrow/memory"
	"github.com/apache/arrow/go/v11/parquet"
	"github.com/apache/arrow/go/v11/parquet/compress"
	"github.com/apache/arrow/go/v11/parquet/file"
	"github.com/apache/arrow/go/v11/parquet/pqarrow"
	runtimev1 "github.com/rilldata/rill/proto/gen/rill/runtime/v1"
	"github.com/rilldata/rill/runtime/pkg/container"
	"gocloud.dev/blob"
)

// number of rows of a column fetched in one call
// keeping it high seems to improve latency at the cost of accuracy in size of fetched data as per policy
const batchSize = int64(1000)

type extractConfig struct {
	limtInBytes uint64
	strategy    runtimev1.Source_ExtractPolicy_Strategy
}

func downloadParquet(ctx context.Context, bucket *blob.Bucket, obj *blob.ListObject, option extractConfig, fw *os.File) error {
	mem := memory.NewCheckedAllocator(memory.DefaultAllocator)
	reader := NewBlobObjectReader(ctx, bucket, obj)

	props := parquet.NewReaderProperties(mem)
	props.BufferedStreamEnabled = true

	pf, err := file.NewParquetReader(reader, file.WithReadProps(props))
	if err != nil {
		return err
	}
	defer pf.Close()

	arrowReadProperties := pqarrow.ArrowReadProperties{BatchSize: batchSize, Parallel: true}
	// reader to convert parquet objects to arrow objects
	fileReader, err := pqarrow.NewFileReader(pf, arrowReadProperties, mem)
	if err != nil {
		return err
	}

	numRowGroups := pf.NumRowGroups()
	if numRowGroups == 0 {
		return fmt.Errorf("invalid parquet")
	}

	records, err := estimateRecords(ctx, pf, fileReader, option)
	if err != nil {
		return err
	}
	defer func() {
		for _, rec := range records {
			rec.Release()
		}
	}()

	schema, err := fileReader.Schema()
	if err != nil {
		return err
	}

	table := array.NewTableFromRecords(schema, records)
	defer table.Release()

	// duck db requires root Repetitions to be required
	// keeping compressions as uncompressed since file will be immediately consumed and deleted
	wp := parquet.NewWriterProperties(
		parquet.WithRootRepetition(parquet.Repetitions.Required),
		parquet.WithCompression(compress.Codecs.Uncompressed),
		parquet.WithAllocator(mem))
	return pqarrow.WriteTable(
		table,
		fw,
		table.NumRows(),
		wp,
		pqarrow.NewArrowWriterProperties(pqarrow.WithAllocator(mem)),
	)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rangeInt(start, end int, rev bool) []int {
	result := make([]int, end-start)
	for i := start; i < end; i++ {
		result[i] = i
	}

	if rev {
		reverse(result)
	}
	return result
}

func containerForRecordLimiting(option *extractConfig) (container.Container[arrow.Record], error) {
	switch option.strategy {
	case runtimev1.Source_ExtractPolicy_STRATEGY_TAIL:
		return container.NewTailContainer(int(option.limtInBytes), func(rec arrow.Record) { rec.Release() })
	case runtimev1.Source_ExtractPolicy_STRATEGY_HEAD:
		return container.NewBoundedContainer[arrow.Record](int(option.limtInBytes))
	default:
		// No option selected - this should not be used for partial downloads though
		return container.NewUnboundedContainer[arrow.Record]()
	}
}

// estimateRecords estimates the number of rows to fetch based on extract policy
// each arrow.Record will hold batchSize number of rows
func estimateRecords(ctx context.Context, reader *file.Reader, pqToArrowReader *pqarrow.FileReader, config extractConfig) ([]arrow.Record, error) {
	rowIndexes := rangeInt(0, reader.NumRowGroups(), config.strategy == runtimev1.Source_ExtractPolicy_STRATEGY_TAIL)

	// row group indices that we need
	reqRowIndices := make([]int, 0)
	var cumSize uint64
	var rows int64
	for _, index := range rowIndexes {
		reqRowIndices = append(reqRowIndices, index)
		rowGroup := reader.RowGroup(index)
		rowGroupSize := rowGroup.ByteSize()
		rowCount := rowGroup.NumRows()

		if cumSize+uint64(rowGroupSize) > config.limtInBytes {
			// taking entire rowgroup crosses allowed size
			perRowSize := uint64(rowGroupSize / rowCount)
			rows += int64((config.limtInBytes - cumSize) / perRowSize)
			break
		}
		rows += rowCount
	}

	r, err := pqToArrowReader.GetRecordReader(ctx, rangeInt(0, reader.RowGroup(0).NumColumns(), false), reqRowIndices)
	if err != nil {
		return nil, err
	}
	defer r.Release()

	// one record has batchsize rows
	numRecords := rows / batchSize
	if numRecords == 0 {
		// if parquet file has less than batchSize rows or user selects less than batchSize rows
		numRecords = 1
	}

	c, err := containerForRecordLimiting(&extractConfig{strategy: config.strategy, limtInBytes: uint64(numRecords)})
	if err != nil {
		return nil, err
	}

	for r.Next() && !c.IsFull() {
		rec := r.Record()
		rec.Retain()
		c.Add(rec)
	}
	return c.Items(), nil
}
