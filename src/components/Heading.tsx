import type { Component } from 'solid-js'
import { cx } from './util/classes'

type Level = 1 | 2 | 3 | 4 | 5 | 6

export interface HeadingProps {
  level: Level
  children: string | Element
}

export const Heading: Component<HeadingProps> = ({ level, children }) => {
  const commonStyles = ['']
  switch (level) {
    case 1:
      return <h1 class={cx(commonStyles, 'font-xl')}>{children}</h1>
    case 2:
      return <h2 class={cx(commonStyles, 'font-xl')}>{children}</h2>
    case 3:
      return <h3 class={cx(commonStyles, 'font-xl')}>{children}</h3>
    case 4:
      return <h4 class={cx(commonStyles, 'font-xl')}>{children}</h4>
    case 5:
      return <h5 class={cx(commonStyles, 'font-xl')}>{children}</h5>
    case 6:
      return <h6 class={cx(commonStyles, 'font-xl')}>{children}</h6>
  }
}
