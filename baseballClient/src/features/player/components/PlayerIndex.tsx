import {
  Typography,
  Card,
  CardContent,
  styled,
  ThemeProvider,
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody
} from '@mui/material'
import { theme } from '../../../themes/theme'
import { useSearchPlayers } from '../hooks/useSearchPlayers'
import { SearchParam } from '../types'

export const PlayerIndex = ({ param }: { param: SearchParam }): JSX.Element => {
  const { players } = useSearchPlayers(param)

  const headers = ['名前', '背番号', '打数', 'ヒット数', '本塁打', '打点', '四球']
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
              <TableRow key={p.name} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
                <TableCell component="th" scope="row">
                  {p.name}
                </TableCell>
                <TableCell align="right">{p.uniformNumber}</TableCell>
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
