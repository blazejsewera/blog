import { render } from '@solidjs/testing-library'
import { DraftInfo } from '../DraftInfo'

describe('<DraftInfo />', () => {
  it('renders properly without description', () => {
    const { container, unmount } = render(() => <DraftInfo />)
    expect(container).toMatchSnapshot()
    unmount()
  })

  it('renders properly with description', () => {
    const description = 'description'
    const { container, unmount } = render(() => (
      <DraftInfo description={description} />
    ))
    expect(container).toMatchSnapshot()
    unmount()
  })
})
