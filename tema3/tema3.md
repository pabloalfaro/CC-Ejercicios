# Autoevaluación semana 4

### Ejercicio 1. Buscar alguna demo interesante de Docker y ejecutarla localmente, o en su defecto, ejecutar la imagen anterior y ver cómo funciona y los procesos que se llevan a cabo la primera vez que se ejecuta y las siguientes ocasiones.

Lo primero que hago es instalar docker siguiendo los pasos de la siguiente [página](https://docs.docker.com/engine/install/ubuntu). Después lo configuro para poder usarlo sin privilegios de super usuario siguiendo [estos](https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user) pasos.

![Instalación](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%201/instalacion%20de%20docker.png)

![Super usuario](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%201/trabajar%20con%20docker%20sin%20sudo.png)

Como se ve en la anterior imagen, lo primero que hago es ejecutar holamundo. No encuentra la imagen local por lo que la busca de manera remota. Una vez que a he utilizado compruebo los contenedores y las imágenes.

![images ps](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%201/listado%20de%20dockers.png)

Tras haber hecho las pruebas anteriores ejecuto la imágen de los ejecicios con el siguiente resultado:

![pulpo](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%201/im%C3%A1gen%20pulpo.png)


### Ejercicio 2. Tomar algún programa simple, “Hola mundo” impreso desde el intérprete de línea de órdenes, y comparar el tamaño de las imágenes de diferentes sistemas operativos base, Fedora, CentOS y Alpine, por ejemplo.

Para hacer este ejericio primero creo el archivo [Dockerfile](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%202/Dockerfile) con la configuración.

`FROM golang:alpine`

`WORKDIR /Escritorio/Ejercicios/tema3/Capturas/Ejercicio\ 2`

`COPY ./holamundo.go ./`

`CMD ["go", "run", "holamundo.go"]`

Para construir y ejecutar la imágen utilizo los comandos que se ven en la imágen de abajo:

![build](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%202/holamundo_alpine.png)


### Ejercicio 3. Crear a partir del contenedor anterior una imagen persistente con commit.

![Commit](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%203/commit.png)

Con `docker ps -l` veo el listado de los contenedores. Para crear una imagen persistente utilizo `docker commit ID nombre_que_quiera`. Al terminar compruebo que se ha creado la nueva imágen.


### Ejercicio 4. Examinar la estructura de capas que se forma al crear imágenes nuevas a partir de contenedores que se hayan estado ejecutando.

Para hacer esto utilizo `docker history nombre_de_la_imágen`.

![history](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema3/Capturas/Ejercicio%204/history.png)
