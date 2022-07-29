<script context="module">
  import { onMount } from 'svelte'
  // Todo: move it to endpoints
  export async function load({ fetch, params }) {
    // ToDo: Use Axios
    
    const { org, repo, branch, build } = params
    const res = await fetch(
      `http://localhost:8080/api/v1/pipelines/${org}/${repo}/${branch}/${build}`
    )

    const pipeline = await res.json()

    if (res.ok) {
      return {
        props: {
          pipeline,
          org,
          repo,
          branch,
          build,
        },
      }
    }
    return {
      status: res.status,
      // ToDo: use new Error()
      error: 'Could not fetch pipelines',
    }
  }

  export const stopPipelineHandler = async ( params ) => {
    const { owner, repository, branch, build } = params
    const res = await fetch(`http://localhost:8080/api/v1/pipelines/${owner}/${repository}/${branch}/${build}`, {
      method: "PUT"
    })
    
    const result = await res.json()

    // TODO: improve error handling (@rajatgupta24)
    if (result){
      console.log("pipeline has successfully stopped")
    }
    else {
      console.log("some error occured, pipeline is not stopped")
    }
  }
</script>

<script lang="ts">
  import Modal from '$src/lib/Components/Modal.svelte'
  import ArchivedLog from '$src/lib/Components/Pipelines/ArchivedLog.svelte'
  import StreamingLog from '$src/lib/Components/Pipelines/StreamingLog.svelte'
  
  import { diffTimes, displayTime } from '$src/lib/formatDate'
  import { fromUnixTime } from 'date-fns'
  
  export let pipeline
  
  let {
    spec: {
      gitOwner: owner,
      gitBranch: branch,
      gitRepository: repository,
      build,
      context,
      steps,
      completedTimestamp,
      startedTimestamp,
      status,
    },
    metadata: { name, namespace },
  } = pipeline

  let show = false
  let title = ""
  let msg = ""
  let item = ""

  const toggleModal = () => {
    show = !show
  }

  let callbacks = [
    {
      label: "Yes",
      action: () => {
        stopPipelineHandler({
          owner,
          branch,
          repository,
          build
        });
        show = false
      }
    },
    {
      label: "No",
      action: () => {
        show = false
      }
    },
  ]

  const startPipeline = () => {
    title = "Start Pipeline"
    msg = "Are you sure, you want to start"
    item = name

    toggleModal()
  }

  const stopPipeline = () => {
    title = "Stop Pipeline"
    msg = "Are you sure, you want to stop"
    item = name

    toggleModal()
  }
</script>

<svelte:head>
  <title>{'Pipeline build'}</title>
</svelte:head>

{#if pipeline}
  <main class="h-full pb-16 overflow-y-auto">
    <div class="container px-6 mx-auto grid">
      <div class="flex justify-between align-baseline">
        <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Pipeline details</h2>

        {#if status.toLowerCase() === "pending" || status.toLowerCase() === "running"}
          <div class="flex  text-purple-600 dark:text-purple-300">
            <button 
              class="rounded-md focus:outline-none focus:shadow-outline-purple-500"
              on:click={startPipeline}
              aria-label="Toggle color mode"
            >
              <svg class="ml-4 w-10 h-10" aria-hidden="true" fill="currentColor" viewBox="0 0 512 512">
                <!--! Font Awesome Pro 6.1.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc. -->
                <path d="M512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256zM176 168V344C176 352.7 180.7 360.7 188.3 364.9C195.8 369.2 205.1 369 212.5 364.5L356.5 276.5C363.6 272.1 368 264.4 368 256C368 247.6 363.6 239.9 356.5 235.5L212.5 147.5C205.1 142.1 195.8 142.8 188.3 147.1C180.7 151.3 176 159.3 176 168V168z"/>
              </svg>
            </button>

            <button
              class="rounded-md focus:outline-none focus:shadow-outline-purple-500"
              aria-label="Toggle color mode"
              on:click={stopPipeline}
            >
              <svg class="ml-4 w-10 h-10" aria-hidden="true" fill="currentColor" viewBox="0 0 512 512">
                <!--! Font Awesome Pro 6.1.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc. -->
                <path d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM175 208.1L222.1 255.1L175 303C165.7 312.4 165.7 327.6 175 336.1C184.4 346.3 199.6 346.3 208.1 336.1L255.1 289.9L303 336.1C312.4 346.3 327.6 346.3 336.1 336.1C346.3 327.6 346.3 312.4 336.1 303L289.9 255.1L336.1 208.1C346.3 199.6 346.3 184.4 336.1 175C327.6 165.7 312.4 165.7 303 175L255.1 222.1L208.1 175C199.6 165.7 184.4 165.7 175 175C165.7 184.4 165.7 199.6 175 208.1V208.1z"/>
              </svg>
            </button>
          </div>
        {/if}
      </div>

      <Modal title={title} msg={msg} item={item} show={show} on:click={toggleModal} callbacks={callbacks}/>

      <div class="grid gap-6 mb-8 md:grid-cols-3">
        <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
          <ul class="list-none">
            <li class="px-4 py-2">Pipeline activity: {name}</li>
            <li class="px-4 py-2">Organization: {owner}</li>
            <li class="px-4 py-2">Repository: {repository}</li>
            <li class="px-4 py-2">Branch: {branch}</li>
            <li class="px-4 py-2">Build: {build}</li>
            <li class="px-4 py-2">Stages: {steps.length}</li>
            <li class="px-4 py-2">Steps: 6</li>
          </ul>
        </div>
        <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
          <ul class="list-none">
            <li class="px-4 py-2">Context: {context}</li>
            <li class="px-4 py-2">
              Namespace: {pipeline.metadata.namespace} (Should this be shown?)
            </li>
            <li class="px-4 py-2">Author: release</li>
            <li class="px-4 py-2">Commit: release</li>
            <li class="px-4 py-2">Event: release</li>
          </ul>
        </div>
        <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
          <ul class="list-none">
            <li class="px-4 py-2">Status: {status}</li>
            <li class="px-4 py-2">
              Started: {displayTime(Date.parse(startedTimestamp))}
            </li>
            <li class="px-4 py-2">
              Finished: {displayTime(Date.parse(completedTimestamp))}
            </li>
            <li class="px-4 py-2">
              Duration: {diffTimes(Date.parse(completedTimestamp), Date.parse(startedTimestamp))}
            </li>
            {#if status == 'failed'}
              <li class="px-4 py-2">Failed reason:</li>
            {/if}
          </ul>
        </div>
      </div>
      <!-- Repeat this block for all stages -->
      {#if status == 'Running' || status == 'pending'}
        <StreamingLog {owner} {repository} {branch} {build} />
      {:else}
        <ArchivedLog {owner} {repository} {branch} {build} />
      {/if}
    </div>
  </main>
{/if}
