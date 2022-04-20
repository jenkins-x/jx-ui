<script lang="ts">
  import { clickOutside } from '$lib/IOEvents/click'
  import { keydownEscape } from '$lib/IOEvents/keydown'

  let isModalOpen = false

  const openModal = () => {
    isModalOpen = true
  }

  const closeModal = () => {
    isModalOpen = false
  }
  let pipelines = [
    {
      id: 1,
      repo: 'jenkins-x/jx',
      branch: 'main',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Running',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
    {
      id: 2,
      repo: 'jenkins-x/jx-ui',
      branch: 'main',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Pending',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
    {
      id: 3,
      repo: 'jenkins-x/jx-docs',
      branch: 'main',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Cancelled',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
    {
      id: 4,
      repo: 'jenkins-x/jx3-versions',
      branch: 'PR-1567',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Timed out',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
  ]
  function onSubmit(e) {
    console.log('Test')
    const formData = new FormData(e.target)
    let id = pipelines.length + 1
    const data = { id }
    for (let field of formData) {
      const [key, value] = field
      data[key] = value
    }
    data['branch'] = 0
    data['status'] = 'Not Started'
    data['build'] = 0

    pipelines = [...pipelines, data]
    isModalOpen = false
  }
</script>

<svelte:head>
  <title>Dashboard</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Dashboard</h2>
    <!-- CTA -->

    <!-- Cards -->
    <div class="grid gap-6 mb-8 md:grid-cols-1 xl:grid-cols-3">
      <!-- pipelines -->
      <div class="flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <div
          class="p-3 mr-4 text-orange-500 bg-orange-100 rounded-full dark:text-orange-100 dark:bg-orange-500"
        >
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
            <path
              d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
            />
          </svg>
        </div>
        <div>
          <a href="/pipelines">
            <p class="mb-2 text-sm font-medium text-gray-600 dark:text-gray-400">Pipelines</p>
            <p class="text-lg font-semibold text-gray-700 dark:text-gray-200">50</p>
          </a>
        </div>
      </div>
      <!-- Users -->
      <div class="flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <div
          class="p-3 mr-4 text-orange-500 bg-orange-100 rounded-full dark:text-orange-100 dark:bg-orange-500"
        >
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
            <path
              d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
            />
          </svg>
        </div>
        <div>
          <p class="mb-2 text-sm font-medium text-gray-600 dark:text-gray-400">Repositories</p>
          <p class="text-lg font-semibold text-gray-700 dark:text-gray-200">100</p>
        </div>
      </div>
      <!-- Documents -->
      <div class="flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <div
          class="p-3 mr-4 text-orange-500 bg-orange-100 rounded-full dark:text-orange-100 dark:bg-orange-500"
        >
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
            <path
              d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
            />
          </svg>
        </div>
        <div>
          <p class="mb-2 text-sm font-medium text-gray-600 dark:text-gray-400">Total Secrets</p>
          <p class="text-lg font-semibold text-gray-700 dark:text-gray-200">1000</p>
        </div>
      </div>
    </div>
    <div class="px-4 py-3 mb-8 rounded-lg flex flex-row-reverse">
      <div
        class="px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple"
      >
        <a href="/pipelines"> Import a repository </a>
      </div>
    </div>

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
                        href="/pipelines/{pipeline.id}-{pipeline.repo
                          .split(' ')
                          .join('-')
                          .toLowerCase()}"><p class="font-semibold">{pipeline.repo}</p></a
                      >
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3 text-sm"> {pipeline.branch} </td>
                <td class="px-4 py-3 text-sm"> {pipeline.build} </td>
                <td class="px-4 py-3 text-xs">
                  <span
                    class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-full dark:bg-green-700 dark:text-green-100"
                  >
                    {pipeline.status}
                  </span>
                </td>
                <td class="px-4 py-3 text-sm"> {pipeline.startDate} </td>
                <td class="px-4 py-3 text-sm"> {pipeline.endDate} </td>
                <td class="px-4 py-3 text-sm"> {pipeline.duration} </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
      <div
        class="grid px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800"
      >
        <span class="flex items-center col-span-3"> Showing 21-30 of 100 </span>
        <span class="col-span-2" />
        <!-- Pagination -->
        <span class="flex col-span-4 mt-2 sm:mt-auto sm:justify-end">
          <nav aria-label="Table navigation">
            <ul class="inline-flex items-center">
              <li>
                <button
                  class="px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple"
                  aria-label="Previous"
                >
                  <svg aria-hidden="true" class="w-4 h-4 fill-current" viewBox="0 0 20 20">
                    <path
                      d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                      clip-rule="evenodd"
                      fill-rule="evenodd"
                    />
                  </svg>
                </button>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                  1
                </button>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                  2
                </button>
              </li>
              <li>
                <button
                  class="px-3 py-1 text-white transition-colors duration-150 bg-purple-600 border border-r-0 border-purple-600 rounded-md focus:outline-none focus:shadow-outline-purple"
                >
                  3
                </button>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                  4
                </button>
              </li>
              <li>
                <span class="px-3 py-1">...</span>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                  8
                </button>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                  9
                </button>
              </li>
              <li>
                <button
                  class="px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple"
                  aria-label="Next"
                >
                  <svg class="w-4 h-4 fill-current" aria-hidden="true" viewBox="0 0 20 20">
                    <path
                      d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                      clip-rule="evenodd"
                      fill-rule="evenodd"
                    />
                  </svg>
                </button>
              </li>
            </ul>
          </nav>
        </span>
      </div>
    </div>
  </div>
</main>
