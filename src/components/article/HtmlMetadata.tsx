import type { Component } from 'solid-js'

export interface MetadataProps {
  title: string
  date: string
  author: string
  license: string
  language: string
  siteName: string
  abstract?: string
  keywords?: string[]
  imgUrl?: string
  imgDescription?: string
}

const BasicMetadata: Component = () => (
  <>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1.0" />
    <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
    <link rel="stylesheet" href="/style/swisstypefaces.css" />
  </>
)

export const HtmlMetadata: Component<MetadataProps> = ({
  title,
  date,
  author,
  license,
  language,
  siteName,
  abstract,
  keywords,
  imgUrl,
  imgDescription,
}) => {
  return (
    <>
      <BasicMetadata />

      <meta name="og:type" content="article" />

      <meta name="dcterms.title" content={title} />
      <meta name="og:title" content={title} />

      <meta name="created" content={date} />
      <meta name="date" content={date} />
      <meta name="dcterms.date" content={date} />
      <meta name="article:published_time" content={date} />

      <meta name="author" content={author} />
      <meta name="dcterms.creator" content={author} />
      <meta name="article:author" content={author} />

      <meta name="dcterms.rights" content={`${license} ${author}`} />
      <meta name="dcterms.license" content={license} />

      <meta name="dcterms.language" content={language} />

      <meta name="og:site_name" content={siteName} />

      {abstract && (
        <>
          <meta name="description" content={abstract} />
          <meta name="dcterms.abstract" content={abstract} />
          <meta name="og:description" content={abstract} />
        </>
      )}

      {keywords && (
        <>
          <meta name="keywords" content={keywords.join(', ')} />
        </>
      )}

      {imgUrl && <meta name="og:image" content={imgUrl} />}

      {imgDescription && <meta name="og:image:alt" content={imgDescription} />}
    </>
  )
}
