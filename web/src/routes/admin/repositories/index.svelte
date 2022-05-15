<script context="module">
  // Todo: move it to endpoints
  export async function load({ fetch }) {
    // ToDo: Use Axios
    const res = await fetch('http://localhost:8080/api/v1/repositories')
    const repositories = await res.json()
    if (res.ok) {
      const repositoriesProcessed = repositories.map((k) => k.spec)
      return {
        props: {
          repositoriesProcessed,
        },
      }
    }
    return {
      status: res.status,
      // ToDo: use new Error()
      error: 'Could not fetch repositories',
    }
  }
</script>

<script lang="ts">
  import { onMount } from 'svelte'
  import { Grid, GridOptions } from 'ag-grid-community'

  import 'ag-grid-community/dist/styles/ag-grid.css'
  import 'ag-grid-community/dist/styles/ag-theme-alpine-dark.css'
  export let repositoriesProcessed

  let gridOptions: GridOptions = {
    defaultColDef: {
      editable: true,
      sortable: true,
      resizable: true,
      filter: true,
      flex: 1,
      minWidth: 100,
    },
    columnDefs: [
      { headerName: 'Org', field: 'org' },
      { headerName: 'Provider Name', field: 'providerName' },
      { headerName: 'Repo', field: 'repo' },
    ],
    rowData: repositoriesProcessed,
    pagination: true,
    paginationAutoPageSize: true,
  }
  onMount(() => {
    var eGridDiv = document.querySelector<HTMLElement>('#repo-grid')
    new Grid(eGridDiv, gridOptions)
  })
</script>

<svelte:head>
  <title>Repositories</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Repositories</h2>
    <!-- New Table -->
    <div class="w-screen h-screen overflow-hidden rounded-lg shadow-xs">
      <div id="repo-grid" class="h-3/4 w-3/4 ag-theme-alpine-dark" />
    </div>
  </div>
</main>
