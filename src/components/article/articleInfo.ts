import type { MarkdownInstance } from 'astro'

type WithFrontmatterDate<T> = Omit<T, 'date'> & { date: string }

export type Update = {
  date: Date
  diffUrl: string
}

export type FrontmatterUpdate = WithFrontmatterDate<Update>

export type ArticleInfo = {
  title: string
  date: Date
  url: string
  file: string
  subtitle?: string
  updates?: Update[]
  draft?: boolean
  draftDescription?: string
  author?: string
  abstract?: string
  keywords?: string[]
  language?: string
  license?: string
  imgUrl?: string
  minImgUrl?: string
  imgDescription?: string
}

export type Frontmatter = Omit<WithFrontmatterDate<ArticleInfo>, 'updates'> & {
  updates?: FrontmatterUpdate[]
}

export type Article = MarkdownInstance<Frontmatter>
