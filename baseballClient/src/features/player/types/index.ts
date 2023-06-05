export type FetchPlayerApiModel = {
  player: PlayerModel
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
