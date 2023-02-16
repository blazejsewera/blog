import type { MarkdownInstance } from 'astro'
import type { Article } from './articleInfo'
import { toArticleInfo } from './metadata'
import { articleInfoByDateDesc } from './sort'

export const getReverseChronologicalArticlesFromGlob = (
  globResult: MarkdownInstance<Record<string, any>>[],
) =>
  (globResult as unknown as Article[])
    .filter(({ frontmatter }) => !!frontmatter.title && !!frontmatter.date)
    .map(toArticleInfo)
    .slice()
    .sort(articleInfoByDateDesc)
