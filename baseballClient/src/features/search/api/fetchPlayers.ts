import { useQuery } from 'react-query'
import { FetchPlayerApiModel } from '../types'
import { createApiClient } from '../../../lib/axios/axiosClient'

const fetchPlayer = async (id: string): Promise<FetchPlayerApiModel> => {
  const api = createApiClient({
    path: '/v1/players/{id}',
    httpMethod: 'get',
    params: {
      paths: {
        id: id
      }
    }
  })
  const response = await api.request()
  if (response.result === 'success') {
    return {
      player: {
        id: response.data.player.id || '',
        name: response.data.player.name || '',
        uniformNumber: response.data.player.uniform_number || 0,
        atBats: response.data.player.at_bats || 0,
        hits: response.data.player.hits || 0,
        homeRuns: response.data.player.home_runs || 0,
        runsBattedIn: response.data.player.runs_batted_in || 0,
        walks: response.data.player.walks || 0
      }
    }
  } else {
    console.log('request error', response.error)
    throw response.error
  }
}

export const useFetchPlayer = (id: string) => {
  return useQuery<FetchPlayerApiModel, Error>({
    queryKey: ['id', id],
    queryFn: () => fetchPlayer(id)
  })
}
