export function battingAverageToDisplay(hits: number, atBats: number): string {
  const result = battingAverage(hits, atBats)
  // 0.333 -> .333
  return result.toFixed(3).substr(1)
}

export function battingAverage(hits: number, atBats: number): number {
  if (atBats === 0) return 0
  return hits / atBats
}

export function parseValue(type: string, value: string, checked: boolean) {
  if (type === 'number') {
    return Number(value)
  } else if (type === 'checkbox') {
    return checked
  } else {
    return value
  }
}
