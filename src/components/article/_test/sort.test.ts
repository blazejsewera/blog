import { describe, it, expect } from 'vitest'
import type { ArticleInfo } from '../articleInfo'
import { articleInfoByDateDesc, byDateDesc } from '../sort'

describe('compare', () => {
  it('returns -1 if the first date is later', () => {
    // given
    const date1 = new Date('2000-01-02')
    const date2 = new Date('2000-01-01')

    // when
    const actual = byDateDesc(date1, date2)

    // then
    expect(actual).toEqual(-1)
  })

  it('returns 1 if the first date is sooner', () => {
    // given
    const date1 = new Date('2000-01-01')
    const date2 = new Date('2000-01-02')

    // when
    const actual = byDateDesc(date1, date2)

    // then
    expect(actual).toEqual(1)
  })
})

describe('sort array', () => {
  const unsorted = ['2000-01-01', '2000-01-03', '1999-12-31', '2000-01-02']
  const sortedDesc = ['2000-01-03', '2000-01-02', '2000-01-01', '1999-12-31']

  it('sorts dates descending', () => {
    const toDate = (it: string) => new Date(it)
    // given
    const unsortedDates = unsorted.slice().map(toDate)
    const expected = sortedDesc.slice().map(toDate)

    // when
    const actual = unsortedDates.slice().sort(byDateDesc)

    // then
    expect(actual).toEqual(expected)
  })

  it('sorts articles by date descending', () => {
    const toArticleInfo = (it: string): ArticleInfo => ({
      file: 'file',
      url: 'url',
      title: 'title',
      date: new Date(it),
    })

    // given
    const unsortedArticles = unsorted.slice().map(toArticleInfo)
    const expected = sortedDesc.slice().map(toArticleInfo)

    // when
    const actual = unsortedArticles.slice().sort(articleInfoByDateDesc)

    // then
    expect(actual).toEqual(expected)
  })
})
