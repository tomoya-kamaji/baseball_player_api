import { useFetch } from '@/lib/hooks/useFetch'

export function useFetchPlayerById(id: string) {
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
  if (error) return { player: undefined, error }
  if (!data) return { player: undefined, error: undefined }

  const player = {
    id: data.player.id,
    name: data.player.name,
    uniformNumber: data.player.uniform_number,
    atBats: data.player.at_bats,
    hits: data.player.hits,
    homeRuns: data.player.home_runs,
    runsBattedIn: data.player.runs_batted_in,
    walks: data.player.walks
  }
  return { player, error: undefined }
}
