<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Playlist } from '$lib/spotify';
  import * as Table from "$lib/components/ui/table";
	import * as Card from '$lib/components/ui/card';
	import * as Select from '$lib/components/ui/select';
	import { Separator } from '$lib/components/ui/separator';
	import CardFooter from '$lib/components/ui/card/card-footer.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Switch } from '$lib/components/ui/switch';
  import { Label } from "$lib/components/ui/label";
	import { tableview } from '$lib/stores';
	import { AspectRatio } from '$lib/components/ui/aspect-ratio';
	export let data: PageData;
  import { fade } from 'svelte/transition';


  let loading = true;
  let playlistData: Playlist;

  let numToKey = ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"]
  let selectedKey:{value:number, label:string} | undefined = undefined

  function closestDistanceKeys(start:number, target:number) {
    let distance = target - start;
    if (distance > 6) {
      distance -= 12;
    } else if (distance < -6) {
      distance += 12;
    }
    return distance;
}

  onMount(async () => {
    const res = await fetch(`/api/getbpm?id=${data.playlistId}`);
    const playdata = await res.json() as Playlist
    playlistData = playdata
    loading = false;
  });

</script>
{#if loading}
<div class="items-center justify-center flex w-screen h-screen" transition:fade>
  <svg
  xmlns="http://www.w3.org/2000/svg"
  width={30}
  height={30}
  viewBox="0 0 24 24"
  fill="none"
  stroke="currentColor"
  stroke-width="20"
  stroke-linecap="round"
  stroke-linejoin="round"
  class="animate-spin"
>
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
</div>
{:else}
<!-- {#each playlistData.tracks as song}
  <p>{song.name}</p>
  <p>{song.artistName}</p>
{/each} -->
<section class="sm:mx-[10%] mx-[5%] mt-20 mb-20">
  <h1 class="font-medium lg:text-4xl md:text-3xl mb-4 text-2xl">Tempos en toonsoort voor playlist:</h1>
  <div class="flex gap-4 pb-4">
    <div class="sm:flex items-center space-x-2 hidden">
      <Switch bind:checked={$tableview} id="airplane-mode" />
      <Label for="airplane-mode">Gebruik tabel</Label>
    </div>
    <Select.Root bind:selected={selectedKey}>
      <Select.Trigger class="w-[180px]">
        <Select.Value placeholder="Toonsoort" />
      </Select.Trigger>
      <Select.Content>
        {#each numToKey as key, i}  
          <Select.Item value={i}>{key}</Select.Item>
        {/each}
      </Select.Content>
    </Select.Root>
  </div>
  <Separator />
  {#if $tableview}
  <div class="mt-4">
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.Head></Table.Head>
          <Table.Head>Naam</Table.Head>
          <Table.Head>Artiest</Table.Head>
          <!-- <Table.Head>Audio</Table.Head> -->
          <Table.Head>Toonsoort</Table.Head>
          <Table.Head>Time signature</Table.Head>
          <Table.Head class="text-right">BPM</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each playlistData.tracks as song, i (i)}
          <Table.Row>
            <Table.Cell><div class="h-10 w-10"><img src={song.image} alt="spotify cover" class="rounded-sm object-cover" /></div></Table.Cell>
            <Table.Cell class="font-medium">
              <a href={song.url} class="underline" target="_blank">{song.name}</a>
            </Table.Cell>
            <Table.Cell>{song.artistName}</Table.Cell>
            <!-- <Table.Cell><audio controls src={song.previewURL}/></Table.Cell> -->
            <Table.Cell>{#if selectedKey}
              <p class="text-gray-400"><span class="font-medium text-black dark:text-white">{numToKey[song.key]}</span> → {numToKey[selectedKey.value]} ({closestDistanceKeys(selectedKey.value, song.key)})</p>
            {:else}
              <p>{numToKey[song.key]}</p>
            {/if}</Table.Cell>
            <Table.Cell>{song.timeSignature}/4</Table.Cell>
            <Table.Cell class="text-right">{song.BPM}</Table.Cell>
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>
  {:else}
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mt-4">
    {#each playlistData.tracks as song, i (i)}
    <Card.Root>
      <Card.Header>
        <a class="flex gap-4 items-center" href={song.url}>
          <img src={song.image} alt="spotify cover" class="rounded-md object-cover h-12" />
          <div>
            <Card.Title>{song.name}</Card.Title>
            <Card.Description>{song.artistName}</Card.Description>
          </div>
        </a>
      </Card.Header>
      <Card.Content>
        <p class="font-light text-gray-400"><span class="font-medium text-black dark:text-white">{song.BPM}</span> BPM</p>
        <p class="font-light text-gray-400" ><span class="font-medium text-black dark:text-white">{song.timeSignature} / 4</span> Maatsoort</p>
        {#if selectedKey}
          <p class="font-light text-gray-400"><span class="font-medium text-black dark:text-white">{numToKey[song.key]}</span> → {numToKey[selectedKey.value]} ({closestDistanceKeys(selectedKey.value, song.key)})</p>
        {:else}
          <p class="font-light text-gray-400"><span class="font-medium text-black dark:text-white">{numToKey[song.key]}</span> Toonsoort</p>
        {/if}
      </Card.Content>
      <!-- <Card.CardFooter>
        <Button href={song.url} target="_blank">Open spotify</Button>
      </Card.CardFooter> -->
    </Card.Root>
    {/each}
  </div>
  {/if}
</section>
{/if}
