import type { MarkdownInstance } from 'astro'

export type Frontmatter = {
  title: string
  date: string
  url: string
  file: string
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

export type ArticleInfo = {
  title: string
  date: Date
  url: string
  file: string
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

export type Article = MarkdownInstance<Frontmatter>

export const toArticleInfo = (article: Article): ArticleInfo => {
  if (!article.frontmatter.title || !article.frontmatter.date)
    return {
      url: '',
      file: '',
      title: 'ERROR, NO TITLE',
      date: new Date('2000-01-01'),
    }

  const date = new Date(article.frontmatter.date)

  const minImgUrl =
    article.frontmatter.minImgUrl === undefined ||
    article.frontmatter.minImgUrl === ''
      ? article.frontmatter.imgUrl
      : article.frontmatter.minImgUrl

  return {
    ...article.frontmatter,
    minImgUrl,
    date,
    url: article.url,
    file: article.file,
  } as ArticleInfo
}
