<script lang="ts">
  import ArrowDown from "@rilldata/web-common/components/icons/ArrowDown.svelte";
  import { getStateManagers } from "@rilldata/web-common/features/dashboards/state-managers/state-managers";
  import { metricsExplorerStore } from "@rilldata/web-common/features/dashboards/stores/dashboard-stores";
  import {
    ExpandedState,
    SortingState,
    TableOptions,
    Updater,
    flexRender,
    createSvelteTable,
    getCoreRowModel,
    getExpandedRowModel,
  } from "@tanstack/svelte-table";
  import {
    createVirtualizer,
    defaultRangeExtractor,
  } from "@tanstack/svelte-virtual";
  import { isTimeDimension } from "./pivot-utils";
  import type { Readable } from "svelte/motion";
  import { derived } from "svelte/store";
  import { getPivotConfig } from "./pivot-data-store";
  import type { PivotDataRow, PivotDataStore } from "./types";
  import Resizer from "@rilldata/web-common/layout/Resizer.svelte";
  import { extractSamples } from "@rilldata/web-common/components/virtualized-table/init-widths";
  import { clamp } from "@rilldata/web-common/lib/clamp";

  // Distance threshold (in pixels) for triggering data fetch
  const ROW_THRESHOLD = 200;
  const OVERSCAN = 60;
  const ROW_HEIGHT = 24;
  const HEADER_HEIGHT = 30;
  const MEASURE_PADDING = 16;
  const MIN_COL_WIDTH = 150;
  const MAX_COL_WIDTH = 600;
  const MAX_INIT_COL_WIDTH = 400;
  const MIN_MEASURE_WIDTH = 60;
  const MAX_MEAUSRE_WIDTH = 300;
  const INIT_MEASURE_WIDTH = 90;

  export let pivotDataStore: PivotDataStore;

  const stateManagers = getStateManagers();

  const { dashboardStore, metricsViewName } = stateManagers;

  const config = getPivotConfig(stateManagers);

  const pivotDashboardStore = derived(dashboardStore, (dashboard) => {
    return dashboard?.pivot;
  });

  const options: Readable<TableOptions<PivotDataRow>> = derived(
    [pivotDashboardStore, pivotDataStore],
    ([pivotConfig, pivotData]) => {
      let tableData = pivotData.data;
      if (pivotData.totalsRowData) {
        tableData = [pivotData.totalsRowData, ...pivotData.data];
      }
      return {
        data: tableData,
        columns: pivotData.columnDef,
        state: {
          expanded: pivotConfig.expanded,
          sorting: pivotConfig.sorting,
        },
        onExpandedChange,
        getSubRows: (row) => row.subRows,
        onSortingChange,
        getExpandedRowModel: getExpandedRowModel(),
        getCoreRowModel: getCoreRowModel(),
        enableSortingRemoval: false,
        enableExpanding: true,
      };
    },
  );

  const table = createSvelteTable(options);

  let containerRefElement: HTMLDivElement;
  let stickyRows = [0];
  let rowScrollOffset = 0;
  let scrollLeft = 0;
  let initialMeasureIndexOnResize = 0;
  let initLengthOnResize = 0;
  let initScrollOnResize = 0;
  let percentOfChangeDuringResize = 0;
  let resizingMeasure = false;

  $: ({
    expanded,
    sorting,
    rowPage,
    columns,
    rows: rowPills,
  } = $pivotDashboardStore);

  $: timeDimension = $config.time.timeDimension;
  $: hasDimension = rowPills.dimension.length > 0;
  $: reachedEndForRows = !!$pivotDataStore?.reachedEndForRowData;
  $: assembled = $pivotDataStore.assembled;

  $: measureNames = columns.measure.map((measure) => measure.title) ?? [];
  $: measureCount = measureNames.length;
  $: measureLengths = measureNames.map((name) =>
    Math.max(INIT_MEASURE_WIDTH, name.length * 7 + MEASURE_PADDING),
  );
  $: measureGroups = headerGroups[headerGroups.length - 2]?.headers?.slice(
    hasDimension ? 1 : 0,
  ) ?? [null];
  $: measureGroupsLength = measureGroups.length;
  $: totalMeasureWidth = measureLengths.reduce((acc, val) => acc + val, 0);
  $: totalLength = measureGroupsLength * totalMeasureWidth;

  $: headerGroups = $table.getHeaderGroups();
  $: totalHeaderHeight = headerGroups.length * HEADER_HEIGHT;
  $: headers = headerGroups[0].headers;
  $: firstColumnName = hasDimension
    ? String(headers[0]?.column.columnDef.header)
    : null;
  $: firstColumnWidth =
    hasDimension && firstColumnName
      ? calculateFirstColumnWidth(firstColumnName)
      : 0;

  $: rows = $table.getRowModel().rows;
  $: virtualizer = createVirtualizer<HTMLDivElement, HTMLTableRowElement>({
    count: rows.length,
    getScrollElement: () => containerRefElement,
    estimateSize: () => ROW_HEIGHT,
    overscan: OVERSCAN,
    initialOffset: rowScrollOffset,
    rangeExtractor: (range) => {
      const next = new Set([...stickyRows, ...defaultRangeExtractor(range)]);

      return [...next].sort((a, b) => a - b);
    },
  });

  $: virtualRows = $virtualizer.getVirtualItems();
  $: totalRowSize = $virtualizer.getTotalSize();

  $: rowScrollOffset = $virtualizer?.scrollOffset || 0;

  // In this virtualization model, we create buffer rows before and after our real data
  // This maintains the "correct" scroll position when the user scrolls
  $: [before, after] = virtualRows.length
    ? [
        (virtualRows[1]?.start ?? virtualRows[0].start) - ROW_HEIGHT,
        totalRowSize - virtualRows[virtualRows.length - 1].end,
      ]
    : [0, 0];

  $: if (resizingMeasure && containerRefElement && measureLengths) {
    containerRefElement.scrollTo({
      left:
        initScrollOnResize +
        percentOfChangeDuringResize * (totalLength - initLengthOnResize),
    });
  }

  function onExpandedChange(updater: Updater<ExpandedState>) {
    // Something is off with tanstack's types
    //@ts-expect-error-free
    //eslint-disable-next-line
    expanded = updater(expanded);
    metricsExplorerStore.setPivotExpanded($metricsViewName, expanded);
  }

  function onSortingChange(updater: Updater<SortingState>) {
    if (updater instanceof Function) {
      sorting = updater(sorting);
    } else {
      sorting = updater;
    }
    metricsExplorerStore.setPivotSort($metricsViewName, sorting);
  }

  const handleScroll = (containerRefElement?: HTMLDivElement | null) => {
    if (containerRefElement) {
      const { scrollHeight, scrollTop, clientHeight } = containerRefElement;
      const bottomEndDistance = scrollHeight - scrollTop - clientHeight;
      scrollLeft = containerRefElement.scrollLeft;

      // Fetch more data when scrolling near the bottom end
      if (
        bottomEndDistance < ROW_THRESHOLD &&
        !$pivotDataStore.isFetching &&
        !reachedEndForRows
      ) {
        metricsExplorerStore.setPivotRowPage($metricsViewName, rowPage + 1);
      }
    }
  };

  function calculateFirstColumnWidth(firstColumnName: string) {
    const rows = $pivotDataStore.data;

    // Dates are displayed as shorter values
    if (isTimeDimension(firstColumnName, timeDimension)) return MIN_COL_WIDTH;

    const samples = extractSamples(
      rows.map((row) => row[firstColumnName]),
    ).filter((v): v is string => typeof v === "string");

    const maxValueLength = samples.reduce((max, value) => {
      return Math.max(max, value.length);
    }, 0);

    const finalBasis = Math.max(firstColumnName.length, maxValueLength);
    const pixelLength = finalBasis * 7;
    const final = clamp(MIN_COL_WIDTH, pixelLength + 16, MAX_INIT_COL_WIDTH);

    return final;
  }

  function onResizeStart(e: MouseEvent) {
    initLengthOnResize = totalLength;
    initScrollOnResize = scrollLeft;

    const offset =
      e.clientX -
      containerRefElement.getBoundingClientRect().left -
      firstColumnWidth -
      measureLengths.reduce((rollingSum, length, i) => {
        return i <= initialMeasureIndexOnResize
          ? rollingSum + length
          : rollingSum;
      }, 0) +
      4;

    percentOfChangeDuringResize = (scrollLeft + offset) / totalLength;
  }
