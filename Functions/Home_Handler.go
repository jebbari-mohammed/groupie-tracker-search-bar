package functions

import (
	"html/template"
	"net/http"

	
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, 405)
		return
	}

	if r.URL.Path != "/" {
		ErrorHandler(w, r, 404)
		return
	}
	Error, err := Fitch_Global(w, r , Url_Artists) , Fitch_Global(w, r, Url_Locations)

	if Error != nil || err != nil {
		ErrorHandler(w,r,500)
		return
	}
	newlocation := []string{}
	founded := map[string]bool{}
	for _, location := range Locations.Index {
		for _, loc := range location.Locatins {
			if !founded[loc] {
			newlocation = append(newlocation, loc)
			founded[loc] = true
		}
	}
}

	tmpl, error := template.ParseFiles("Template/index.html")
	if error != nil {
		ErrorHandler(w, r, 500)
		return
	}
	data := page{
		Ar:  Artists,
		Loc: newlocation,
	}
	
	error =tmpl.Execute(w, data)
	if error!=nil{
		ErrorHandler(w,r,500)
		return
	}
}
