import {
  ThemeProvider,
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Link
} from '@mui/material'
import { theme } from '../../../themes/theme'
import { useSearchPlayers } from '../api/useSearchPlayers'
import { Link as RouterLink } from 'react-router-dom'
import { battingAverageToDisplay } from '../function'
import { useSearchParamStore } from '../store/useCountStare'

export const PlayerIndex = (): JSX.Element => {
  const { searchParam } = useSearchParamStore()
  console.log(searchParam)
  const { players, error } = useSearchPlayers(searchParam)

  if (error) return <div>読み込みに失敗しました</div>
  const headers = ['名前', '背番号', '打率', '打数', 'ヒット数', '本塁打', '打点', '四球']
  return (
    <ThemeProvider theme={theme}>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              {headers.map((header, index) => (
                <TableCell align={index === 0 ? 'left' : 'right'} key={header}>
                  {header}
                </TableCell>
              ))}{' '}
            </TableRow>
          </TableHead>
          <TableBody>
            {players.map((p) => (
              <TableRow key={p.id} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
                <TableCell component="th" scope="row">
                  <Link component={RouterLink} to={`/player/${p.id}`}>
                    {p.name}
                  </Link>
                </TableCell>
                <TableCell align="right">{p.uniformNumber}</TableCell>
                <TableCell align="right">{battingAverageToDisplay(p.hits, p.atBats)}</TableCell>
                <TableCell align="right">{p.atBats}</TableCell>
                <TableCell align="right">{p.hits}</TableCell>
                <TableCell align="right">{p.homeRuns}</TableCell>
                <TableCell align="right">{p.runsBattedIn}</TableCell>
                <TableCell align="right">{p.walks}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>{' '}
    </ThemeProvider>
  )
}
