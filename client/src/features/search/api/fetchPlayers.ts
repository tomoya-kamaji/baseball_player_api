import { createApiClient } from '@/lib/axios/axiosClient'

const fetchPlayers = async (id: string) => {
  const api = createApiClient({
    path: '/players/{id}',
    httpMethod: 'get',
    params: {
      paths: {
        id: id
      }
    }
  })
  const response = await api.request()
  if (response.result === 'success') {
    return response.data
  } else {
    throw response.error
  }
}

export default fetchPlayers
