<script lang="ts">
  import type { GridOptions } from 'ag-grid-community'
  import Table from '$lib/Components/Table.svelte'
  import { HTMLInColumnCell } from '$lib/cellRenderer'

  export let data

  let gridOptions: GridOptions = {
    defaultColDef: {
      sortable: true,
      resizable: true,
      filter: true,
      flex: 1,
      minWidth: 100,
    },
    columnDefs: [
      { headerName: 'Org', field: 'org' },
      { headerName: 'Provider Name', field: 'providerName' },
      {
        headerName: 'Repo',
        field: 'repo',
        cellRenderer: (params) => {
          let keyData = params.data.repo
          return HTMLInColumnCell(params.data.url, params.data.repo)
        },
      },
    ],
    rowData: data.repositoriesProcessed,
    pagination: true,
    paginationAutoPageSize: true,
  }
</script>

<svelte:head>
  <title>Repositories</title>
</svelte:head>

<main class="h-full overflow-hidden">
  <div class="container px-6 mx-auto grid overflow-hidden">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Repositories</h2>
    <!-- New Table -->
    <div class="w-full h-screen overflow-hidden rounded-lg shadow-xs">
      <Table tableId="repo-grid" {gridOptions} />
    </div>
  </div>
</main>
