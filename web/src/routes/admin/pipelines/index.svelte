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
  import { isValid, formatDistanceStrict, format } from 'date-fns'
  import { Grid, GridOptions } from 'ag-grid-community'

  import 'ag-grid-community/dist/styles/ag-grid.css'
  import 'ag-grid-community/dist/styles/ag-theme-alpine-dark.css'
  import { onMount } from 'svelte'
  export let pipelinesProcessed

  function diffTimes(i: number, j: number) {
    if (isValid(i) && isValid(j)) {
      return formatDistanceStrict(j, i)
    }
    return ''
  }

  function displayTime(i: number) {
    if (isValid(i)) {
      return format(i, 'Pp')
    } else {
      return ''
    }
  }

  let gridOptions: GridOptions = {
    defaultColDef: {
      editable: true,
      enableValue: true,
      sortable: true,
      resizable: true,
      filter: true,
      flex: 1,
      minWidth: 100,
    },
    columnDefs: [
      { headerName: 'Repository', field: 'gitRepository' },
      { headerName: 'Branch', field: 'gitBranch' },
      { headerName: 'Build', field: 'build' },
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
  onMount(() => {
    var eGridDiv = document.querySelector<HTMLElement>('#pipeline-grid')
    new Grid(eGridDiv, gridOptions)
  })
</script>

<svelte:head>
  <title>Pipelines</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Pipelines</h2>
    <!-- New Table -->
    <div class="w-screen h-screen overflow-hidden rounded-lg shadow-xs">
      <div id="pipeline-grid" class="h-3/4 w-3/4 ag-theme-alpine-dark" />
    </div>
  </div>
</main>
