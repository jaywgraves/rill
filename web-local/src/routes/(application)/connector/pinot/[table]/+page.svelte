<script lang="ts">
  import { page } from "$app/stores";
  import { featureFlags } from "@rilldata/web-common/features/feature-flags";
  import TablePreviewWorkspace from "@rilldata/web-common/features/tables/TablePreviewWorkspace.svelte";
  import { error } from "@sveltejs/kit";
  import { onMount } from "svelte";

  const { readOnly } = featureFlags;

  $: table = $page.params.table;

  onMount(() => {
    if ($readOnly) {
      throw error(404, "Page not found");
    }
  });
</script>

<svelte:head>
  <title>Rill Developer | {table}</title>
</svelte:head>

<TablePreviewWorkspace connector="pinot" {table} />
