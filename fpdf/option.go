package fpdf

import (
	"io"

	"github.com/go-swiss/fonts"
)

type Config struct {
	Title   string
	Subject string

	Orientation string // Default Portrait
	PaperSize   string // Default A4

	Logo       io.Reader
	LogoFormat string

	// Header
	HeaderFunc func(impl Impl, fontsCache fonts.Cache) func()

	// Footer
	FooterFunc func(impl Impl, fontsCache fonts.Cache) func()
}
