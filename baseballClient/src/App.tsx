/** @jsxImportSource @emotion/react */
import React from 'react'
// import { css } from '@emotion/react'
import { QueryClient, QueryClientProvider } from 'react-query'
import { RouterConfig } from './routes'

// const style = css``
const App: React.FC = () => {
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <RouterConfig />
    </QueryClientProvider>
  )
}

export default App
