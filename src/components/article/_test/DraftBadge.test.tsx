import { describe, expect, it } from 'vitest'
import { render } from '@solidjs/testing-library'
import { DraftBadge } from '../DraftBadge'

describe('<DraftBadge />', () => {
  it('renders properly', () => {
    const { container, unmount } = render(() => <DraftBadge />)
    expect(container).toMatchSnapshot()
    unmount()
  })
})
