# Autoevaluación semana 2

### Ejercicio 1. Instalar alguno de los entornos virtuales de node.js (o de cualquier otro lenguaje con el que se esté familiarizado) y, con ellos, instalar la última versión existente, la versión minor más actual de la 4.x y lo mismo para la 0.11 o alguna impar (de desarrollo).

Lo primero que he hecho ha sido instalar nodeenv de la siguiente manera:
`pip3 install nodeenv`
![Instalación nodeenv](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio1/instalacion%20nodeenv.png)

Con nodeenv ya instalado he creado un entorno nuevo, lo he llamado env:
`nodeenv env`
![nodeenv entorno](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio1/instalacion%20del%20entorno.png)

Antes de elegir las versiones que quiero instalar saco un listado de las disponibles:
`nodeenv -l`
![nodeenv listado](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio1/listado%20versiones.png)

Me decido a instalar las versiones 4.9.1 y la 0.11.16. Para hacerlo utilizo el siguente comando:
`nodeenv --node=version nombre`
![nodeenv versiones](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio1/instalacion%20versiones.png)




### Ejercicio 2. Crear una descripción del módulo usando package.json. En caso de que se trate de otro lenguaje, usar el método correspondiente.

Para hacer este ejercicio he usado npm. En primer lugar lo he descargado, con npm listo he ido a la carpeta del pruebas del proyecto y he usado el comando init para crear el fichero package.json:	
`npm init`
![npm init](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio2/Creaci%C3%B3n%20json.png)


	
### Ejercicio 3. Descargar el repositorio de ejemplo anterior, instalar las herramientas necesarias (principalmente Scala y sbt) y ejecutar el ejemplo desde sbt. Alternativamente, buscar otros marcos para REST en Scala tales como Finatra o Scalatra y probar los ejemplos que se incluyan en el repositorio.

Antes de empezar he tenido que descargar las versiones 8 de java y la 2.11 de scala.

Luego he seguido los pasos de esta web para instalar sbt:  https://www.scala-sbt.org/1.x/docs/es/Installing-sbt-on-Linux.html

En mi caso uso una distribución de Ubuntu, los pasos que he tenido que seguir son los siguientes:

`echo "deb https://dl.bintray.com/sbt/debian /" | sudo tee -a /etc/apt/sources.list.d/sbt.list`

`curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | sudo apt-key add`

`sudo apt-get update`

`sudo apt-get install sbt`

Ya con todo instalado he clonado el repositorio https://github.com/JJ/spray-test, he iniciado sbt, he hecho `test` y finalmente `re-start`.

![clone](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio3/clone.png)

![test](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio3/test.png)

![re-start](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio3/re-start.png)

Tras esto, me he dispuesto a realizar las distintas pruebas que se indicaban en la información del repositorio.

![pruebas](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio3/pruebas.png)


### Ejercicio 4. Para la aplicación que se está haciendo, escribir una serie de aserciones y probar que efectivamente no fallan. Añadir tests para una nueva funcionalidad, probar que falla y escribir el código para que no lo haga. A continuación, ejecutarlos desde mocha (u otro módulo de test de alto nivel), usando descripciones del test y del grupo de test de forma correcta. Si hasta ahora no has subido el código que has venido realizando a GitHub, es el momento de hacerlo, porque lo vamos a necesitar un poco más adelante.



### Ejercicio 5. Ejercicio: Haced los dos primeros pasos antes de pasar al tercero. 
1. Darse de alta. Muchos están conectados con GitHub por lo que puedes usar directamente el usuario ahí. A través de un proceso de autorización, acceder al contenido e incluso informar del resultado de los tests.

![inicio](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio5/inicio.png)

2. Activar el repositorio en el que se vaya a aplicar la integración continua. Travis permite hacerlo directamente desde tu configuración; en otros se dan de alta desde la web de GitHub.

![sinc](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema2/capturas/ejercicio5/travis.png)

3. Crear un fichero de configuración para que se ejecute la integración y añadirlo al repositorio.



