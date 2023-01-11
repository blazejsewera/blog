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
        'min-w-full',
        'md:w-80',
        'shadow-xl',
        'break-inside-avoid',
      )}
      href={url}
    >
      <div
        class={cx(
          'bg-clip-border',
          'bg-origin-border',
          'bg-center',
          'bg-cover',
          'aspect-video',
          'flex',
        )}
        style={{ 'background-image': `url('${imgUrl}')` }}
        title={imgDesc}
      >
        <div class={cx('ml-4', 'mr-auto', 'mb-4', 'mt-auto')}>
          {draft && <DraftBadge />}
        </div>
      </div>
      <div class={cx('p-4')}>
        <div class={cx('text-xl', 'mb-3')}>{title}</div>
        <div class={cx('text-sm', 'text-neutral-400')}>
          {date.toLocaleDateString('en-US', { dateStyle: 'long' })}
        </div>
      </div>
    </a>
  )
}
