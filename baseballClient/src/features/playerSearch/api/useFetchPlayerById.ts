import { useFetch } from '../../../lib/hooks/useFetch'
import { FetchPlayerApiModel } from '../types'

export function useFetchPlayerById(id: string): FetchPlayerApiModel {
  const { data, error } = useFetch({
    path: '/v1/players/{id}',
    params: { paths: { id: id } },
    onError: (e) => {
      // result='error'の場合のエラー
      if (e.status === 400) {
        console.error(e.data.message)
      }
    }
  })
  if (error) return { player: undefined, error: new Error(error.data.message) }
  if (!data) return { player: undefined, error: undefined }

  const player = {
    id: data.player.id || '',
    name: data.player.name || '',
    uniformNumber: data.player.uniform_number || 0,
    atBats: data.player.at_bats || 0,
    hits: data.player.hits || 0,
    homeRuns: data.player.home_runs || 0,
    runsBattedIn: data.player.runs_batted_in || 0,
    walks: data.player.walks || 0
  }
  return { player, error: undefined }
}
