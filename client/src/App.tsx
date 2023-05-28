/** @jsxImportSource @emotion/react */
import React from 'react'
import { css } from '@emotion/react'
import Button from '@mui/material/Button'
import { PlayersTable } from './features/search/conponents/Table'
import { QueryClient, QueryClientProvider } from 'react-query'

const style = css`
  background-color: red;
  &:hover {
    background-color: red;
  }
`
const App: React.FC = () => {
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <Button variant="contained" css={style}>
        <PlayersTable />
      </Button>
    </QueryClientProvider>
  )
}

export default App
