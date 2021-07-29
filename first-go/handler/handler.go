package handler

import (
	"hello-world/entity"
	"html/template"
	"log"
	"net/http"
	"path"
)


func HandleIndex(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/"{
        http.NotFound(w, r)
        return
    }

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Home Page"))
}

func HandleAbout(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("About Page"))
}

func HandleProduct(w http.ResponseWriter, r *http.Request){
    
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	data := []entity.Product{
		{Id: 1, Name: "Chicken", Price: 13000, Stock: 2},
		{Id: 2, Name: "Noodle", Price: 15000, Stock: 3},
		{Id: 3, Name: "Vegetable", Price: 20000, Stock: 11},
		{Id: 4, Name: "Fried Rice", Price: 17000, Stock: 8},
	}

	err = tmpl.Execute(w, data)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	// json.NewEncoder(w).Encode(data)
}

func PostGet(w http.ResponseWriter, r *http.Request){
	method := r.Method // GET ATAU POST

	switch method{
		case "GET":
			w.Write([]byte("Ini adalah GET"))
		case "POST":
			w.Write([]byte("Ini adalah POST"))
		default:
			http.Error(w, "Error", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err!= nil{
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return	
		}
		err = tmpl.Execute(w, nil)
		if err != nil{
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return	
		}
	return
	}
	http.Error(w, "error", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()
		if err!= nil{
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return	
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name": name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err!= nil{
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return	
		}
		err = tmpl.Execute(w, data)
		if err != nil{
			log.Println(err)
			http.Error(w, "Error", http.StatusInternalServerError)
			return	
		}
		return
	}
	http.Error(w, "error", http.StatusBadRequest)
}