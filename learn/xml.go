package learn

import (
	"encoding/xml"
	"fmt"
)

type ProgLang struct {
	XMLName xml.Name `xml:"proglang"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p ProgLang) String() string {
	return fmt.Sprintf("Programming Id: %v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func XML() {
	python := &ProgLang{Id: 1, Name: "Python"}
	python.Origin = []string{"Netherlands", "Europe"}

	out, _ := xml.MarshalIndent(python, " ", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p ProgLang

	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	cpp := &ProgLang{Id: 25, Name: "C++"}
	cpp.Origin = []string{"Denmark", "Europe"}

	type Nesting struct {
		XMLName  xml.Name    `xml:"nesting"`
		ProgLang []*ProgLang `xml:"parent>child>proglang"`
	}

	nesting := &Nesting{}
	nesting.ProgLang = []*ProgLang{python, cpp}

	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
}
