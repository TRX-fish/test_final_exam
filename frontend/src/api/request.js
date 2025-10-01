import axios from 'axios'

const api = axios.create({
  baseURL: 'http://127.0.0.1:5000', // 后端API地址
  timeout: 5000
})

// 请求拦截器：自动加token
api.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers['Authorization'] = 'Bearer ' + token
  }
  return config
}, error => Promise.reject(error))

export default api 