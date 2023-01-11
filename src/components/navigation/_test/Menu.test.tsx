import { describe, expect, it } from 'vitest'
import { render } from '@solidjs/testing-library'
import { Menu } from '../Menu'

describe('<Menu />', () => {
  async function testVisibility(visible: boolean, visibilityClass: string) {
    const baseSite = 'https://www.sewera.dev'
    const blogSite = 'https://blog.sewera.dev'
    const github = 'https://github.com/blazejsewera'
    const { findByRole, unmount } = render(() => (
      <Menu {...{ visible, baseSite, blogSite, github }} />
    ))

    const menu = await findByRole('navigation')
    expect(menu).toHaveClass(visibilityClass)
    unmount()
  }

  it('is not visible with visible = false', async () =>
    testVisibility(false, '-right-96'))

  it('is visible with visible = true', async () =>
    testVisibility(true, 'right-0'))
})
