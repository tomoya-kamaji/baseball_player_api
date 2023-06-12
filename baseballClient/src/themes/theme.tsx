import { createTheme } from '@mui/material'

export const theme = createTheme({
  typography: {
    fontFamily: 'Roboto, sans-serif'
  },
  palette: {
    primary: {
      main: '#03734f' // プライマリカラーを青色に設定
    },
    secondary: {
      main: '#6e836e' // セカンダリカラーを緑色に設定
    },
    text: {
      primary: '#0a0a0a' // テキストの色を黒色に設定
    },
    background: {
      default: '#f5f5f5' // 背景色をグレーに設定
    }
  }
})
