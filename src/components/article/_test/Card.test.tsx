import { render } from '@solidjs/testing-library'
import { describe, expect, it } from 'vitest'
import { Card } from '../Card'

describe('<Card />', () => {
  it('renders properly', async () => {
    const title = 'title'
    const date = new Date('2000-01-01')
    const url = 'url'

    const { unmount, findByText } = render(() => (
      <Card {...{ title, date, url }} />
    ))

    const renderedDate = await findByText('January 1, 2000')
    expect(renderedDate).toBeVisible()

    unmount()
  })
})
