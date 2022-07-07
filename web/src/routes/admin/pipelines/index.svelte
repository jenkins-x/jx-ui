<script context="module">
  // Todo: move it to endpoints
  export async function load({ fetch }) {
    // ToDo: Use Axios
    const res = await fetch('http://localhost:8080/api/v1/pipelines')
    const pipelines = await res.json()
    if (res.ok) {
      const pipelinesProcessed = pipelines.map((k) => k.spec)
      return {
        props: {
          pipelinesProcessed,
        },
      }
    }
    return {
      status: res.status,
      // ToDo: use new Error()
      error: 'Could not fetch pipelines',
    }
  }
</script>

<script lang="ts">
  import { diffTimes, displayTime } from '$lib/formatDate'
  import Table from '$lib/Components/Table.svelte'
  import type { GridOptions } from 'ag-grid-community'
  import { HTMLInColumnCell } from '$lib/cellRenderer'
  export let pipelinesProcessed

  let gridOptions: GridOptions = {
    defaultColDef: {
      sortable: true,
      resizable: true,
      filter: true,
      flex: 1,
    },
    columnDefs: [
      { headerName: 'Repository', field: 'gitRepository' },
      { headerName: 'Branch', field: 'gitBranch' },
      {
        headerName: 'Build',
        field: 'build',
        cellRenderer: (params) => {
          return HTMLInColumnCell(
            `/pipelines/${params.data.gitOwner}/${params.data.gitRepository}/${params.data.gitBranch}/${params.data.build}`,
            params.data.build
          )
        },
      },
      { headerName: 'Status', field: 'status' },
      {
        headerName: 'Start Time',
        field: 'startedTimestamp',
        valueFormatter: (params) => displayTime(Date.parse(params.data.startedTimestamp)),
      },
      {
        headerName: 'End Time',
        field: 'completedTimestamp',
        valueFormatter: (params) => displayTime(Date.parse(params.data.completedTimestamp)),
      },
      {
        headerName: 'Duration',
        field: 'duration',
        valueFormatter: (params) =>
          diffTimes(
            Date.parse(params.data.startedTimestamp),
            Date.parse(params.data.completedTimestamp)
          ),
      },
    ],
    rowData: pipelinesProcessed,
    pagination: true,
    paginationAutoPageSize: true,
  }
</script>

<svelte:head>
  <title>Pipelines</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Pipelines</h2>
    <!-- New Table -->
    <div class="w-screen h-screen overflow-hidden rounded-lg shadow-xs">
      <Table tableId="pipeline-grid" {gridOptions} />
    </div>
  </div>
</main>
