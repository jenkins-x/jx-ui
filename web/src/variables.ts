export default {
  authentication: {
    driver: 'laravelSanctum', // null | laravelSanctum
    config: {
      url: 'http://localhost:9200/sanctum/csrf-token',
    },
  },
}
