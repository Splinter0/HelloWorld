package ext

import (
	"encoding/xml"
)

//Extraction
type Query struct {
	XMLName xml.Name `xml:"urlset"`
	Pars    []string `xml:"url>p"`
	Chapts  []string `xml:"url>h1"`
}

//type Par string
//type Chapt string

//Objects
type Paragraph struct {
	Text string
}
type Chapter struct {
	Text       string
	Paragraphs []Paragraph
}

func conChapt(t string) *Chapter {
	c := &Chapter{
		Text: t,
	}
	return c
}

func conPar(t string) *Paragraph {
	p := &Paragraph{
		Text: t,
	}
	return p
}

func (c *Chapter) add(p *Paragraph) {
	c.Paragraphs = append(c.Paragraphs, *p)
}

/*
func main() {
	xmlFile, err := os.Open("bible.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(b, &q)
	/*
		for i := range q.Pars {
			fmt.Println(q.Pars[i])
		}
		fmt.Println(q.Chapts)

	c := conChapt(q.Chapts[0])
	p := conPar(q.Pars[0])
	c.add(p)
	fmt.Println(c.Text)
	fmt.Println(c.Paragraphs[0].Text)
}
*/
