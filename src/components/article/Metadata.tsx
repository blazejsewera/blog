import type { Component } from 'solid-js'

export interface MetadataProps {
  title: string
  date: string
  author: string
  license: string
  language: string
  siteName: string
  abstract?: string
  imgUrl?: string
  imgDescription?: string
}

export const Metadata: Component<MetadataProps> = ({
  title,
  date,
  author,
  license,
  language,
  siteName,
  abstract,
  imgUrl,
  imgDescription,
}) => {
  return (
    <>
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

      {imgUrl && <meta name="og:image" content={imgUrl} />}

      {imgDescription && <meta name="og:image:alt" content={imgDescription} />}
    </>
  )
}