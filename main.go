package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){

	
	// http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Println("Escuchando en puerto :8080")

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) { //En vez de agregar la función en el segundo parámetro, había que crearla
		http.ServeFile(rw, req, "./public/index.html") //Testear en upload
		log.Println("Ip Entrante", req.RemoteAddr)
		log.Println("")
	})

	http.HandleFunc("/upload", upload)

	
	// log.Println("IP:")
	http.ListenAndServe(":8080", nil )

	 
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
		log.Println("Tamaño de ", tamaño, "Bytes") //SizeFile previamente guardado en una variable
		log.Println("Tamaño de ", tamañoEnKb, "KB")
		log.Println("Tamaño de ", tamañoEnMb, "MB")
		log.Println("") //Espacios para una mayor claridad de lectura
		log.Println("")
	}
}





