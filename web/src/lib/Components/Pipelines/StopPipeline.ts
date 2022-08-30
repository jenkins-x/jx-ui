export const stopPipelineHandler = async (params) => {
    const { owner, repository, branch, build } = params
    const res = await fetch(
      `http://localhost:9200/api/v1/pipelines/${owner}/${repository}/${branch}/${build}`,
      {
        method: 'PUT',
      }
    )
    const result = await res.json()
  
    // TODO: improve error handling (@rajatgupta24)
    if (result) {
      console.log('pipeline has successfully stopped')
    } else {
      console.log('some error occured, pipeline is not stopped')
    }
  }
