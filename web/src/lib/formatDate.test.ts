import { it, expect } from 'vitest'
import { diffTimes, displayTime } from './formatDate'

it('returns difference in time between two valid dates', () => {
  expect(diffTimes(1, 2)).not.toBe('')
})

it('returns empty string if the dates are not valid', () => {
  expect(diffTimes(1, NaN)).toBe('')
})

it('displays valid time according to the format specified', () => {
  expect(displayTime(1)).not.toBe('')
})

it('displays empty string if string is not valid', () => {
  expect(displayTime(NaN)).toBe('')
})
