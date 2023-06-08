import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'

import { PlayerIndexRoute, PlayerDetailRoute } from '../features/player/routes/Player'

function HomeRedirect() {
  return <Navigate to="/player/" replace />
}

export const RouterConfig = () => {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomeRedirect />} />
          <Route path="/player/" element={<PlayerIndexRoute />} />
          <Route path="/player/:id" element={<PlayerDetailRoute />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}
