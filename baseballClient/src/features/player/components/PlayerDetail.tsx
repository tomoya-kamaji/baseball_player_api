import { Typography, Card, CardContent, styled, ThemeProvider } from '@mui/material'
import { theme } from '../../../themes/theme'
import { useFetchPlayerById } from '../hooks/useFetchPlayerById'

const CustomCard = styled(Card)`
  width: 300px;
  margin: 16px;
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
  transition: 0.3s;
  background-color: #f1f1f1;
  &:hover {
    box-shadow: 0 16px 32px 0 rgba(0, 0, 0, 0.2);
  }
`
export const PlayerDetail = ({ playerId }: { playerId: string }): JSX.Element => {
  const { player } = useFetchPlayerById(playerId)
  if (player == undefined) return <div>対象の選手が存在しない</div>

  return (
    <ThemeProvider theme={theme}>
      <CustomCard>
        <CardContent>
          <Typography variant="h5" align="left">
            選手情報{' '}
          </Typography>
          <Typography align="left">名前: {player.name}</Typography>
          <Typography align="left">背番号: {player.uniformNumber}</Typography>
          <Typography align="left">打数: {player.atBats}</Typography>
          <Typography align="left">ヒット: {player.hits}</Typography>
          <Typography align="left">ホームラン: {player.homeRuns}</Typography>
          <Typography align="left">四球: {player.runsBattedIn}</Typography>
        </CardContent>
      </CustomCard>
    </ThemeProvider>
  )
}
