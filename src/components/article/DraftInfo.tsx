import type { Component } from 'solid-js'
import { cx } from '../util/classes'
import { DraftBadge } from './DraftBadge'

export interface DraftInfoProps {
  description?: string
}

export const DraftInfo: Component<DraftInfoProps> = ({ description }) => (
  <div class={cx(['mx-16', 'mb-8', 'flex'])}>
    <DraftBadge />
    <span class={cx(['text-orange-600', 'inline', 'my-auto', 'ml-4'])}>
      {description ?? 'This is a draft article. It may be incomplete.'}
    </span>
  </div>
)
