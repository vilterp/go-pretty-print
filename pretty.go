package prettyprint

import (
	"bytes"
	"fmt"
	"strings"
)

// Based on http://homepages.inf.ed.ac.uk/wadler/papers/prettier/prettier.pdf
// TODO: this is the paper's naive implementation; use the more efficient one
// TODO: alternative layout combinators (best, group, etc)

type Doc interface {
	// Render returns the pretty-printed representation.
	String() string
	// String returns a representation of the doc tree, for debugging.
	Debug() string
}

// Text

// tried to alias this to string; didn't work
type text struct {
	str string
}

var _ Doc = &text{}

func Text(s string) *text {
	return &text{
		str: s,
	}
}

func Textf(format string, args ...interface{}) *text {
	return Text(fmt.Sprintf(format, args...))
}

func (s *text) String() string {
	return s.str
}

func (s *text) Debug() string {
	return fmt.Sprintf("Text(%#v)", s.str)
}

// Indent

type indent struct {
	doc      Doc
	indentBy int
}

func Indent(by int, d Doc) Doc {
	return &indent{
		doc:      d,
		indentBy: by,
	}
}

func (i *indent) String() string {
	indent := strings.Repeat(" ", i.indentBy)
	lines := strings.Split(i.doc.String(), "\n")
	buf := bytes.NewBufferString("")
	for idx, line := range lines {
		if idx > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(indent)
		buf.WriteString(line)
	}
	return buf.String()
}

func (i *indent) Debug() string {
	return fmt.Sprintf("Indent(%d, %s)", i.indentBy, i.doc.String())
}

// Empty

type empty struct{}

var Empty = &empty{}

func (e *empty) String() string {
	return ""
}

func (empty) Debug() string {
	return "Empty"
}

// Seq

type concat struct {
	docs []Doc
}

var _ Doc = &concat{}

func Seq(docs []Doc) Doc {
	return &concat{
		docs: docs,
	}
}

func (c *concat) String() string {
	buf := bytes.NewBufferString("")
	for _, doc := range c.docs {
		buf.WriteString(doc.String())
	}
	return buf.String()
}

func (c *concat) Debug() string {
	docStrs := make([]string, len(c.docs))
	for idx := range c.docs {
		docStrs[idx] = c.docs[idx].String()
	}
	return fmt.Sprintf("Seq(%s)", strings.Join(docStrs, ", "))
}

// Newline

type newline struct{}

var Newline = &newline{}

func (newline) String() string {
	return "\n"
}

func (newline) Debug() string {
	return "Newline"
}
