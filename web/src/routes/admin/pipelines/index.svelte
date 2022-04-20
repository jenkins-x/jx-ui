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
  let projects = [
    {
      id: 1,
      name: 'Belmont high school',
      nDoc: 100,
      amount: 90000,
      type: 'Public',
      status: 'In progress',
      startDate: '6/10/2020',
      endDate: '6/10/2020',
    },
    {
      id: 2,
      name: 'Boston Public School',
      nDoc: 100,
      amount: 90000,
      type: 'Public',
      status: 'Completed',
      startDate: '6/10/2020',
      endDate: '6/10/2020',
    },
    {
      id: 3,
      name: 'Arsenal Yards',
      nDoc: 100,
      amount: 90000,
      type: 'Mixed use',
      status: 'Not started',
      startDate: '6/10/2020',
      endDate: '6/10/2020',
    },
    {
      id: 4,
      name: 'Revere Cinema',
      nDoc: 100,
      amount: 90000,
      type: 'Private',
      status: 'Not started',
      startDate: '6/10/2020',
      endDate: '6/10/2020',
    },
  ]
  function onSubmit(e) {
    console.log('Test')
    const formData = new FormData(e.target)
    let id = projects.length + 1
    const data = { id }
    for (let field of formData) {
      const [key, value] = field
      data[key] = value
    }
    data['nDoc'] = 0
    data['status'] = 'Not Started'
    data['amount'] = 0

    projects = [...projects, data]
    isModalOpen = false
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
              <th class="px-4 py-3">Pipelines</th>
              <th class="px-4 py-3"></th>
              <th class="px-4 py-3">Amount (in dollars)</th>
              <th class="px-4 py-3">Type</th>
              <th class="px-4 py-3">Status</th>
              <th class="px-4 py-3">Start Date</th>
              <th class="px-4 py-3">End Date</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
            {#each projects as project}
              <tr class="text-gray-700 dark:text-gray-400">
                <td class="px-4 py-3">
                  <div class="flex items-center text-sm">
                    <!-- Avatar with inset shadow -->
                    <div class="relative hidden w-8 h-8 mr-3 rounded-full md:block">
                      <img
                        class="object-cover w-full h-full rounded-full"
                        src="https://images.unsplash.com/flagged/photo-1570612861542-284f4c12e75f?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ"
                        alt=""
                        loading="lazy"
                      />
                      <div class="absolute inset-0 rounded-full shadow-inner" aria-hidden="true" />
                    </div>
                    <div>
                      <a
                        href="/projects/{project.id}-{project.name
                          .split(' ')
                          .join('-')
                          .toLowerCase()}"><p class="font-semibold">{project.name}</p></a
                      >
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3 text-sm"> {project.nDoc} </td>
                <td class="px-4 py-3 text-sm"> {project.amount} </td>
                <td class="px-4 py-3 text-sm"> {project.type} </td>
                <td class="px-4 py-3 text-xs">
                  <span
                    class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-full dark:bg-green-700 dark:text-green-100"
                  >
                    {project.status}
                  </span>
                </td>
                <td class="px-4 py-3 text-sm"> {project.startDate} </td>
                <td class="px-4 py-3 text-sm"> {project.endDate} </td>
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
