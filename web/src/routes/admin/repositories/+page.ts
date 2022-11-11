import { error } from '@sveltejs/kit'
import { PUBLIC_BASE_PATH } from '$env/static/public'

// Todo: move it to endpoints
export async function load({ fetch }) {
  // ToDo: Use Axios
  const config = { method: 'get', headers: { origin: 'http://localhost:3000' } }
  const request = new Request(`${PUBLIC_BASE_PATH}/api/v1/repositories`, config)
  const res = await fetch(request)
  const repositories = await res.json()
  if (res.ok) {
    const repositoriesProcessed = repositories.map((k) => k.spec)
    return {
      repositoriesProcessed,
    }
  }
  throw error(500, 'Could not fetch repositories')
}
