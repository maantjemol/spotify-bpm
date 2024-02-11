<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import Input from "$lib/components/ui/input/input.svelte";
	import { Separator } from "$lib/components/ui/separator";
  import { goto } from '$app/navigation';

  let spotifyUrl = "";

  // https://open.spotify.com/playlist/2OjFH3O7GqtU8F9eb2iqEr?si=e3469439eb7145ea
  // 2OjFH3O7GqtU8F9eb2iqEr
  const getIdFromUrl = (url: string) => {
    try {
      return new URL(url).pathname.split("/")[2] ?? "";
    } catch (e) {
      return "";
    }
  }

  $: spotifyId = spotifyUrl ? getIdFromUrl(spotifyUrl) : "";

</script>
<section class="flex items-center justify-center w-screen h-screen flex-col px-[10%]">
  <h1 class="font-medium text-4xl">Playlist zoeken</h1>
  <Separator class="w-[100px] my-4"/>
  <div class="w-[500px] flex gap-4">
    <Input placeholder="playlist url" type="search" bind:value={spotifyUrl} />
    <Button disabled={!spotifyId} on:click={() => goto(`/playlist/${spotifyId}`)}>Zoeken</Button>
  </div>
</section>