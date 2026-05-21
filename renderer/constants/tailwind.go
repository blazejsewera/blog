package constants

const (
	TailwindUpstreamURL  = "https://github.com/tailwindlabs/tailwindcss/releases/download"
	TailwindBaseFilename = "tailwindcss"
	TailwindBinary       = RendererDir + "/bin/" + TailwindBaseFilename
	TailwindChecksum     = RendererDir + "/bin/tailwindcss.checksum.txt"
	TailwindVersion      = "v4.3.0"
	TailwindConfigFile   = RendererDir + "/preprocess/tailwind/tailwind.config.css"
	TailwindStyleFile    = SiteDir + "/style.css"
)
