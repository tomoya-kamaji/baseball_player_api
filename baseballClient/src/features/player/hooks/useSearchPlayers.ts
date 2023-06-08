import { useFetch } from '../../../lib/hooks/useFetch'
import { SearchParam, SearchPlayerApiModel } from '../types'

export function useSearchPlayers(param: SearchParam): SearchPlayerApiModel {
  const { data, error } = useFetch({
    path: '/v1/players/search',
    params: {
      query: {
        limit: param.limit || 10,
        max_hits: param.maxHits || 1000,
        min_hits: param.minHits || 0,
        max_home_runs: param.maxHomeRuns || 1000,
        min_home_runs: param.minHomeRuns || 0,
        max_runs_batted_in: param.maxRunsBattedIn || 1000,
        min_runs_batted_in: param.minRunsBattedIn || 0,
        sort_field: param.sortField || 'hits',
        sort_order: param.sortOrder || 'desc'
      }
    },
    onError: (e) => {
      // result='error'の場合のエラー
      if (e.status === 400) {
        console.error(e.data.message)
      }
    }
  })
  if (error) return { players: [], error: new Error(error.data.message) }
  if (!data) return { players: [], error: undefined }

  const players = data.players.map((p) => {
    return {
      id: p.id || '',
      name: p.name || '',
      uniformNumber: p.uniform_number || 0,
      atBats: p.at_bats || 0,
      hits: p.hits || 0,
      homeRuns: p.home_runs || 0,
      runsBattedIn: p.runs_batted_in || 0,
      walks: p.walks || 0
    }
  })
  return { players, error: undefined }
}
