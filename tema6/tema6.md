### 1. Crear un pod con dos o más contenedores, de forma que se pueda usar uno desde el otro. Uno de los contenedores contendrá la aplicación que queramos desplegar.

Para este ejercicio he buscado información en algunos post y he llegado a [este](https://www.redhat.com/sysadmin/compose-podman-pods) de RedHat en el que muestra cómo hacer 
una composición de dos contenedores. Para empezar descargué podman siguiendo los pasos que se indicaban en el enlace de los ejercicios. El primer paso he sido crear el pod de
la siguiente manera:

![pod_create](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio1/podman_create.png)

Teniendo este pod, he creado los dos contenedores que indicaban en el ejemplo. El primero es una base de datos y el segundo es WordPress:

![mariadb](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio1/mariadb.png)

![wordpress](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio1/wordpress.png)

Para ver el pod y los contedores creados he usado las siguientes instrucciones:

![ps_ls](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio1/ls_ps.png)

De esta manera ya estaría configurada la composición. Se puede ver WordPress ejecutándose en localhost:8080.

![local](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio1/local.png)

### 2. Usar un miniframework REST para crear un servicio web y introducirlo en un contenedor, y componerlo con un cliente REST que sea el que finalmente se ejecuta y sirve como “frontend”.

Para este ejericio he buscado información para diferentes aspectos, cómo hacer [peticiones](https://parzibyte.me/blog/2019/05/21/peticion-post-get-put-delete-go-net-http/) desde
el cliente, cómo construir el [docker](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e) y cómo escribir el [docker-compose](https://docs.docker.com/compose/).

Lo primero que he hecho ha sido programar el servidor y el cliente. El servidor lo he sacado de mi proyecto, ha sido una adaptación simple para el ejericio. El 
[servidor](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio2/Servidor/servidor.go) ha sido el siguiente:

~~~
package main

import "github.com/kataras/iris/v12"

func newApp() *iris.Application {
    	app := iris.New()
	
	usuario := app.Party("/usuario")
	{
		usuario.Get("", buscarUsuario)
	}
	
   	return app
}

func buscarUsuario(ctx iris.Context){
	ctx.JSON(iris.Map{
		"username": "PabloA",
		"nombre": "Pablo",
		"apellido": "Alfaro",
		"correo": "pabloalfaro@correo.ugr.es",
		"ciudad": "Pamplona",
	})
	return
}

func main() {
	app := newApp()
	app.Listen(":3000")
}
~~~

Tiene programado un GET en el puerto 3000. El [cliente](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio2/Cliente/cliente.go) hace esa petición 
y muestra la respuesta:

~~~
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
~~~

Teniendo ya ambos programados he pasado a programar los docker. Este sería el del [servidor](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio2/Servidor/Dockerfile):

~~~
FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o servidor .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/servidor .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/servidor"]
~~~

El del [cliente](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio2/Cliente/Dockerfile) es este otro:

~~~

FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
#COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o cliente .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/cliente .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/cliente"]
~~~

Para construir los docker lo he hecho de esta manera para el servidor:

`docker build . -t servidor`

El del cliente es similar. El siguiente paso ha sido crear la [composición](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio2/docker-compose.yml).
Mi docker-compose ha sido este:

~~~
version: '3.7'

services:
  cliente:
    image: cliente
    ports:
      - 8080:8080
    depends_on:
      - servidor

  servidor:
    image: servidor
    ports:
      - 3000:3000
~~~

Ejecutando docker-compose up se puede ver el siguiente resultado:

![dockerCompose_Up](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema6/Capturas/Ejercicio2/dockerComose_up.png)

Se puede ver que se contectan y el cliente obtiene la respuesta después de hacer la petición.
