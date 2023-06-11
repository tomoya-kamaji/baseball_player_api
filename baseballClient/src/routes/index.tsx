import { PlayerDetailRoute, PlayerSearchRoute } from '../features/playerSearch'
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'

function HomeRedirect() {
  return <Navigate to="/player/" replace />
}

export const RouterConfig = () => {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomeRedirect />} />
          <Route path="/player/" element={<PlayerSearchRoute />} />
          <Route path="/player/:id" element={<PlayerDetailRoute />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}
