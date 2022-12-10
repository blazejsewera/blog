import type { Component } from 'solid-js'

type Level = 1 | 2 | 3 | 4 | 5 | 6

export interface HeadingProps {
  level: Level
  children: string | Element
}

export const Heading: Component<HeadingProps> = ({ level, children }) => {
  return <h1>{children}</h1>
}
