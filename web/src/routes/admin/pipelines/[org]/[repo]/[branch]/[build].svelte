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
</script>

<script lang="ts">
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
</script>

<svelte:head>
  <title>{'Pipeline build'}</title>
</svelte:head>

{#if pipeline}
  <main class="h-full pb-16 overflow-y-auto">
    <div class="container px-6 mx-auto grid">
      <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Pipeline details</h2>
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
