import type {
  Update,
  Article,
  ArticleInfo,
  Frontmatter,
  FrontmatterUpdate,
} from './articleInfo'

export function toArticleInfo(article: Article): ArticleInfo {
  const stubArticleInfo: ArticleInfo = {
    url: '',
    file: '',
    title: 'ERROR, NO TITLE',
    date: new Date('2000-01-01'),
  }

  const metadata = article.frontmatter

  if (!metadata.title || !metadata.date) return stubArticleInfo
  const date = parseDate(metadata.date)
  const minImgUrl = getMinImgUrlOrFallbackToImgUrl(metadata)
  const updates = metadata.updates?.map(toArticleInfoUpdate)

  return {
    ...metadata,
    minImgUrl,
    date,
    updates,
    url: article.url,
    file: article.file,
  } as ArticleInfo
}

function getMinImgUrlOrFallbackToImgUrl(metadata: Frontmatter) {
  return metadata.minImgUrl === undefined || metadata.minImgUrl === ''
    ? metadata.imgUrl
    : metadata.minImgUrl
}

function toArticleInfoUpdate(update: FrontmatterUpdate): Update {
  return { diffUrl: update.diffUrl, date: parseDate(update.date) }
}

function parseDate(dateString: string): Date {
  return new Date(dateString)
}
