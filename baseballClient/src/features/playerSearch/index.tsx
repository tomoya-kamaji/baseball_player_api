import { useParams } from 'react-router-dom'
import { PlayerDetail, PlayerIndex } from './components'

export const PlayerSearchRoute = () => {
  return <PlayerIndex />
}

export const PlayerDetailRoute = () => {
  const { id } = useParams()
  if (id == undefined) return <div></div>
  return <PlayerDetail playerId={id} />
}
