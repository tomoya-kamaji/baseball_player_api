import { useParams } from 'react-router-dom'
import { PlayerSearchContainer } from './components/PlayerSearchContainer'
import { PlayerDetail } from './components/PlayerDetail'

export const PlayerSearchRoute = () => {
  return <PlayerSearchContainer />
}

export const PlayerDetailRoute = () => {
  const { id } = useParams()
  if (id == undefined) return <div></div>
  return <PlayerDetail playerId={id} />
}
