import { it, expect } from 'vitest'
import { HTMLInColumnCell } from './cellRenderer'

it('returns text with link if link is valid', () => {
  const output = HTMLInColumnCell('https://jenkins-x.io/', 'Jenkins X')
  expect(output).toContain('Jenkins X')
  expect(output).toContain('https://jenkins-x.io/')
})
