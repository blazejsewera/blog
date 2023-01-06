import type { Component } from 'solid-js'
import { cx } from '../util/classes'

export const DraftBadge: Component = () => (
  <div
    class={cx([
      'inline-flex',
      'rounded-full',
      'border-2',
      'w-min',
      'h-min',
      'px-3',
      'py-0.5',
      'uppercase',
      'text-xs',
      'border-orange-600',
      'text-orange-600',
      'bg-orange-100',
    ])}
  >
    <div class="m-auto">draft</div>
  </div>
)
