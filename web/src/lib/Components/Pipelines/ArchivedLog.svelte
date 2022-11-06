<script lang="ts">
  import { onMount } from 'svelte'
  import { default as AnsiUp } from 'ansi_up'
  import DisplayLog from './DisplayLog.svelte'
  import { PUBLIC_BASE_PATH } from '$env/static/public'
  export let owner: string
  export let repository: string
  export let branch: string
  export let build: string
  let logs = []
  let log_processed = []
  onMount(async () => {
    const ansi_up = new AnsiUp()
    const res2 = await fetch(
      `${PUBLIC_BASE_PATH}/api/v1/logs_archived/${owner}/${repository}/${branch}/${build}`
    )
    logs = await res2.json()
    log_processed = logs.map((log) => ansi_up.ansi_to_html(log))
  })
</script>

{#if logs}
  <DisplayLog logs={log_processed} />
{/if}
