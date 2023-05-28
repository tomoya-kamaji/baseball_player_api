/** @jsxImportSource @emotion/react */
import React from 'react'
import { css } from '@emotion/react'
import { PlayerDetails } from './features/search/conponents/Table'
import { QueryClient, QueryClientProvider } from 'react-query'

// const style = css``
const App: React.FC = () => {
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <PlayerDetails />
    </QueryClientProvider>
  )
}

export default App
