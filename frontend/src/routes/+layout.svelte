<script>
  import "../app.css";
  import { Avatar } from "$lib/components/ui/avatar";
	import AvatarFallback from "$lib/components/ui/avatar/avatar-fallback.svelte";
  import { page } from "$app/stores";
	import { Button } from "$lib/components/ui/button";
  export let data;
  const { session } = data;
</script>

<nav class={`sticky top-0 z-10 flex flex-row justify-between bg-white/40 backdrop-blur py-4 ${$page.url.pathname == "/" ? "my-8" : ""} px-24 items-center`}>
  <a href="/" class="font-robomono font-medium">
    0rdina1
  </a>
  <div class={`flex flex-row ${$page.url.pathname.includes("topics") ? "text-yellow-700" : "text-gray-400"} hover:text-yellow-600 transition-all duration-300 ease-in-out`}>
    <a href="/topics">Topics</a>
  </div>
  {#if session}
    <a href="/profile">
      <Avatar>
        <AvatarFallback>A</AvatarFallback>
      </Avatar>
    </a>
  {:else}
    <a href="/profile/signup">
      <Button>Sign up</Button>
    </a>   
  {/if}
</nav>
<slot />