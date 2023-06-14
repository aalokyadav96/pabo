package main

import (
	"net/http"
	)


type Result struct {
	FromLang string
	Original_Text string
	Translated_Text string
	ToLang string
}

func Translate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "itranslate.html", nil)
	} else {
		tt := r.FormValue("orig")
		var res Result
		res = Result{FromLang: "EN", Original_Text: r.FormValue("orig"), Translated_Text: tt, ToLang: "KO"}
		tmpl.ExecuteTemplate(w, "head.html", nil)
		tmpl.ExecuteTemplate(w, "translate.html", res)
		tmpl.ExecuteTemplate(w, "footer.html", nil)
	}
}



