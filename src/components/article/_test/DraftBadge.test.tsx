import { render } from '@solidjs/testing-library'
import { DraftBadge } from '../DraftBadge'

describe('<DraftBadge />', () => {
  const { container, unmount } = render(() => <DraftBadge />)

  expect(container).toMatchSnapshot()

  unmount()
})
