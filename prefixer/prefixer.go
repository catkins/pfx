package prefixer

import (
	"bufio"
	"fmt"
	"io"
	"text/template"
	"time"

	"github.com/fatih/color"
)

//TODO: use github.com/mitchellh/go-linereader to read from channel instead of bufio.Scanner
//TODO: support STDERR and separate template for error stream
//TODO: add simple mode where template is just program name with different colours for STDIN/STDERR

// Prefixer takes a stream of lines, and prefixes them with provided prefix
type Prefixer struct {
	prefix   string
	scanner  *bufio.Scanner
	output   io.Writer
	template *template.Template
}

// NewPrefixer builds a new prefixer
func NewPrefixer(prefix string, scanner *bufio.Scanner, output io.Writer) *Prefixer {
	p := new(Prefixer)
	p.scanner = scanner
	p.output = output
	p.template = parseTemplate(prefix)

	return p
}

type prefixData struct {
	Now     time.Time
	IsError bool
}

// PrefixLines output each line in the prefixer's scanner to the prefixer's writer
func (p *Prefixer) PrefixLines() {
	for p.scanner.Scan() {
		p.template.Execute(p.output, prefixData{Now: time.Now()})

		fmt.Fprintf(p.output, " %s\n", p.scanner.Text())
	}
}

func parseTemplate(prefix string) *template.Template {
	return template.Must(
		template.New("prefix-format").
			Funcs(templateFunctions()).
			Parse(prefix))
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"shorttime": shortTime,

		"blue":    color.BlueString,
		"black":   color.BlackString,
		"red":     color.RedString,
		"yellow":  color.YellowString,
		"magenta": color.MagentaString,
		"cyan":    color.CyanString,
		"green":   color.GreenString,
		"white":   color.WhiteString,

		"hiblue":    color.HiBlueString,
		"hiblack":   color.HiBlackString,
		"hired":     color.HiRedString,
		"hiyellow":  color.HiYellowString,
		"himagenta": color.HiMagentaString,
		"hicyan":    color.HiCyanString,
		"higreen":   color.HiGreenString,
		"hiwhite":   color.HiWhiteString}
}

func shortTime(t time.Time) (string, error) {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()), nil
}
