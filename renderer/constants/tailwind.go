package constants

const (
	TailwindUpstreamURL  = "https://github.com/tailwindlabs/tailwindcss/releases/download"
	TailwindBaseFilename = "tailwindcss"
	TailwindBinary       = RendererPrefix + "bin/" + TailwindBaseFilename
	TailwindChecksum     = RendererPrefix + "bin/tailwindcss.checksum.txt"
	TailwindVersion      = "v3.3.5"
	TailwindConfigFile   = RendererPrefix + "preprocess/tailwind/tailwind.config.js"
	TailwindStyleFile    = SiteDir + "/style/tailwind.css"
)
