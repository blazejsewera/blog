import type { Component } from 'solid-js'
import { cx } from '../util/classes'
import { BIG_TITLE_FONT, HeadingProps } from './common'

export const H2: Component<HeadingProps> = ({ children }) => (
  <h2 class={cx(BIG_TITLE_FONT, 'text-2xl')}>{children}</h2>
)
