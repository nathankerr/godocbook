package main

import (
	"bytes"
	"github.com/nathankerr/godocbook/Godeps/_workspace/src/github.com/GeertJohan/go.rice"
	"github.com/nathankerr/godocbook/Godeps/_workspace/src/github.com/docopt/docopt-go"
	"github.com/nathankerr/godocbook/Godeps/_workspace/src/github.com/russross/blackfriday"
	"io"
	"log"
	"os"
	"runtime"
	"text/template"

	"github.com/nathankerr/godocbook/Godeps/_workspace/src/code.google.com/p/go.tools/godoc"
	"github.com/nathankerr/godocbook/Godeps/_workspace/src/code.google.com/p/go.tools/godoc/vfs"
	"github.com/nathankerr/godocbook/Godeps/_workspace/src/code.google.com/p/go.tools/godoc/vfs/gatefs"
)

func main() {
	log.SetFlags(log.Lshortfile)

	usage := `Create printable Go package documentation

Usage:
	godocbook [options] <package>...

Options:
`

	arguments, err := docopt.Parse(usage, nil, true, "godocbook", false)
	if err != nil {
		log.Fatalln(err)
	}

	doc := newDoc()

	packages := arguments["<package>"].([]string)
	package_bytes := make([]string, 0, len(packages))
	for _, pkg := range packages {
		md := &bytes.Buffer{}
		doc.For(md, pkg)

		package_bytes = append(package_bytes, string(md2Tex(md.Bytes())))
		// println(string(tex))
	}

	writeBookTex("book.tex", package_bytes)
}

type doc struct {
	pres *godoc.Presentation
	fs   vfs.NameSpace
}

func newDoc() *doc {
	d := doc{}

	// build an fs including GOROOT and GOPATH
	d.fs = vfs.NameSpace{}
	fsGate := make(chan bool, 20)
	d.fs.Bind("/", gatefs.New(vfs.OS(runtime.GOROOT()), fsGate), "/", vfs.BindReplace)

	box, err := rice.FindBox("templates")
	if err != nil {
		log.Fatalln(err)
	}

	// create a presentation based on the fs
	corpus := godoc.NewCorpus(d.fs)
	d.pres = godoc.NewPresentation(corpus)
	packageText, err := box.String("package.txt")
	if err != nil {
		log.Fatalln(err)
	}
	d.pres.PackageText, err = template.New("package.txt").Funcs(d.pres.FuncMap()).Delims("[", "]").Parse(packageText)
	if err != nil {
		log.Fatalln(err)
	}

	return &d
}

func (d *doc) For(w io.Writer, packageName string) {
	err := godoc.CommandLine(w, d.fs, d.pres, []string{packageName})
	if err != nil {
		log.Fatalln(err)
	}
}

// customized latex renderer
type renderer struct {
	blackfriday.Renderer
}

func newRenderer(latexFlags int) blackfriday.Renderer {
	return renderer{blackfriday.LatexRenderer(latexFlags)}
}

func (r renderer) DocumentHeader(out *bytes.Buffer) {
	// no-op
}

func (r renderer) DocumentFooter(out *bytes.Buffer) {
	// no-op
}

func (r renderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()
	switch level - 1 {
	case 0:
		out.WriteString("\n\\chapter{")
	case 1:
		out.WriteString("\n\\section{")
	case 2:
		out.WriteString("\n\\subsection{")
	case 3:
		out.WriteString("\n\\subsubsection{")
	case 4:
		out.WriteString("\n\\paragraph{")
	case 5:
		out.WriteString("\n\\subparagraph{")
	case 6:
		out.WriteString("\n\\textbf{")
	}
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("}\n")
}

func (options renderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	if lang == "" {
		out.WriteString("\n\\begin{lstlisting}\n")
	} else {
		out.WriteString("\n\\begin{lstlisting}[language=")
		out.WriteString(lang)
		out.WriteString("]\n")
	}
	out.Write(text)
	out.WriteString("\n\\end{lstlisting}\n")
}

func md2Tex(input []byte) []byte {
	latexFlags := 0
	renderer := newRenderer(latexFlags)

	extensions := 0

	return blackfriday.Markdown(input, renderer, extensions)
}

func writeBookTex(filename string, packages []string) {
	box, err := rice.FindBox("templates")
	if err != nil {
		log.Fatalln(err)
	}

	bookTex, err := box.String("book.tex")
	if err != nil {
		log.Fatalln(err)
	}
	bookTemplate, err := template.New("book.tex").Delims("[[", "]]").Parse(bookTex)
	if err != nil {
		log.Fatalln(err)
	}

	err = bookTemplate.Execute(os.Stdout, packages)
	if err != nil {
		log.Fatalln(err)
	}
}
