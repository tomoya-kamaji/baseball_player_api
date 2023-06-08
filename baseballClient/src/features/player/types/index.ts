export type SearchParam = {
  limit?: number
  maxHits?: number
  maxHomeRuns?: number
  maxRunsBattedIn?: number
  minHits?: number
  minHomeRuns?: number
  minRunsBattedIn?: number
  sortField?: 'hits' | 'home_runs' | 'runs_batted_in'
  sortOrder?: 'asc' | 'desc'
}

export type FetchPlayerApiModel = {
  player: PlayerModel
}

export type SearchPlayerApiModel = {
  players: PlayerModel[]
  error: Error | undefined
}

export type PlayerModel = {
  id: string
  name: string
  uniformNumber: number
  atBats: number
  hits: number
  homeRuns: number
  runsBattedIn: number
  walks: number
}
