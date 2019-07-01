package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

/*

Taken from here:

https://www.calhoun.io/intro-to-templates-p3-functions/

Example usage:

{{htmlSafe "<!--[if IE 6]>"}}
<meta http-equiv="Content-Type" content="text/html; charset=Unicode">
{{htmlSafe "<![endif]-->"}}


Or

{{ifIE}}
<meta http-equiv="Content-Type" content="text/html; charset=Unicode">
{{endif}}

*/

func main() {
	var err error
	testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
		"ifIE": func() template.HTML {
			return template.HTML("<!--[if IE]>")
		},
		"endif": func() template.HTML {
			return template.HTML("<![endif]-->")
		},
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := testTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
