<script lang="ts">
  /**
   * TimestampBound.svelte
   * ---------------------
   * This component will render the label bound on the TimestampDetail.svelte graph.
   * It also enables a shift + click to copy the bound as a query-ready timestamp.
   */
  import { eventBus } from "@rilldata/web-common/lib/event-bus/event-bus";
  import Shortcut from "@rilldata/web-common/components/tooltip/Shortcut.svelte";
  import StackingWord from "@rilldata/web-common/components/tooltip/StackingWord.svelte";
  import Tooltip from "@rilldata/web-common/components/tooltip/Tooltip.svelte";
  import TooltipContent from "@rilldata/web-common/components/tooltip/TooltipContent.svelte";
  import TooltipShortcutContainer from "@rilldata/web-common/components/tooltip/TooltipShortcutContainer.svelte";
  import TooltipTitle from "@rilldata/web-common/components/tooltip/TooltipTitle.svelte";
  import {
    createShiftClickAction,
    isClipboardApiSupported,
  } from "@rilldata/web-common/lib/actions/shift-click-action";
  import {
    datePortion,
    timePortion,
  } from "@rilldata/web-common/lib/formatters";
  import { removeLocalTimezoneOffset } from "@rilldata/web-common/lib/time/timezone";

  const { shiftClickAction } = createShiftClickAction();

  export let value: Date;
  export let label = "value";
  export let align: "left" | "right" = "left";
  let valueWithoutOffset: Date | undefined;
  $: if (value instanceof Date)
    valueWithoutOffset = removeLocalTimezoneOffset(value);
</script>

<Tooltip
  alignment={align == "left" ? "start" : "end"}
  distance={8}
  suppress={!isClipboardApiSupported()}
>
  <button
    class="text-{align} text-gray-500"
    style:line-height={1.1}
    use:shiftClickAction
    on:shift-click={async () => {
      if (valueWithoutOffset === undefined) return;
      const exportedValue = `TIMESTAMP '${valueWithoutOffset.toISOString()}'`;
      await navigator.clipboard.writeText(exportedValue);
      eventBus.emit("notification", {
        message: `copied ${exportedValue} to clipboard`,
      });
      // update this to set the active animation in the tooltip text
    }}
  >
    {#if valueWithoutOffset}
      <div>
        {datePortion(valueWithoutOffset)}
      </div>
      <div>
        {timePortion(valueWithoutOffset)}
      </div>
    {:else}
      loading...
    {/if}
  </button>
  <TooltipContent slot="tooltip-content">
    <TooltipTitle>
      <svelte:fragment slot="name"
        >{#if valueWithoutOffset === undefined}
          loading...
        {:else}
          {valueWithoutOffset.toISOString()}
        {/if}</svelte:fragment
      >
      <svelte:fragment slot="description">{label}</svelte:fragment>
    </TooltipTitle>
    <TooltipShortcutContainer>
      <div>
        <StackingWord key="shift">Copy</StackingWord> to clipboard
      </div>
      <Shortcut>
        <span
          style="
          font-family: var(--system); 
          font-size: 11.5px;
        ">⇧</span
        > + Click
      </Shortcut>
    </TooltipShortcutContainer>
  </TooltipContent>
</Tooltip>
