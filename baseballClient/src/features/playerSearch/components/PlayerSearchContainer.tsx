import { PlayerList } from './PlayerList'
import { PlayerSearchFilter } from './PlayerSearchFilter'

export const PlayerSearchContainer = (): JSX.Element => {
  return (
    <div>
      <h2>選手検索</h2>
      <div style={{ display: 'flex', flexDirection: 'row' }}>
        <PlayerSearchFilter />
        <PlayerList />
      </div>
    </div>
  )
}
