import { error } from '@sveltejs/kit'
// Todo: move it to endpoints
export async function load({ fetch, params }) {
  // ToDo: Use Axios
  const config = { method: 'get', headers: { origin: 'http://localhost:3000' } }
  const { org, repo, branch, build } = params
  const request = new Request(`http://localhost:9200/api/v1/pipelines/${org}/${repo}/${branch}/${build}`, config)
  const res = await fetch(request)

  const pipeline = await res.json()
  if (res.ok) {
    return {
      pipeline,
      org,
      repo,
      branch,
      build,
    }
  }
  throw error(500, 'Could not fetch pipelines')
}
