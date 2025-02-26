<script lang="ts">
  import { afterNavigate, beforeNavigate, goto } from "$app/navigation";
  import { page } from "$app/stores";
  import WorkspaceError from "@rilldata/web-common/components/WorkspaceError.svelte";
  import Editor from "@rilldata/web-common/features/editor/Editor.svelte";
  import FileWorkspaceHeader from "@rilldata/web-common/features/editor/FileWorkspaceHeader.svelte";
  import { FILES_WITHOUT_AUTOSAVE } from "@rilldata/web-common/features/editor/config";
  import { getExtensionsForFile } from "@rilldata/web-common/features/editor/getExtensionsForFile";
  import { addLeadingSlash } from "@rilldata/web-common/features/entity-management/entity-mappers";
  import { fileArtifacts } from "@rilldata/web-common/features/entity-management/file-artifacts";
  import { ResourceKind } from "@rilldata/web-common/features/entity-management/resource-selectors";
  import { directoryState } from "@rilldata/web-common/features/file-explorer/directory-store";
  import UnsavedSourceDialog from "@rilldata/web-common/features/sources/editor/UnsavedSourceDialog.svelte";
  import { extractFileExtension } from "@rilldata/web-common/features/sources/extract-file-name";
  import WorkspaceContainer from "@rilldata/web-common/layout/workspace/WorkspaceContainer.svelte";
  import WorkspaceEditorContainer from "@rilldata/web-common/layout/workspace/WorkspaceEditorContainer.svelte";
  import { workspaces } from "@rilldata/web-common/layout/workspace/workspace-stores";
  import {
    createRuntimeServiceGetFile,
    createRuntimeServicePutFile,
  } from "@rilldata/web-common/runtime-client";
  import { runtime } from "@rilldata/web-common/runtime-client/runtime-store";
  import { onMount } from "svelte";
  import SourceModelPage from "../../[type=workspace]/[name]/+page.svelte";
  import ChartPage from "../../chart/[name]/+page.svelte";
  import CustomDashboardPage from "../../custom-dashboards/[name]/+page.svelte";
  import DashboardPage from "../../dashboard/[name]/edit/+page.svelte";

  const UNSUPPORTED_EXTENSIONS = [".parquet", ".db", ".db.wal"];

  let interceptedUrl: string | null = null;

  const putFile = createRuntimeServicePutFile(); // TODO: optimistically update the Get File cache

  $: filePath = addLeadingSlash($page.params.file);
  $: fileExtension = extractFileExtension(filePath);
  $: fileTypeUnsupported = UNSUPPORTED_EXTENSIONS.includes(fileExtension);

  $: fileQuery = createRuntimeServiceGetFile(
    $runtime.instanceId,
    { path: filePath },
    {
      query: {
        enabled: !fileTypeUnsupported,
      },
    },
  );
  $: fileError = !!$fileQuery.error;
  $: fileErrorMessage = $fileQuery.error?.response.data.message;
  $: fileArtifact = fileArtifacts.getFileArtifact(filePath);
  $: name = fileArtifact.name;
  $: resourceKind = $name?.kind;

  $: isSource = resourceKind === ResourceKind.Source;
  $: isModel = resourceKind === ResourceKind.Model;
  $: isDashboard = resourceKind === ResourceKind.MetricsView;
  $: isChart = resourceKind === ResourceKind.Component;
  $: isCustomDashboard = resourceKind === ResourceKind.Dashboard;
  $: isOther =
    !isSource && !isModel && !isDashboard && !isChart && !isCustomDashboard;

  onMount(() => {
    expandDirectory(filePath);

    // TODO: Focus on the code editor
  });

  afterNavigate(() => {
    expandDirectory(filePath);

    // TODO: Focus on the code editor
  });

  beforeNavigate((e) => {
    if (!hasUnsavedChanges || interceptedUrl) {
      localContent = null;
      return;
    }

    e.cancel();

    if (e.to) interceptedUrl = e.to.url.href;
  });

  let blob = "";
  $: blob = $fileQuery.data?.blob ?? blob;

  let localContent: string | null = null;
  $: hasUnsavedChanges = localContent !== null && localContent !== blob;

  $: pathname = $page.url.pathname;
  $: workspace = workspaces.get(pathname);
  $: autoSave = workspace.editor.autoSave;
  $: disableAutoSave = FILES_WITHOUT_AUTOSAVE.includes(filePath);

  async function save() {
    if (localContent === null) return;

    await $putFile.mutateAsync({
      instanceId: $runtime.instanceId,
      data: {
        path: filePath,
        blob: localContent,
      },
    });
  }

  function revert() {
    localContent = null;
  }

  // TODO: move this logic into the DirectoryState
  // TODO: expand all directories in the path, not just the last one
  function expandDirectory(filePath: string) {
    const directory = filePath.split("/").slice(0, -1).join("/");
    directoryState.expand(directory);
  }

  function handleConfirm() {
    if (!interceptedUrl) return;
    const url = interceptedUrl;
    localContent = null;
    hasUnsavedChanges = false;
    interceptedUrl = null;
    goto(url).catch(console.error);
  }

  function handleCancel() {
    interceptedUrl = null;
  }
</script>

{#if fileTypeUnsupported}
  <WorkspaceError message="Unsupported file type." />
{:else if fileError}
  <WorkspaceError message={`Error loading file: ${fileErrorMessage}`} />
{:else if isSource || isModel}
  <SourceModelPage data={{ fileArtifact }} />
{:else if isDashboard}
  <DashboardPage data={{ fileArtifact }} />
{:else if isChart}
  {#key $page.params.file}
    <ChartPage data={{ fileArtifact }} />
  {/key}
{:else if isCustomDashboard}
  <CustomDashboardPage data={{ fileArtifact }} />
{:else if isOther}
  <WorkspaceContainer inspector={false}>
    <FileWorkspaceHeader
      filePath={$page.params.file}
      {hasUnsavedChanges}
      slot="header"
    />
    <div
      slot="body"
      class="editor-pane size-full overflow-hidden flex flex-col"
    >
      <WorkspaceEditorContainer>
        <Editor
          key={filePath}
          remoteContent={blob}
          {hasUnsavedChanges}
          extensions={getExtensionsForFile(filePath)}
          {disableAutoSave}
          bind:localContent
          bind:autoSave={$autoSave}
          on:save={save}
          on:revert={revert}
        />
      </WorkspaceEditorContainer>
    </div>
  </WorkspaceContainer>
{/if}

{#if interceptedUrl}
  <UnsavedSourceDialog
    context="file"
    on:confirm={handleConfirm}
    on:cancel={handleCancel}
  />
{/if}
