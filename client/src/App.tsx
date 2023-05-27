/** @jsxImportSource @emotion/react */
import React from 'react'
import { css } from '@emotion/react'
import Button from '@mui/material/Button'
import { PlayersTable } from './features/search/conponents/Table'
import fetchPlayers from './features/search/api/fetchPlayers'

const style = css`
  background-color: red;
  &:hover {
    background-color: red;
  }
`

const App: React.FC = () => {
  return (
    <Button variant="contained" css={style}>
      <PlayersTable />
    </Button>
  )
}

export default App
