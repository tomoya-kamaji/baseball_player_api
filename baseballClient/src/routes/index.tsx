import { BrowserRouter, Route, Routes } from 'react-router-dom'

import { Player } from '../features/player/routes/Player'

export const RouterConfig = () => {
  return (
    <>
      <BrowserRouter>
        <Routes>
          {/* <Route path="/" element={<Home />} /> */}
          <Route path="/player/:id" element={<Player />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}
