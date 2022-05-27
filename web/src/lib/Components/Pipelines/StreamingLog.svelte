<script lang="ts">
  import { onMount } from 'svelte'
  import { default as AnsiUp } from 'ansi_up'

  export let owner: string
  export let repository: string
  export let branch: string
  export let build: string
  let log = ''

  onMount(() => {
    const ansi_up = new AnsiUp()
    const eventSource = new EventSource(
      `http://localhost:8080/api/v1/logs/${owner}/${repository}/${branch}/${build}`
    )
    eventSource.onopen = function (e) {
      console.log('Connection Opened')
    }
    eventSource.onmessage = (event) => {
      // ToDo: not very performant, need better ways to displaying live/streaming logs
      log += ansi_up.ansi_to_html(event.data) + '<br />'
    }
    eventSource.onerror = (e) => {
      console.log('Connection error')
      console.log(e)
    }
  })
</script>

<p>{@html log}</p>
