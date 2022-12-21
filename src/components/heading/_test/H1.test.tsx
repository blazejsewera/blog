import { H1 } from '../H1'
import { render } from '@solidjs/testing-library'

describe('<H1 />', () => {
  it('is accessible as first-level header', async () => {
    const { findByRole, unmount } = render(() => <H1>Test</H1>)
    const h1 = await findByRole('heading', { level: 1 })
    expect(h1).toBeInTheDocument()
    unmount()
  })

  it('has proper text', async () => {
    const expected = 'Test'
    const { findByText, unmount } = render(() => <H1>{expected}</H1>)
    const actual = await findByText(expected)
    expect(actual).toBeInTheDocument()
    unmount()
  })

  it('is properly styled', () => {
    const { container, unmount } = render(() => <H1>Test</H1>)
    expect(container).toMatchSnapshot()
    unmount()
  })
})
