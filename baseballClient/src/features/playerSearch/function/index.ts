export function battingAverageToDisplay(hits: number, atBats: number): string {
  const result = battingAverage(hits, atBats)
  // 0.333 -> .333
  return result.toFixed(3).substr(1)
}

export function battingAverage(hits: number, atBats: number): number {
  if (atBats === 0) return 0
  return hits / atBats
}
