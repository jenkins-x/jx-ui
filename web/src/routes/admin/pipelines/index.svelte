<script context="module">
  // Todo: move it to endpoints
  export async function load({ fetch }) {
    // ToDo: Use Axios
    const res = await fetch('http://localhost:8080/api/v1/pipelines')
    const pipelines = await res.json()
    if (res.ok) {
      return {
        props: {
          pipelines
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
  export let pipelines

  let gridOptions: GridOptions = {
    defaultColDef: {
      sortable: true,
      resizable: true,
      filter: true,
      flex: 1,
    },
    columnDefs: [
      { headerName: 'Repository', field: 'GitRepository' },
      { headerName: 'Branch', field: 'GitBranch' },
      {
        headerName: 'Build',
        field: 'Build',
        cellRenderer: (params) => {
          return HTMLInColumnCell(
            `/pipelines/${params.data.GitOwner}/${params.data.GitRepository}/${params.data.GitBranch}/${params.data.Build}`,
            params.data.Build
          )
        },
      },
      { headerName: 'Status', field: 'Status' },
      {
        headerName: 'Start Time',
        field: 'StartTime',
        valueFormatter: (params) => displayTime(Date.parse(params.data.StartTime)),
      },
      {
        headerName: 'End Time',
        field: 'EndTime',
        valueFormatter: (params) => displayTime(Date.parse(params.data.EndTime)),
      },
      {
        headerName: 'Duration',
        field: 'duration',
        valueFormatter: (params) =>
          diffTimes(
            Date.parse(params.data.StartTime),
            Date.parse(params.data.EndTime)
          ),
      },
    ],
    rowData: pipelines,
    pagination: true,
    paginationAutoPageSize: true,
  }
</script>

<svelte:head>
  <title>Pipelines</title>
</svelte:head>

<main class="h-full overflow-hidden">
  <div class="container px-6 mx-auto grid overflow-hidden">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Pipelines</h2>
    <!-- New Table -->
    <div class="w-full h-screen overflow-hidden rounded-lg shadow-xs">
      <Table tableId="pipeline-grid" {gridOptions} />
    </div>
  </div>
</main>
