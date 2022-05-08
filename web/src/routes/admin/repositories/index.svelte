<script context="module">
  // Todo: move it to endpoints
  export async function load({ fetch }) {
    // ToDo: Use Axios
    const res = await fetch('http://localhost:8080/api/v1/repositories')
    const repositories = await res.json()
    if (res.ok) {
      return {
        props: {
          repositories,
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
  export let repositories
</script>

<svelte:head>
  <title>Repositories</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Repositories</h2>
    <!-- New Table -->
    <div class="w-full overflow-hidden rounded-lg shadow-xs">
      <div class="w-full overflow-x-auto">
        <table class="w-full whitespace-no-wrap">
          <thead>
            <tr
              class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800"
            >
              <th class="px-4 py-3">Org</th>
              <th class="px-4 py-3">Provider Name</th>
              <th class="px-4 py-3">Repo</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
            {#each repositories as repository}
              <tr class="text-gray-700 dark:text-gray-400">
                <td class="px-4 py-3 text-sm"> {repository.spec.org} </td>
                <td class="px-4 py-3 text-sm"> {repository.spec.providerName} </td>
                <td class="px-4 py-3 text-sm">
                  <a href={repository.spec.url}>{repository.spec.repo}</a>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</main>
