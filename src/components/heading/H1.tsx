import type { Component } from 'solid-js'
import { cx } from '../util/classes'
import { BIG_TITLE_FONT, HeadingProps } from './common'

export const H1: Component<HeadingProps> = ({ children }) => (
  <h1 class={cx(BIG_TITLE_FONT, 'text-4xl')}>{children}</h1>
)