</script>

<div
  class="table-wrapper"
  class:with-row-dimension={hasDimension}
  style:--row-height="{ROW_HEIGHT}px"
  style:--header-height="{HEADER_HEIGHT}px"
  style:--total-header-height="{totalHeaderHeight + headerGroups.length}px"
  bind:this={containerRefElement}
  on:scroll={() => handleScroll(containerRefElement)}
>
  <table style:width="{totalLength}px">
    {#if firstColumnName && firstColumnWidth}
      <colgroup>
        <col
          style:width="{firstColumnWidth}px"
          style:max-width="{firstColumnWidth}px"
        />
      </colgroup>
    {/if}

    {#each measureGroups as _}
      <colgroup>
        {#each measureLengths as length}
          <col style:width="{length}px" style:max-width="{length}px" />
        {/each}
      </colgroup>
    {/each}

    <thead>
      {#each headerGroups as headerGroup, group (headerGroup.id)}
        <tr>
          {#each headerGroup.headers as header, i (header.id)}
            {@const sortDirection = header.column.getIsSorted()}
            {@const isFirstColumn = i === 0}
            {@const canResize = hasDimension && (isFirstColumn || group !== 0)}
            {@const measureIndex = (i - 1) % measureLengths.length}
            <th colSpan={header.colSpan}>
              {#if canResize}
                <Resizer
                  side="right"
                  direction="EW"
                  min={isFirstColumn ? MIN_COL_WIDTH : MIN_MEASURE_WIDTH}
                  max={isFirstColumn ? MAX_COL_WIDTH : MAX_MEAUSRE_WIDTH}
                  basis={isFirstColumn ? MIN_COL_WIDTH : INIT_MEASURE_WIDTH}
                  dimension={isFirstColumn
                    ? firstColumnWidth
                    : measureLengths[measureIndex]}
                  onMouseDown={(e) => {
                    resizingMeasure = !isFirstColumn;
                    initialMeasureIndexOnResize = measureIndex;
                    if (resizingMeasure) onResizeStart(e);
                  }}
                  onUpdate={(d) => {
                    if (isFirstColumn) {
                      firstColumnWidth = d;
                    } else {
                      measureLengths[measureIndex] = d;
                    }
                  }}
                />
              {/if}

              <button
                class="header-cell"
                class:cursor-pointer={header.column.getCanSort()}
                class:select-none={header.column.getCanSort()}
                on:click={header.column.getToggleSortingHandler()}
              >
                {#if !header.isPlaceholder}
                  <p class="truncate">
                    {header.column.columnDef.header}
                  </p>
                  {#if sortDirection}
                    <span
                      class="transition-transform -mr-1"
                      class:-rotate-180={sortDirection === "asc"}
                    >
                      <ArrowDown />
                    </span>
                  {/if}
                {/if}
              </button>
            </th>
          {/each}
        </tr>
      {/each}
    </thead>
    <tbody>
      <tr style:height="{before}px" />
      {#each virtualRows as row (row.index)}
        {@const cells = rows[row.index].getVisibleCells()}
        <tr>
          {#each cells as cell, i (cell.id)}
            {@const result =
              typeof cell.column.columnDef.cell === "function"
                ? cell.column.columnDef.cell(cell.getContext())
                : cell.column.columnDef.cell}
            <td
              class="ui-copy-number"
              class:border-r={i % measureCount === 0 && i}
            >
              <div class="cell">
                {#if result?.component && result?.props}
                  <svelte:component
                    this={result.component}
                    {...result.props}
                    {assembled}
                  />
                {:else if typeof result === "string" || typeof result === "number"}
                  {result}
                {:else}
                  <svelte:component
                    this={flexRender(
                      cell.column.columnDef.cell,
                      cell.getContext(),
                    )}
                  />
                {/if}
              </div>
            </td>
          {/each}
        </tr>
      {/each}
      <tr style:height="{after}px" />
    </tbody>
  </table>
</div>

<style lang="postcss">
  * {
    @apply border-slate-200;
  }

  table {
    @apply p-0 m-0 border-spacing-0 border-separate w-fit;
    @apply font-normal select-none;
    @apply bg-white table-fixed;
  }

  .table-wrapper {
    @apply overflow-auto h-fit max-h-full w-fit max-w-full;
    @apply border rounded-md z-40;
  }

  /* Pin header */
  thead {
    @apply sticky top-0;
    @apply z-30 bg-white;
  }

  tbody .cell {
    height: var(--row-height);
  }

  th {
    @apply p-0 m-0 text-xs;
    @apply border-r border-b relative;
  }

  th:last-of-type,
  td:last-of-type {
    @apply border-r-0;
  }

  th,
  td {
    @apply whitespace-nowrap text-xs;
  }

  td {
    @apply text-right;
    @apply p-0 m-0;
  }

  .header-cell {
    @apply px-2 bg-white size-full;
    @apply flex items-center gap-x-1 w-full truncate;
    height: var(--header-height);
  }

  .cell {
    @apply size-full p-1 px-2;
  }

  /* The leftmost header cells have no bottom border unless they're the last row */
  .with-row-dimension thead > tr:not(:last-of-type) > th:first-of-type {
    @apply border-b-0;
  }

  .with-row-dimension tr > th:first-of-type {
    @apply sticky left-0 z-20;
    @apply bg-white;
  }

  .with-row-dimension tr > td:first-of-type {
    @apply sticky left-0 z-10;
    @apply bg-white;
  }

  tr > td:first-of-type:not(:last-of-type) {
    @apply border-r font-medium;
  }

  /* The totals row */
  tbody > tr:nth-of-type(2) {
    @apply bg-slate-100 sticky z-20 font-semibold;
    top: var(--total-header-height);
  }

  /* The totals row header */
  tbody > tr:nth-of-type(2) > td:first-of-type {
    @apply font-semibold;
  }

  tr:hover,
  tr:hover .cell {
    @apply bg-slate-100;
  }
</style>
