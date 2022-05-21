<script lang="ts">
  export let logs = []
</script>

<div class="grid gap-6 mb-8 md:grid-cols-1">
  <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
    {#each logs as log}
      <!-- Run through html sanitizer to avoid XSS attack  -->
      {@html log} <br />
    {/each}
  </div>
</div>
