package constants

const (
	RendererPrefix = "renderer/"
	DistDir        = "dist"
	SiteDir        = "_site"
	PageDir        = RendererPrefix + "page"

	TailwindUpstreamURL = "https://github.com/tailwindlabs/tailwindcss/releases/download"
	TailwindBinary      = RendererPrefix + "bin/tailwindcss"
	TailwindChecksum    = RendererPrefix + "bin/tailwindcss.checksum.txt"
	TailwindVersion     = "v3.3.5"
	TailwindConfigFile  = RendererPrefix + "preprocess/tailwind/tailwind.config.js"
	TailwindStyleFile   = SiteDir + "/style/tailwind.css"

	FontUpstreamURL = "https://www.sewera.dev/swty/"
	FontDir         = SiteDir + "/font"
	FontList        = "font.list"

	MdExt   = ".md"
	HtmlExt = ".html"
)
