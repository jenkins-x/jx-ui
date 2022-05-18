<script lang="ts" context="module">
  import _i from '$lib/translator'
  import Plotly from 'plotly.js-dist-min'
  import { onMount } from 'svelte'
  import { page } from '$app/stores'
  let data = [
    {
      z: [
        [1, null, 30, 50, 1],
        [20, 1, 60, 80, 30],
        [30, 60, 1, -10, 20],
        [1, null, 30, 50, 1],
        [20, 1, 60, 80, 30],
        [30, 60, 1, -10, 20],
        [1, null, 30, 50, 1],
        [20, 1, 60, 80, 30],
        [30, 60, 1, -10, 20],
        [1, null, 30, 50, 1],
        [20, 1, 60, 80, 30],
        [30, 60, 1, -10, 20],
      ],
      x: ['2017', '2018', '2019', '2020', '2021'],
      y: ['Jan', 'Feb', 'Mar', 'April', 'May', 'June', 'July', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      type: 'heatmap',
      colorscale: 'Jet',
      hoverongaps: true,
    },
  ]

  let projectProgess = [
    {
      type: 'bar',
      x: [50],
      y: ['Project-1'],
      orientation: 'h',
      text: 'In Progress',
    },
  ]

  var layout = {
    xaxis: { range: [0, 100], autorange: false },
  }

  onMount(() => {
    Plotly.newPlot('heatmap', data)
    Plotly.newPlot('progress-bar', projectProgess, layout)
  })
  // const { id, name } = { ...$page.params }
</script>

<svelte:head>
  <title>{_i('Charts')}</title>
</svelte:head>

<main class="h-full pb-16 overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">#{id}: {name}</h2>
    <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Project Summary</h4>
    <div class="grid gap-6 mb-8 md:grid-cols-2">
      <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <ul class="list-disc">
          <li class="px-4 py-2">Type: PDF</li>
          <li class="px-4 py-2">Type: PDF</li>
          <li class="px-4 py-2">Type: PDF</li>
          <li class="px-4 py-2">Type: PDF</li>
        </ul>
        <p class="text-gray-600 dark:text-gray-400">
          <br />
          Size: 12 MB <br />
          Upload date: April 1, 2050 <br />
        </p>
      </div>
    </div>
    <div class="grid gap-6 mb-8 md:grid-cols-1">
      <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <h4 class="mb-4 font-semibold text-gray-800 dark:text-gray-300">Documents</h4>
        <div
          id="heatmap"
          class="flex justify-center mt-4 text-sm text-gray-600 dark:text-gray-400"
        />
      </div>
    </div>
    <div class="grid gap-6 mb-8 md:grid-cols-2">
      <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <h4 class="mb-4 font-semibold text-gray-800 dark:text-gray-300">Project progress</h4>
        <div
          id="progress-bar"
          class="flex justify-center mt-4 text-sm text-gray-600 dark:text-gray-400"
        />
      </div>
    </div>
  </div>
</main>
