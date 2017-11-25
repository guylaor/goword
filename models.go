package goword

type WordDocument struct {
	Paragraphs []Paragraph
}

type Paragraph struct {
	Style Style `xml:"pPr>pStyle"`
	Rows  []Row `xml:"r"`
}

type Style struct {
	Val string `xml:"val,attr"`
}
type Row struct {
	Text string `xml:"t"`
}

// methods
func (w WordDocument) AsText() string {
	text := ""
	for _, v := range w.Paragraphs {
		for _, rv := range v.Rows {
			text += rv.Text
		}
		text += "\n"
	}
	return text
}
