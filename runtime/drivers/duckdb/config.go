package duckdb

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/c2h5oh/datasize"
	"github.com/mitchellh/mapstructure"
)

// config represents the DuckDB driver config
type config struct {
	// DSN is the connection string
	DSN string `mapstructure:"dsn"`
	// PoolSize is the number of concurrent connections and queries allowed
	PoolSize int `mapstructure:"pool_size"`
	// Slots controls the amount of resources to allocate to the database
	Slots int `mapstructure:"slots"`
	// MemoryGBPerSlot is the amount of memory in bytes per slot
	MemoryGBPerSlot int `mapstructure:"memory_per_slot"`
	// CPUPerSlot is the number of CPU cores per slot
	CPUPerSlot int `mapstructure:"cpu_per_slot"`
	// StorageGBPerSlot is the amount of storage in GB per slot
	StorageGBPerSlot int `mapstructure:"storage_per_slot"`
	// AllowHostAccess denotes whether to limit access to the local environment and file system
	AllowHostAccess bool `mapstructure:"allow_host_access"`
	// ErrorOnIncompatibleVersion controls whether to return error or delete DBFile created with older duckdb version.
	ErrorOnIncompatibleVersion bool `mapstructure:"error_on_incompatible_version"`
	// ExtTableStorage controls if every table is stored in a different db file
	ExtTableStorage bool `mapstructure:"external_table_storage"`
	// StorageLimitBytes is the maximum size of all database files
	StorageLimitBytes int64 `mapstructure:"_"`
	// DBFilePath is the path where the database is stored. It is inferred from the DSN (can't be provided by user).
	DBFilePath string `mapstructure:"-"`
	// ExtStoragePath is the path where the database files are stored in case external_table_storage is true. It is inferred from the DSN (can't be provided by user).
	ExtStoragePath string `mapstructure:"-"`
}

func newConfig(cfgMap map[string]any) (*config, error) {
	cfg := &config{
		PoolSize: 1, // Default value
	}
	err := mapstructure.WeakDecode(cfgMap, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not decode config: %w", err)
	}

	// Parse DSN as URL
	uri, err := url.Parse(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("could not parse dsn: %w", err)
	}
	qry, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("could not parse dsn: %w", err)
	}

	if cfg.Slots != 0 {
		// memory limits
		qry.Add("max_memory", fmt.Sprintf("%dGB", cfg.Slots*cfg.MemoryGBPerSlot))

		// cpu limits
		qry.Add("threads", strconv.Itoa(cfg.Slots*cfg.CPUPerSlot))
		cfg.PoolSize = cfg.Slots * cfg.CPUPerSlot

		// storage limits
		cfg.StorageLimitBytes = int64(cfg.Slots*cfg.StorageGBPerSlot) * int64(datasize.GB)

		// Remove from query string (so not passed into DuckDB config)
		qry.Del("slots")

		// Rebuild DuckDB DSN (which should be "path?key=val&...")
		uri.RawQuery = qry.Encode()
		cfg.DSN = uri.String()
	}

	// Infer DBFilePath
	cfg.DBFilePath = uri.Path
	if cfg.ExtTableStorage {
		cfg.ExtStoragePath = filepath.Dir(cfg.DBFilePath)
	}

	// We also support overriding the pool size via the DSN by setting "rill_pool_size" as a query argument.
	if qry.Has("rill_pool_size") {
		// Parse as integer
		cfg.PoolSize, err = strconv.Atoi(qry.Get("rill_pool_size"))
		if err != nil {
			return nil, fmt.Errorf("could not parse dsn: 'rill_pool_size' is not an integer")
		}

		// Remove from query string (so not passed into DuckDB config)
		qry.Del("rill_pool_size")

		// Rebuild DuckDB DSN (which should be "path?key=val&...")
		uri.RawQuery = qry.Encode()
		cfg.DSN = uri.String()
	}

	// Check pool size
	if cfg.PoolSize < 1 {
		return nil, fmt.Errorf("duckdb pool size must be >= 1")
	}

	return cfg, nil
}
