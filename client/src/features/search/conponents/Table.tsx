import { styled } from '@mui/system'
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from '@mui/material'
import { useFetchPlayer } from '../api/fetchPlayers'

const StyledTableCell = styled(TableCell)`
  /* Emotionのスタイルを適用する */
`

export const PlayersTable = (): JSX.Element => {
  const { data, isLoading, isError, error } = useFetchPlayer('1')
  console.log(data)
  console.log(isLoading)
  console.log(isError)
  console.log(error)

  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <StyledTableCell></StyledTableCell>
            <StyledTableCell>Column 2</StyledTableCell>
            <StyledTableCell>Column 3</StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          <TableRow>
            <StyledTableCell>Value 1</StyledTableCell>
            <StyledTableCell>Value 2</StyledTableCell>
            <StyledTableCell>Value 3</StyledTableCell>
          </TableRow>
          <TableRow>
            <StyledTableCell>Value 4</StyledTableCell>
            <StyledTableCell>Value 5</StyledTableCell>
            <StyledTableCell>Value 6</StyledTableCell>
          </TableRow>
          {/* 追加の行 */}
        </TableBody>
      </Table>
    </TableContainer>
  )
}
