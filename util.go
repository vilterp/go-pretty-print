package prettyprint

// TODO: make this split over line breaks if there's
// not enough width
// TODO: formatting with trailing commas would be nice
func Join(docs []Doc, sep Doc) Doc {
	var out []Doc
	for idx, doc := range docs {
		if idx > 0 {
			out = append(out, sep)
		}
		out = append(out, doc)
	}
	return Seq(out)
}

func Surround(before string, doc Doc, after string) Doc {
	return Seq([]Doc{Text(before), doc, Text(after)})
}

var Comma = Text(",")
var CommaSpace = Text(", ")

var CommaNewline = Seq([]Doc{Comma, Newline})

type Formatter interface {
	Format() Doc
}

func Block(start string, doc Doc, end string) Doc {
	return SeqV(
		Text(start), Newline,
		Indent(2, doc),
		Newline, Text(end),
	)
}

func SeqV(docs ...Doc) Doc {
	return Seq(docs)
}
