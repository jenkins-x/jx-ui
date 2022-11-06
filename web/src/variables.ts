import { PUBLIC_BASE_PATH } from '$env/static/public'
export default {
  authentication: {
    driver: 'laravelSanctum', // null | laravelSanctum
    config: {
      url: `${PUBLIC_BASE_PATH}/sanctum/csrf-token`,
    },
  },
}
