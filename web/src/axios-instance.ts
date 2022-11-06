import axios from 'axios'
import { PUBLIC_BASE_PATH } from '$env/static/public'

const instance = axios.create({
  baseURL: `${PUBLIC_BASE_PATH}/api/v1`,
})

export default instance
