import { error } from '@sveltejs/kit'

// Todo: move it to endpoints
export async function load({ fetch }) {
  // ToDo: Use Axios
  const res = await fetch('http://localhost:9200/api/v1/repositories')
  const repositories = await res.json()
  if (res.ok) {
    const repositoriesProcessed = repositories.map((k) => k.spec)
    return {
      repositoriesProcessed,
    }
  }
  throw error(500, 'Could not fetch repositories')
}
