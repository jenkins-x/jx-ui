<script lang="ts">
  // import { barConfig, lineConfig, pieConfig } from '$lib/Templates/Admin/Config/charts'
  // import Bar from 'svelte-chartjs/src/Bar.svelte'
  // import Line from 'svelte-chartjs/src/Line.svelte'
  // import Pie from 'svelte-chartjs/src/Pie.svelte'

  import _i from '$lib/translator'
</script>

<svelte:head>
  <title>{_i('Charts')}</title>
</svelte:head>

<main class="h-full pb-16 overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Charts</h2>
    <!-- CTA -->

    <p class="mb-8 text-gray-600 dark:text-gray-400">
      {_i('Charts are provided by')}
      <a
        class="text-purple-600 dark:text-purple-400 hover:underline"
        href="https://www.chartjs.org/"
      >
        {_i('Chart.js')}
      </a>
      . {_i(
        'Note that the default legends are disabled and you should provide a description for your charts in HTML. See source code for examples.'
      )}
    </p>

    <div class="grid gap-6 mb-8 md:grid-cols-2">
      <!-- Doughnut/Pie chart -->
      <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <h4 class="mb-4 font-semibold text-gray-800 dark:text-gray-300">Doughnut/Pie</h4>
        <!-- <Pie data={pieConfig} /> -->
        <div class="flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400">
          <!-- Chart legend -->
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-blue-600 rounded-full" />
            <span>{_i('Shirts')}</span>
          </div>
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-teal-500 rounded-full" />
            <span>{_i('Shoes')}</span>
          </div>
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full" />
            <span>{_i('Bags')}</span>
          </div>
        </div>
      </div>
      <!-- Lines chart -->
      <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <h4 class="mb-4 font-semibold text-gray-800 dark:text-gray-300">Lines</h4>
        <!-- <Line data={lineConfig} /> -->
        <div class="flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400">
          <!-- Chart legend -->
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-teal-500 rounded-full" />
            <span>Organic</span>
          </div>
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full" />
            <span>Paid</span>
          </div>
        </div>
      </div>
      <!-- Bars chart -->
      <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
        <h4 class="mb-4 font-semibold text-gray-800 dark:text-gray-300">Bars</h4>
        <!-- <Bar data={barConfig} /> -->
        <div class="flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400">
          <!-- Chart legend -->
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-teal-500 rounded-full" />
            <span>{_i('Shoes')}</span>
          </div>
          <div class="flex items-center">
            <span class="inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full" />
            <span>{_i('Bags')}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</main>
