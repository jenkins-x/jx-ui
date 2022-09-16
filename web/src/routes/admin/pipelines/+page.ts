import { error } from '@sveltejs/kit'

// Todo: move it to endpoints
export const load = async({ fetch }) => {
  // ToDo: Use Axios
  const config = { method: 'get', headers: { origin: 'http://localhost:3000' } }
  const request = new Request('http://localhost:9200/api/v1/pipelines', config)
  const res = await fetch(request)
  
  const pipelines = await res.json()

  if (res.ok) {
    return {
      pipelines,
    }
  }
  throw error(500, 'Could not fetch pipelines')
}
