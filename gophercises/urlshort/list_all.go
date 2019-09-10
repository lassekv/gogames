package urlshort

import (
	"fmt"
	"io"
	"text/template"

	"github.com/lassekv/gogames/gophercises/urlshort/dynamo"
)

type templateObject struct {
	Records []dynamo.Record
}

func listAll(records []dynamo.Record, io io.Writer) {
	t, err := template.ParseFiles("../template.html")
	if err != nil {
		fmt.Printf("Unable to parse template %v\n", err)
	}
	t.Execute(io, templateObject{Records: records})
}
