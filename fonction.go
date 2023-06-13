package main

import (
	"html/template"
	"net/http"
	"os"
	"web/etat"
)

type Data struct {
	Output string
}

// creation d un serveur avec le port 8080
const port = ":8080"

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/404.html")
		template.Execute(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		template, _ := template.ParseFiles("./templates/405.html")
		template.Execute(w, r)
		return
	}
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Error while parsing the HTML file", http.StatusInternalServerError)
		template, _ := template.ParseFiles("./templates/500.html")
		template.Execute(w, r)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error while parsing the HTML file", http.StatusInternalServerError)
		template, _ := template.ParseFiles("./templates/500.html")
		template.Execute(w, r)
		return
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	da := ""
	if r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		template, _ := template.ParseFiles("./templates/404.html")
		template.Execute(w, r)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		template, _ := template.ParseFiles("./templates/405.html")
		template.Execute(w, r)
		return
	}

	testTemplet, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Error while parsing the HTML file", http.StatusInternalServerError)
		template, _ := template.ParseFiles("./templates/500.html")
		template.Execute(w, r)
		return
	}

	if r.Method == http.MethodPost {
		d := r.FormValue("entree")
		banner := r.FormValue("web")
		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			w.WriteHeader(http.StatusInternalServerError)
			template, _ := template.ParseFiles("./templates/500.html")
			template.Execute(w, r)
			return
		}
		file, _ := os.ReadFile(banner + ".txt")
		texte := string(file)
		if (len(texte) != 6623 && banner == "standard") || (len(texte) != 7463 && banner == "shadow") || (len(texte) != 5558 && banner == "thinkertoy") {
			w.WriteHeader(http.StatusInternalServerError)
			template, _ := template.ParseFiles("./templates/500.html")
			template.Execute(w, r)
			return
		}
		// Le texte ne contient pas de caractères spéciaux
		// http.Error(w, "pas de caractere special", http.StatusInternalServerError)
		if !Special(d) {
			w.WriteHeader(http.StatusBadRequest)
			template, _ := template.ParseFiles("./templates/400.html")
			template.Execute(w, r)
			return
		}
		// si la longueurscii a de la chaine est vide
		// http.Error(w, "", http.StatusInternalServerError)
		if len(d) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			template, _ := template.ParseFiles("./templates/400.html")
			template.Execute(w, r)
			return
		}

		da, err = etat.Ascifs(d, banner)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			template, _ := template.ParseFiles("./templates/500.html")
			template.Execute(w, r)
			return
		}
		data.Output = da
		testTemplet.Execute(w, data)
	}
}

func Special(s string) bool {
	for _, r := range s {
		if r < ' ' && r != '\n' && r != '\r' || r > '~' {
			return false
		}
	}
	return true
}
