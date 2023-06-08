import { useParams } from 'react-router-dom'
import { PlayerDetail, PlayerIndex } from '../components'

export const PlayerIndexRoute = () => {
  return (
    <PlayerIndex
      param={{
        limit: 10,
        maxHits: undefined,
        maxHomeRuns: undefined,
        maxRunsBattedIn: undefined,
        minHits: undefined,
        minHomeRuns: undefined,
        minRunsBattedIn: undefined,
        sortField: 'hits',
        sortOrder: 'desc'
      }}
    />
  )
}

export const PlayerDetailRoute = () => {
  const { id } = useParams()
  if (id == undefined) return <div></div>
  return <PlayerDetail playerId={id} />
}
