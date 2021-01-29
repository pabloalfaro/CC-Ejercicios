package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	clienteHttp := &http.Client{}
	url := "http://172.17.0.1:3000/usuario"
	peticion, err := http.NewRequest("GET", url, nil)
	// Compruebo el error al crear la petición
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}
	
	respuesta, err := clienteHttp.Do(peticion)
	// Compruebo el error al hacer la petición
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}
	
	// Cierro el cuerpo
	defer respuesta.Body.Close()
	
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	// Compruebo el error al leer la respuesta
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	log.Printf("Código de respuesta: %d", respuesta.StatusCode)
	log.Printf("Encabezados: '%q'", respuesta.Header)
	contentType := respuesta.Header.Get("Content-Type")
	log.Printf("El tipo de contenido: '%s'", contentType)
	log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)

}
