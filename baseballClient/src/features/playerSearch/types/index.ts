export type SearchParam = {
  limit?: number
  maxHits?: number
  maxHomeRuns?: number
  maxRunsBattedIn?: number
  minHits?: number
  minHomeRuns?: number
  minRunsBattedIn?: number
  sortField?: string
  sortOrder?: string
}

export type FetchPlayerApiModel = {
  player: PlayerModel | undefined
  error: Error | undefined
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
  // 打率
}
