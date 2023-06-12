import { IconButton } from '@mui/material'
import AppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'

export const Header = (): JSX.Element => {
  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" style={{ flexGrow: 1 }}>
          ヘッダータイトル
        </Typography>
      </Toolbar>
    </AppBar>
  )
}
