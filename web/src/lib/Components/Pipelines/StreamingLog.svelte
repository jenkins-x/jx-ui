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
  let log = {}

  onMount(() => {
    const ansi_up = new AnsiUp()
    const eventSource = new EventSource(
      `${PUBLIC_BASE_PATH}/api/v1/logs/${owner}/${repository}/${branch}/${build}`
    )
    eventSource.onopen = function (e) {
      console.log('Connection Opened')
    }
    eventSource.onmessage = (event) => {
      log = ansi_up.ansi_to_html(event.data)
      logs.push(log)
      logs = logs
    }
    eventSource.onerror = (e) => {
      console.log('Connection error')
    }
  })
</script>

<DisplayLog {logs} />
