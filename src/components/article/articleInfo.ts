import type { MarkdownInstance } from 'astro'

export type ArticleInfo = {
  title: string
  date: string
  url: string
  file: string
  draft?: boolean
  draftDescription?: string
  author?: string
  abstract?: string
  language?: string
  license?: string
  imgUrl?: string
  minImgUrl?: string
  imgDescription?: string
}

type Article = MarkdownInstance<Record<string, any>>

export const toArticleInfo = (article: Article): ArticleInfo => {
  if (!article.frontmatter.title || !article.frontmatter.date)
    return {
      url: '',
      file: '',
      title: 'ERROR, NO TITLE',
      date: '2000-01-01',
    }

  const minImgUrl =
    article.frontmatter.minImgUrl === undefined ||
    article.frontmatter.minImgUrl === ''
      ? article.frontmatter.imgUrl
      : article.frontmatter.minImgUrl

  return {
    ...article.frontmatter,
    minImgUrl,
    url: article.url,
    file: article.file,
  } as ArticleInfo
}
