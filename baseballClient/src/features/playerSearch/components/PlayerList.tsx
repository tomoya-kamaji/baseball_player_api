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

const headers = [
  {
    name: '名前',
    key: 'name'
  },
  {
    name: '背番号',
    key: 'uniformNumber'
  },
  {
    name: '打率',
    key: 'battingAverage'
  },
  {
    name: '打数',
    key: 'atBats'
  },
  {
    name: 'ヒット数',
    key: 'hits'
  },
  {
    name: '本塁打',
    key: 'homeRuns'
  },
  {
    name: '打点',
    key: 'runsBattedIn'
  },
  {
    name: '四球',
    key: 'walks'
  }
]

export const PlayerList = (): JSX.Element => {
  const { searchParam, setSearchParam } = useSearchParamStore()
  const { players, error } = useSearchPlayers(searchParam)

  const sort = (sortField: string) => {
    setSearchParam({
      sortField: sortField
    })
  }

  if (error) return <div>読み込みに失敗しました</div>
  return (
    <ThemeProvider theme={theme}>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              {headers.map((header, index) => (
                <TableCell align={index === 0 ? 'left' : 'right'} key={header.key} onClick={() => sort(header.key)}>
                  {header.name}
                </TableCell>
              ))}
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
