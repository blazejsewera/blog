package footnoteextension

import (
	"github.com/yuin/goldmark"
	goldmarkast "github.com/yuin/goldmark/ast"
	goldmarkextension "github.com/yuin/goldmark/extension"
	goldmarkextensionast "github.com/yuin/goldmark/extension/ast"
	goldmarkparser "github.com/yuin/goldmark/parser"
	goldmarkrenderer "github.com/yuin/goldmark/renderer"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	goldmarkutil "github.com/yuin/goldmark/util"
)

type footnote struct{}

func NewFootnote() goldmark.Extender {
	return &footnote{}
}

func (e *footnote) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		goldmarkparser.WithBlockParsers(
			goldmarkutil.Prioritized(goldmarkextension.NewFootnoteBlockParser(), 999),
		),
		goldmarkparser.WithInlineParsers(
			goldmarkutil.Prioritized(goldmarkextension.NewFootnoteParser(), 101),
		),
		goldmarkparser.WithASTTransformers(
			goldmarkutil.Prioritized(goldmarkextension.NewFootnoteASTTransformer(), 999),
		),
	)
	m.Renderer().AddOptions(goldmarkrenderer.WithNodeRenderers(
		goldmarkutil.Prioritized(NewFootnoteHTMLRenderer(), 500),
	))
}

type FootnoteHTMLRenderer struct {
	internal goldmarkrenderer.NodeRenderer
}

func NewFootnoteHTMLRenderer() *FootnoteHTMLRenderer {
	return &FootnoteHTMLRenderer{internal: goldmarkextension.NewFootnoteHTMLRenderer(
		goldmarkextension.WithFootnoteLinkClass([]byte("footnote-link")),
		goldmarkextension.WithFootnoteBacklinkClass([]byte("footnote-backlink")),
		goldmarkextension.WithFootnoteBacklinkHTML([]byte(backlinkIcon)),
	)}
}

// language=SVG
const backlinkIcon = `
<svg xmlns="http://www.w3.org/2000/svg" class="inline-block align-middle feather feather-corner-right-up" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#333333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
	<polyline points="10 9 15 4 20 9" />
	<path d="M4 20h7a4 4 0 0 0 4-4V4" />
</svg>
`

func (r *FootnoteHTMLRenderer) RegisterFuncs(reg goldmarkrenderer.NodeRendererFuncRegisterer) {
	r.internal.RegisterFuncs(reg)
	reg.Register(goldmarkextensionast.KindFootnoteList, r.renderFootnoteList)
}

func (r *FootnoteHTMLRenderer) renderFootnoteList(
	w goldmarkutil.BufWriter,
	_ []byte,
	node goldmarkast.Node,
	entering bool,
) (goldmarkast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString(`<div class="footnotes" role="doc-endnotes"`)
		if node.Attributes() != nil {
			goldmarkhtml.RenderAttributes(w, node, goldmarkhtml.GlobalAttributeFilter)
		}
		_ = w.WriteByte('>')
		_, _ = w.WriteString("\n<h2 id=\"footnote-label\">Footnotes</h2>\n")
		_, _ = w.WriteString("<ol>\n")
	} else {
		_, _ = w.WriteString("</ol>\n")
		_, _ = w.WriteString("</div>\n")
	}
	return goldmarkast.WalkContinue, nil
}
