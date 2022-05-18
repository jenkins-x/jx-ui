import { isValid, formatDistanceStrict, format } from 'date-fns'

export function diffTimes(i: number, j: number): string {
  if (isValid(i) && isValid(j)) {
    return formatDistanceStrict(j, i)
  }
  return ''
}

export function displayTime(date: number, displayFormat: string = 'Pp'): string {
  if (isValid(date)) {
    return format(date, displayFormat)
  }
  return ''
}
