<script lang="ts">
  import type { Extension } from "@codemirror/state";
  import { EditorState } from "@codemirror/state";
  import { EditorView } from "@codemirror/view";
  import { createEventDispatcher, onMount } from "svelte";
  import { bindEditorEventsToDispatcher } from "./dispatch-events";
  import { base } from "./presets/base";
  import { yaml } from "./presets/yaml";

  const dispatch = createEventDispatcher();

  export let content: string;
  export let extensions: Extension[] = [];
  export let view: EditorView | undefined = undefined;
  export let whenFocused = false;
  export let key: string;

  let container: HTMLElement;

  onMount(() => {
    view = new EditorView({
      state: EditorState.create({
        doc: content,
        extensions: [
          // any extensions passed as props
          ...extensions,
          // establish a basic editor
          base(),
          // establish the yaml editor, which currently only has
          // syntax highlighting
          yaml(),
          // this will catch certain events and dispatch them to the parent
          bindEditorEventsToDispatcher(dispatch, whenFocused),
        ],
      }),
      parent: container,
    });
  });

  $: if (key) {
    // When the key changes, unfocus the Editor so that the update is dispatched
    view?.contentDOM.blur();
  }

  // reactive statements to dynamically update the editor when inputs change
  $: updateEditorContents(content);

  function updateEditorContents(newContent: string) {
    if (view && !view.hasFocus) {
      let curContent = view.state.doc.toString();
      if (newContent != curContent) {
        view.dispatch({
          changes: {
            from: 0,
            to: curContent.length,
            insert: newContent,
          },
        });
      }
    }
  }
</script>

<div bind:this={container} class="contents" />
