import type { ArticleInfo } from './metadata'

type Comparator<T> = (value1: T, value2: T) => number

export const byDateDesc: Comparator<Date> = (date1, date2) =>
  date2.getTime() >= date1.getTime() ? 1 : -1

export const articleInfoByDateDesc: Comparator<ArticleInfo> = (a1, a2) =>
  byDateDesc(a1.date, a2.date)
