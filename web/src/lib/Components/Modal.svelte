<script lang="ts">
  import { createEventDispatcher } from "svelte";

  export let show = true
  export let title = `Start Pipeline`
  export let msg = `Are you sure you want to start`
  export let item = `pipelineActivity-pa-12345-67890-qwert-12qwe3`
  
  export let callbacks = [
    {
      label: "Confirm",
      action: () => {
        console.log("No callback")
      }
    },
    {
      label: "Cancel",
      action: () => {
        console.log("No callback")
      }
    },
  ]
   
	const dispatch = createEventDispatcher();
	const handleMenuItemAction = (action) => dispatch('click', { action });	
</script>

{#if show}
  <div on:click|self class="absolute w-full h-full top-0 left-0 z-50 bg-black bg-opacity-50">
    <div
      class="absolute bg-white rounded-lg shadow-xs dark:bg-gray-800 w-1/3 h-1/4 inset-1/3 py-4 px-6"
    >
      <h1 class="my-4 text-2xl font-semibold text-gray-700 dark:text-gray-200">{title}</h1>
      <p class="text-lg text-gray-700 dark:text-gray-200">
        {msg}
        <span class="font-bold">{item}</span>?
      </p>
      <div class="my-4 py-3 rounded-lg flex">
        {#each callbacks as item}
          {#if typeof item.action === 'string'}
            <a
            class="mr-4 px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple cursor-pointer"
            href={item.action}
            >
              {item.label}
            </a>
          {:else}
            <div 
              class="mr-4 px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple cursor-pointer"
              on:click={() => item.action()}
            >
              <span>{item.label}</span>
            </div>
          {/if}
        {/each}
      </div>
    </div>
  </div>
{/if}
