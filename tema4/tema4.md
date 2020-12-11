# Autoevauación semana 8

### Ejercicio 1. Haced los dos primeros pasos antes de pasar al tercero. 

Para trabajar con estos sistemas, generalmente hay que ejecutar estos tres pasos:

1. Darse de alta. Muchos están conectados con GitHub por lo que puedes autentificarte directamente desde ahí. A través de un proceso de autorización, puedes acceder al contenido e incluso informar del resultado de los tests a GitHub.

Me he dado de alta en Travis con mi cuenta de GitHub, de esta manera ya tengo ya mis repos listos para activar la integración continua.

![Alta](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema4/Capturas/inicio.png)

2. Activar el repositorio en el que se vaya a aplicar la integración continua. Travis permite hacerlo directamente desde tu configuración; en otros se dan de alta desde la web de GitHub.

He activado mi repositorio del proyecto. En la imagen de abajo se pueden ver mis repos.

![Repos](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema4/Capturas/travis.png)

3. Crear un fichero de configuración para que se ejecute la integración y añadirlo al repositorio.

Mi archivo de configuración indica el lenguaje que estoy utilizando y la versión. Lo que hace es ejecutar mis tests de integración continua con el fichero Makefile.

~~~
language: go

go:
- 1.15

script: 
- make test
~~~

### Ejercicio 2. Configurar integración continua para nuestra aplicación usando Travis o algún otro sitio.

En la imagen de abajo se puede ver cómo se ejecutan los tests al realizar un commit en el repo. En este caso he hecho un cambio en el archivo README.md.

![Test](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema4/Capturas/testTravis.png)
