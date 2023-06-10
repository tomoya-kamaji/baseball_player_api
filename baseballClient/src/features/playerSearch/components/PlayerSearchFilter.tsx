import { Button, FormControl, Grid, InputLabel, MenuItem, Select, TextField } from '@mui/material'
import { useSearchParamStore } from '../store/useCountStare'
// import { ChangeEvent, FormEvent } from 'react'
// import { SearchParam } from '../types'

export const PlayerSearchFilter = () => {
  const { searchParam, setSearchParam } = useSearchParamStore()

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    setSearchParam({ ...searchParam, limit: 10 })
  }

  const handleSelectChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
    const { name, value } = event.target
    setSearchParam({ ...searchParam, limit: 10 })
  }

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    // 検索処理を実行するための関数をここに追加
  }

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        label="Limit"
        name="limit"
        type="number"
        value={searchParam.limit || ''}
        onChange={handleInputChange}
      />
      <TextField
        label="Max Hits"
        name="maxHits"
        type="number"
        value={searchParam.maxHits || ''}
        onChange={handleInputChange}
      />
      <TextField
        label="Max Home Runs"
        name="maxHomeRuns"
        type="number"
        value={searchParam.maxHomeRuns || ''}
        onChange={handleInputChange}
      />
      <TextField
        label="Max Runs Batted In"
        name="maxRunsBattedIn"
        type="number"
        value={searchParam.maxRunsBattedIn || ''}
        onChange={handleInputChange}
      />
      <TextField
        label="Min Hits"
        name="minHits"
        type="number"
        value={searchParam.minHits || ''}
        onChange={handleInputChange}
      />
      <TextField
        label="Min Home Runs"
        name="minHomeRuns"
        type="number"
        value={searchParam.minHomeRuns || ''}
        onChange={handleInputChange}
      />
      <TextField
        label="Min Runs Batted In"
        name="minRunsBattedIn"
        type="number"
        value={searchParam.minRunsBattedIn || ''}
        onChange={handleInputChange}
      />
      <FormControl>
        <InputLabel>Sort Field</InputLabel>
        <Select name="sortField" value={searchParam.sortField || ''}>
          <MenuItem value="hits">Hits</MenuItem>
          <MenuItem value="home_runs">Home Runs</MenuItem>
          <MenuItem value="runs_batted_in">Runs Batted In</MenuItem>
        </Select>
      </FormControl>
      <FormControl>
        <InputLabel>Sort Order</InputLabel>
        <Select name="sortOrder" value={searchParam.sortOrder || ''}>
          <MenuItem value="asc">Ascending</MenuItem>
          <MenuItem value="desc">Descending</MenuItem>
        </Select>
      </FormControl>
      <Button variant="contained" type="submit">
        Search
      </Button>
    </form>
  )
}
