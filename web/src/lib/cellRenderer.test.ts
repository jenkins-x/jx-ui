import { it, expect } from 'vitest'
import { HTMLInColumnCell } from './cellRenderer'

it('returns text with link if link is valid', () => {
  const output = HTMLInColumnCell('https://jenkins-x.io/', 'Jenkins X')
  expect(output).toContain('Jenkins X')
  expect(output).toContain('https://jenkins-x.io/')
})

it('returns only text with link if link is invalid', () => {
  const output = HTMLInColumnCell('', 'Jenkins X')
  expect(output).toContain('Jenkins X')
  // ToDo: test that it does not contain a <a> tag
  expect(output).not.toContain('http')
})
