import type { Component } from 'solid-js'
import { DraftBadge } from './DraftBadge'

export interface DraftInfoProps {
  description?: string
}

export const DraftInfo: Component<DraftInfoProps> = ({ description }) => (
  <div>
    <DraftBadge />
    {description && <div>{description}</div>}
  </div>
)
