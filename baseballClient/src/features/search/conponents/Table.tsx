import { Typography, Box } from '@mui/material'
import { useFetchPlayer } from '../api/fetchPlayers'

// const StyledTableCell = styled()
;`
  /* Emotionのスタイルを適用する */
`

export const PlayerDetails = (): JSX.Element => {
  const { data, isLoading, isError } = useFetchPlayer('6d89d4ec-0770-48f8-a6f2-882bb4b95fc9')
  if (isLoading) return <div>loading...</div>
  if (isError) return <div>データの読み込みに失敗しました</div>

  const name = data?.player.name || ''
  const uniformNumber = data?.player.uniformNumber || 0
  const atBats = data?.player.atBats || 0
  const hits = data?.player.hits || 0
  const homeRuns = data?.player.homeRuns || 0
  const runsBattedIn = data?.player.runsBattedIn || 0

  return (
    <Box>
      <Typography variant="h5" align="left">
        選手情報
      </Typography>
      <Typography align="left">Name: {name}</Typography>
      <Typography align="left">Uniform Number: {uniformNumber}</Typography>
      <Typography align="left">At Bats: {atBats}</Typography>
      <Typography align="left">Hits: {hits}</Typography>
      <Typography align="left">Home Runs: {homeRuns}</Typography>
      <Typography align="left">Runs Batted In: {runsBattedIn}</Typography>
    </Box>
  )
}
