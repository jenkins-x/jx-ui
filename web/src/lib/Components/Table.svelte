<script lang="ts">
  import { Grid, GridOptions } from 'ag-grid-community'

  import 'ag-grid-community/dist/styles/ag-grid.css'
  import 'ag-grid-community/dist/styles/ag-theme-alpine-dark.css'
  import { onMount } from 'svelte'

  export let tableId: string = ''
  export let gridOptions: GridOptions = {}
  onMount(() => {
    var eGridDiv = document.querySelector<HTMLElement>(`#${tableId}`)
    new Grid(eGridDiv, gridOptions)
  })
</script>

<div id={tableId} class="h-3/4 w-9/10 ag-theme-alpine-dark" />
