import type { Component } from 'solid-js'
import { DraftBadge } from './DraftBadge'

export type CardProps = {
  title: string
  date: string
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
    <a href={url}>
      <div>
        <div>
          <img src={imgUrl} alt={imgDesc} />
        </div>
        <div>{draft && <DraftBadge />}</div>
        <div>{title}</div>
        <div>{date}</div>
      </div>
    </a>
  )
}
