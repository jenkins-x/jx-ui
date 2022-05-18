import { render } from '@testing-library/svelte'
import Card from './Card.svelte'

it('Renders a card with the right title', () => {
  const { getByText } = render(Card, { title: 'Pipelines' })
  expect(getByText('Pipelines')).toBeInTheDocument()
})
