import type { Component } from 'solid-js'
import { cx } from '../util/classes'
import { DraftBadge } from './DraftBadge'

export type CardProps = {
  title: string
  date: Date
  url: string
  draft?: boolean
  minImgUrl?: string
  imgDescription?: string
}

export const Card: Component<CardProps> = ({
  title,
  date,
  url,
  draft,
  minImgUrl,
  imgDescription,
}) => {
  const defaultImgUrl = ''
  const imgUrl = minImgUrl ?? defaultImgUrl
  const defaultImgDescription = ''
  const imgDesc = imgDescription ?? defaultImgDescription

  return (
    <a
      class={cx(
        'block',
        'flex',
        'flex-col',
        'w-64',
        'rounded-lg',
        'shadow-lg',
        'break-inside-avoid',
      )}
      href={url}
    >
      <div
        class={cx('bg-clip-border', 'bg-cover', 'aspect-video')}
        style={{ 'background-image': `url('${imgUrl}')` }}
      />
      <div class={cx('p-2')}>
        <div>{draft && <DraftBadge />}</div>
        <div>{title}</div>
        <div>{date.toLocaleDateString('en-US', { dateStyle: 'long' })}</div>
      </div>
    </a>
  )
}
