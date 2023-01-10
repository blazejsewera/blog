import type { MarkdownInstance } from 'astro'
import { describe, expect, it } from 'vitest'
import { ArticleInfo, Frontmatter, toArticleInfo } from '../articleInfo'

describe('ArticleInfo', () => {
  it('maps from Markdown frontmatter metadata', () => {
    const article = {
      file: 'file',
      url: 'url',
      frontmatter: {
        title: 'title',
        date: '2000-01-01',
        draft: true,
        draftDescription: 'draftDescription',
        author: 'author',
        abstract: 'abstract',
        language: 'language',
        license: 'license',
        imgUrl: 'imgUrl',
        imgDescription: 'imgDescription',
      },
    } as MarkdownInstance<Frontmatter>

    const expected = {
      ...article.frontmatter,
      date: new Date('2000-01-01'),
      minImgUrl: 'imgUrl',
      url: 'url',
      file: 'file',
    } as ArticleInfo

    const actual = toArticleInfo(article)

    expect(actual).toEqual(expected)
  })
})
