import type { MarkdownInstance } from 'astro'
import { describe, expect, it } from 'vitest'
import { ArticleInfo, toArticleInfo } from '../articleInfo'

describe('ArticleInfo', () => {
  it('maps from Markdown frontmatter metadata', () => {
    const article = {
      file: 'file',
      url: 'url',
      frontmatter: {
        title: 'title',
        date: 'date',
        draft: true,
        draftDescription: 'draftDescription',
        author: 'author',
        abstract: 'abstract',
        language: 'language',
        license: 'license',
        imgUrl: 'imgUrl',
        imgDescription: 'imgDescription',
      },
    } as MarkdownInstance<ArticleInfo>

    const expected = {
      ...article.frontmatter,
      minImgUrl: 'imgUrl',
      url: 'url',
      file: 'file',
    } as ArticleInfo

    const actual = toArticleInfo(article)

    expect(actual).toStrictEqual(expected)
  })
})
