<script lang="ts">
  import { goto } from "$app/navigation";
  import { createAdminServiceGetCurrentUser } from "@rilldata/web-admin/client";
  import { ADMIN_URL } from "@rilldata/web-admin/client/http-client";
  import CtaButton from "@rilldata/web-common/components/calls-to-action/CTAButton.svelte";
  import CtaContentContainer from "@rilldata/web-common/components/calls-to-action/CTAContentContainer.svelte";
  import CtaHeader from "@rilldata/web-common/components/calls-to-action/CTAHeader.svelte";
  import CtaLayoutContainer from "@rilldata/web-common/components/calls-to-action/CTALayoutContainer.svelte";
  import CtaMessage from "@rilldata/web-common/components/calls-to-action/CTAMessage.svelte";
  import Github from "@rilldata/web-common/components/icons/Github.svelte";
  import GithubRepoInline from "../../../../features/projects/GithubRepoInline.svelte";

  const urlParams = new URLSearchParams(window.location.search);
  const redirectURL = urlParams.get("redirect");
  const remote = new URL(decodeURIComponent(redirectURL)).searchParams.get(
    "remote",
  );
  const user = createAdminServiceGetCurrentUser({
    query: {
      onSuccess: (data) => {
        if (!data.user) {
          goto(`${ADMIN_URL}/auth/login?redirect=${window.location.href}`);
        }
      },
    },
  });

  function handleGoToGithub() {
    window.location.href = redirectURL;
  }
</script>

<svelte:head>
  <title>Connect to Github</title>
</svelte:head>

{#if $user.data && $user.data.user}
  <CtaLayoutContainer>
    <CtaContentContainer>
      <Github className="w-10 h-10 text-gray-900" />
      <CtaHeader>Connect to Github</CtaHeader>
      <CtaMessage>
        Rill projects deploy continuously when you push changes to Github.
      </CtaMessage>
      {#if remote}
        <CtaMessage>
          Please grant access to your repository<br /><GithubRepoInline
            githubUrl={remote}
          />
        </CtaMessage>
      {/if}
      <div class="mt-4 w-full">
        <CtaButton variant="primary" on:click={handleGoToGithub}
          >Connect to Github</CtaButton
        >
      </div>
    </CtaContentContainer>
  </CtaLayoutContainer>
{/if}
