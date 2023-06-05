import { Typography, Card, CardContent, styled, ThemeProvider } from '@mui/material'
import { useFetchPlayer } from '../api/fetchPlayer'
import { theme } from '../../../themes/theme'

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
  const { data, isLoading, isError } = useFetchPlayer(playerId)
  if (isLoading) return <div>loading...</div>
  if (isError) return <div>データの読み込みに失敗しました</div>
  if (data == undefined) {
    return <div>対象の選手が存在しない</div>
  }
  const name = data.player.name
  const uniformNumber = data.player.uniformNumber
  const atBats = data.player.atBats
  const hits = data.player.hits
  const homeRuns = data.player.homeRuns
  const runsBattedIn = data.player.runsBattedIn

  return (
    <ThemeProvider theme={theme}>
      <CustomCard>
        <CardContent>
          <Typography variant="h5" align="left">
            選手情報{' '}
          </Typography>
          <Typography align="left">名前: {name}</Typography>
          <Typography align="left">背番号: {uniformNumber}</Typography>
          <Typography align="left">打数: {atBats}</Typography>
          <Typography align="left">ヒット: {hits}</Typography>
          <Typography align="left">ホームラン: {homeRuns}</Typography>
          <Typography align="left">四球: {runsBattedIn}</Typography>
        </CardContent>
      </CustomCard>
    </ThemeProvider>
  )
}
