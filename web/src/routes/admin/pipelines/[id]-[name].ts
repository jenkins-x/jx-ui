export function get() {
  const pipelines = [
    {
      id: 1,
      repo: 'jenkins-x/jx',
      branch: 'main',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Running',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
    {
      id: 2,
      repo: 'jenkins-x/jx-ui',
      branch: 'main',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Pending',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
    {
      id: 3,
      repo: 'jenkins-x/jx-docs',
      branch: 'main',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Cancelled',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
    {
      id: 4,
      repo: 'jenkins-x/jx3-versions',
      branch: 'PR-1567',
      build: '#3 pr',
      duration: '10 minutes',
      status: 'Timed out',
      startDate: '12/14/2021, 3:12 PM',
      endDate: '12/14/2021, 3:22 PM',
    },
  ]

  return {
    body: {
      pipelines,
    },
  }
}
