import { Metadata } from '../Metadata'
import { render } from '@solidjs/testing-library'

describe('<Metadata />', () => {
  it('renders properly', () => {
    const author = 'author'
    const date = '2022-01-06'
    const language = 'en-US'
    const license = 'CC BY-SA 3.0'
    const siteName = 'blog.sewera.dev'
    const title = 'title'

    const { container, unmount } = render(() => (
      <Metadata {...{ author, date, language, license, siteName, title }} />
    ))

    expect(container).toMatchSnapshot()

    unmount()
  })
})
