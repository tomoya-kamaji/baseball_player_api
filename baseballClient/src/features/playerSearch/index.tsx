import { useParams } from 'react-router-dom'
import { PlayerDetail, PlayerIndex } from './components'
import { SearchParam } from './types/index'

export const PlayerSearchRoute = () => {
  const searchParam: SearchParam = {
    limit: 10,
    maxHits: undefined,
    maxHomeRuns: undefined,
    maxRunsBattedIn: undefined,
    minHits: undefined,
    minHomeRuns: undefined,
    minRunsBattedIn: undefined,
    sortField: 'hits',
    sortOrder: 'desc'
  }
  return <PlayerIndex param={searchParam} />
}

export const PlayerDetailRoute = () => {
  const { id } = useParams()
  if (id == undefined) return <div></div>
  return <PlayerDetail playerId={id} />
}
