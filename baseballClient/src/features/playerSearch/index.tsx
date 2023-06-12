import { useParams } from 'react-router-dom'
import { PlayerSearchContainer } from './components/PlayerSearchContainer'
import { PlayerDetail } from './components/PlayerDetail'
import { Header } from '../../components/Head/HEAD'
import { ThemeProvider } from '@emotion/react'
import { theme } from '../../themes/theme'

export const PlayerSearchRoute = () => {
  return (
    <div>
      <ThemeProvider theme={theme}>
        <Header />
        <PlayerSearchContainer />
      </ThemeProvider>
    </div>
  )
}

export const PlayerDetailRoute = () => {
  const { id } = useParams()
  if (id == undefined) return <div></div>
  return <PlayerDetail playerId={id} />
}
