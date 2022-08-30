import { error } from '@sveltejs/kit'
// Todo: move it to endpoints
export async function load({ fetch, params }) {
  // ToDo: Use Axios

  const { org, repo, branch, build } = params
  const res = await fetch(
    `http://localhost:9200/api/v1/pipelines/${org}/${repo}/${branch}/${build}`
  )

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
