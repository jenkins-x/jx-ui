<script lang="ts">
  import { diffTimes, displayTime } from '$lib/formatDate'
  import Table from '$lib/Components/Table.svelte'
  import type { GridOptions } from 'ag-grid-community'
  import { BTNInColumnCell, HTMLInColumnCell } from '$lib/cellRenderer'
  import { stopPipelineHandler } from '$lib/Components/Pipelines/StopPipeline'
  import Modal from '$src/lib/Components/Modal.svelte'
  
  export let data

  let show: boolean = false
  let title: string
  let msg: string
  let item: string

  let owner: string
  let branch: string
  let repository: string
  let build: string

  let callbacks = [
    {
      label: 'Yes',
      action: () => {
        stopPipelineHandler({
          owner,
          branch,
          repository,
          build,
        })
        show = false
      },
    },
    {
      label: 'No',
      action: () => {
        show = false
      },
    },
  ]

  const toggleModal = () => {
    show = !show
  }

  const stopPipeline = (name, pipeline) => {
    title = 'Stop Pipeline'
    msg = 'Are you sure, you want to stop'
    item = name

    owner = pipeline.owner
    branch = pipeline.branch
    repository = pipeline.repository
    build = pipeline.build

    toggleModal()
  }

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
          diffTimes(Date.parse(params.data.StartTime), Date.parse(params.data.EndTime)),
      },
      {
        headerName: 'Action',
        field: 'action',
        cellRenderer: (params) => {
          return BTNInColumnCell(
            `${params.data.Status}`,
            `${params.data.GitOwner}`,
            `${params.data.GitBranch}`,
            `${params.data.GitRepository}`,
            `${params.data.Build}`,
            `${params.data.Name}`,
            stopPipeline,
          )
        },
      },
    ],
    rowData: data.pipelines,
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
      <Modal {title} {msg} {item} {show} on:click={toggleModal} {callbacks} />
    </div>
  </div>
</main>
