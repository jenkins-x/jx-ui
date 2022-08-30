import { error } from '@sveltejs/kit'

// Todo: move it to endpoints
export const load = async({ fetch }) => {
  // ToDo: Use Axios
  const res = await fetch('http://localhost:9200/api/v1/pipelines')
  const pipelines = await res.json()

  if (res.ok) {
    return {
      pipelines,
    }
  }
  throw error(500, 'Could not fetch pipelines')
}
