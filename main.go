package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
)

func main(){
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}

func upload(w http.ResponseWriter, r *http.Request){
	
	if r.Method == http.MethodPost {
		file, handle, err:= r.FormFile("myFile") //Ver esta linea
		template, err := template.ParseFiles("template/test.html")
		if err != nil {
			log.Printf("Error al cargar el archivo %v", err)
			fmt.Fprintf(w, "Error al cargar el archivo %v", err )
			return
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Error al leer el archivo %v", err)
			fmt.Fprintf(w, "Error al leer el archivo %v", err )
			return
		}

		err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666) // 0666 Permisos de Lectura y Escritura
		if err != nil {
			log.Printf("Error al escribir el archivo %v", err)
			fmt.Fprintf(w, "Error al escribir el archivo %v", err )
			return
		}
		template.Execute(w, nil)
		fmt.Fprint(w, "Archivo cargado exitosamente")
	}
}





