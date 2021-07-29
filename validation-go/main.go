package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
    http.HandleFunc("/", getIndex)
    http.HandleFunc("/upload", getPost)

    fmt.Println("server started at localhost:9001")
    http.ListenAndServe(":9001", nil)
}

//get index.html
func getIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "403 Error", http.StatusBadRequest)
        return
    }

    var tmpl = template.Must(template.ParseFiles("index.html"))
    var err = tmpl.Execute(w, nil)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


//post form image
func getPost(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "", http.StatusBadRequest)
        return
    }

	const max_size = 8 * 1024 * 1024
	//check size > 8MB
	r.Body = http.MaxBytesReader(w, r.Body, max_size)
    if err := r.ParseMultipartForm(max_size); err != nil {
        http.Error(w, "Error 403 Max size image 8MB", http.StatusInternalServerError)
        return
    }

	//write filename & size
	file, handler, err := r.FormFile("data")
	if err != nil {
		http.Error(w, "Error 403 Insert a image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)

	// check filetype
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" { {
		http.Error(w, "Error 403 Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(fmt.Sprintf("temp-images/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Success")
}	