import { useParams } from 'react-router-dom'
import { PlayerDetail } from '../components/PlayerDetail'

export const Player = () => {
  const { id } = useParams()
  if (id == undefined) return <div></div>
  return <PlayerDetail playerId={id} />
}
