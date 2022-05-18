import { isValid, formatDistanceStrict, format } from 'date-fns'

export function diffTimes(i: number, j: number): string {
  if (isValid(i) && isValid(j)) {
    return formatDistanceStrict(j, i)
  }
  return ''
}

export function displayTime(i: number): string {
  if (isValid(i)) {
    return format(i, 'Pp')
  }
  return ''
}
