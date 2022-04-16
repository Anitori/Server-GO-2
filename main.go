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
	log.Println("Escuchando en puerto :8080")
	http.HandleFunc("/upload", upload)

	http.ListenAndServe(":8080", nil)
	log.Println("IP Entrante") 
}

func home (w http.ResponseWriter, r *http.Request){
	log.Println("Ip Entrante", r.RemoteAddr)
}

func upload(w http.ResponseWriter, r *http.Request){
	
	if r.Method == http.MethodPost {
		file, handle, err:= r.FormFile("myFile")
		template, err := template.ParseFiles("template/test.html")

		nombre := handle.Filename
		tamaño := handle.Size

		tamañoEnKb := tamaño / 1000
		tamañoEnMb := tamaño / 1000000

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
		fmt.Fprintf(w, "Archivo cargado exitosamente")
		
		log.Printf("Archivo cargado exitosamente: %v", nombre) //Filename previamente guardado en una variable
		log.Printf("IP proveniente: %v", r.RemoteAddr) //Dirección IP del que sube el archivo
		log.Printf("Tipo de archivo: %v", http.DetectContentType(data) ) 
		log.Printf("Tamaño en bytes: %d", tamaño) //SizeFile previamente guardado en una variable
		log.Printf("Tamaño en Kb: %d", tamañoEnKb)
		log.Printf("Tamaño en Mb: %v", tamañoEnMb)
		log.Println("") //Espacios para una mayor claridad a la hora de lectura
		log.Println("")

		

		
		

	

	}
}





