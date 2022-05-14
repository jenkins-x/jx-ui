<script context="module">
  // Todo: move it to endpoints
  export async function load({ fetch }) {
    // ToDo: Use Axios
    const res = await fetch('http://localhost:8080/api/v1/pipelines')
    const pipelines = await res.json()
    if (res.ok) {
      return {
        props: {
          pipelines,
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
  export let pipelines

  function diffTimes(i: Date, j: Date) {
    if (isValid(i) && isValid(j)) {
      return formatDistanceStrict(j, i)
    } else {
      return ''
    }
  }

  function displayTime(i: Date) {
    if (isValid(i)) {
      return format(i, 'Pp')
    } else {
      return ''
    }
  }
</script>

<svelte:head>
  <title>Pipelines</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Pipelines</h2>
    <!-- New Table -->
    <div class="w-full overflow-hidden rounded-lg shadow-xs">
      <div class="w-full overflow-x-auto">
        <table class="w-full whitespace-no-wrap">
          <thead>
            <tr
              class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800"
            >
              <th class="px-4 py-3">Repository</th>
              <th class="px-4 py-3">Branch</th>
              <th class="px-4 py-3">Build</th>
              <th class="px-4 py-3">Status</th>
              <th class="px-4 py-3">Start Time</th>
              <th class="px-4 py-3">End Time</th>
              <th class="px-4 py-3">Duration</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
            {#each pipelines as pipeline}
              <tr class="text-gray-700 dark:text-gray-400">
                <td class="px-4 py-3">
                  <div class="flex items-center text-sm">
                    <div>
                      <a
                        href="/pipelines/{pipeline.spec.build}-{pipeline.spec.gitRepository
                          .split(' ')
                          .join('-')
                          .toLowerCase()}"
                        ><p class="font-semibold">{pipeline.spec.gitRepository}</p></a
                      >
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3 text-sm"> {pipeline.spec.gitBranch} </td>
                <td class="px-4 py-3 text-sm"> {pipeline.spec.build} </td>
                <td class="px-4 py-3 text-xs">
                  <span
                    class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-full dark:bg-green-700 dark:text-green-100"
                  >
                    {pipeline.spec.status}
                  </span>
                </td>
                <td class="px-4 py-3 text-sm">
                  {displayTime(Date.parse(pipeline.spec.startedTimestamp))}
                </td>
                <td class="px-4 py-3 text-sm">
                  {displayTime(Date.parse(pipeline.spec.completedTimestamp))}
                </td>
                <td class="px-4 py-3 text-sm">
                  {diffTimes(
                    Date.parse(pipeline.spec.startedTimestamp),
                    Date.parse(pipeline.spec.completedTimestamp)
                  )}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</main>
