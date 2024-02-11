<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Playlist } from '$lib/spotify';
  import * as Table from "$lib/components/ui/table";
	import { Separator } from '$lib/components/ui/separator';
	export let data: PageData;

  let loading = true;
  let playlistData: Playlist;

  let numToKey = ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"]

  onMount(async () => {
    const res = await fetch(`/api/getbpm?id=${data.playlistId}`);
    const playdata = await res.json() as Playlist
    playlistData = playdata
    loading = false;
  });

</script>
{#if loading}
<div class="items-center justify-center flex w-screen h-screen">
  <p>Loading...</p>
</div>
{:else}
<!-- {#each playlistData.tracks as song}
  <p>{song.name}</p>
  <p>{song.artistName}</p>
{/each} -->
<section class="mx-[10%] mt-20">
  <h1 class="font-medium text-4xl mb-4">Tempos and keys for playlist:</h1>
  <Separator />
  <div class="mt-4">
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.Head>Name</Table.Head>
          <Table.Head>Artist</Table.Head>
          <Table.Head>Key</Table.Head>
          <Table.Head>Time signature</Table.Head>
          <Table.Head class="text-right">BPM</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each playlistData.tracks as song, i (i)}
          <Table.Row>
            <Table.Cell class="font-medium">
              <a href={song.url} class="underline" target="_blank">{song.name}</a>
            </Table.Cell>
            <Table.Cell>{song.artistName}</Table.Cell>
            <Table.Cell>{numToKey[song.key]}</Table.Cell>
            <Table.Cell>{song.timeSignature}/4</Table.Cell>
            <Table.Cell class="text-right">{song.BPM}</Table.Cell>
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>
</section>
{/if}
