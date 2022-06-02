<script lang="ts">
  import Card from '$lib/Components/Card.svelte'
</script>

<svelte:head>
  <title>Dashboard</title>
</svelte:head>

<main class="h-full overflow-y-auto">
  <div class="container px-6 mx-auto grid">
    <h2 id="dashboard-title" class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">
      Dashboard
    </h2>
    <!-- CTA -->

    <!-- Cards -->
    <div class="grid gap-6 mb-8 md:grid-cols-1 xl:grid-cols-3">
      <Card title="Pipelines" num="50" link="/pipelines" />
      <!-- Users -->
      <Card title="Repositories" num="20" link="/repositories" />
      <!-- Documents -->
      <Card title="External secrets" num="12" link="/repositories" />
    </div>
    <div class="px-4 py-3 mb-8 rounded-lg flex flex-row-reverse">
      <div
        class="px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple"
      >
        <a href="/pipelines"> Import a repository </a>
      </div>
    </div>
  </div>
</main>
