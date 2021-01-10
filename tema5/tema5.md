### 1. Instalar etcd3, averiguar qué bibliotecas funcionan bien con el lenguaje que estemos escribiendo el proyecto (u otro lenguaje), y hacer un pequeño ejemplo de almacenamiento y recuperación de una clave; hacer el almacenamiento desde la línea de órdenes (con etcdctl) y la recuperación desde el mini-programa que hagáis.

Viendo la documentación de [openstack](https://docs.openstack.org/install-guide/environment-etcd-ubuntu.html) hago la instalación usando el comando:

`apt install etcd`

En esta [página](https://etcd.io/docs/v3.3.12/integrations/) he buscado librerías para go. La de V3 para Go tenía mal el link, pero buscando a traves de la de V2 he encontrado el [repo](https://github.com/etcd-io/etcd/tree/master/client/v3) corecto.

En primer lugar he hecho la instalación del cliente. He podido comprobar que la versión actual del momento en el que realizaba la práctica daba problemas. He intentado solucionar el problema buscando información en los issues y por foros. No he encontrado ninguna manera de solucionarlo por lo que opté por usar la versión anterior. Utilizando esta versión no he encontrado ningún problema. La siguiente imagen es instalando la última versión, la que da problemas.

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio1/go_client.png)

Siguiendo los pasos de este [post](https://cloud.ibm.com/docs/databases-for-etcd?locale=es) hago lo siguiente para crear la clave usando etcdctl.

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio1/clave_valor.png)


Habiendo creado ya la clave he programado el programa que la recupera. En mi caso, estoy ulizando GO porque es el que estoy usando en mi proyecto. Para crear el programa he utilizado la información del repo del cliente de Go y este [post](https://pkg.go.dev/go.etcd.io/etcd/clientv3).

~~~
package main

import (  
    "fmt"
    "context"
//    "log"
    "go.etcd.io/etcd/clientv3"
    "time"
//    "strconv"
)

func main() {
   	cli, err := clientv3.New(clientv3.Config{
	Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	resp, err := cli.Get(ctx, "clave")
	cancel()
	if err != nil {
		// handle error!
	}else{
	// use the response
	fmt.Print(resp)
	fmt.Print("\r")
	}
	
	defer cli.Close()
}
~~~

Con el programa anterior consigo que recuperar la clave. El resultado que obtengo en la terminal es el siguiente:

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio1/programa.png)

### 2. Realizar una aplicación básica que use algún microframework para devolver alguna estructura de datos del modelo que se viene usando en el curso u otra que se desee. La intención de este ejercicio es simplemente que se vea el funcionamiento básico de un microframework, especialmente si es alguno que, como express, tiene un generador de código. Se puede usar, por otro lado, el lenguaje y microframework que se desee.

Empezé haciendo la instalación de express en el directorio de trabajo:

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio2/install_express.png)

Comprobé que previamente tenía que crear el fichero `package.json`. Encontré información sobre el archivo `package.json` y cómo crearlo [aquí](https://medium.com/noders/t%C3%BA-yo-y-package-json-9553929fb2e3). En la siguiente imagen se puede ver cómo lo creo y en la siguiente cómo vuelvo a hacer la instalación de express con el fichero ya creado.

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio2/npm_init.png)

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio2/install_express1.png)

Teniendo el entorno confgurado, he utilizado el código del [repo](http://jj.github.io/CC/documentos/temas/Microservicios.html) de los ejercicios:

~~~
#!/usr/bin/env node

var express=require('express');
var app = express();
var port = process.env.PORT || 8080;

app.get('/', function (req, res) {
    res.send( { Portada: true } );
});

app.get('/proc', function (req, res) {
    res.send( { Portada: false} );
});

app.listen(port);
console.log('Server running at http://127.0.0.1:'+port+'/');
~~~

Ejecutando este programa se puede ver en `localhost:8080` lo siguiente:

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio2/ejecuci%C3%B3n.png)

### 3. Programar un microservicio en express (o el lenguaje y marco elegido) que incluya variables como en el caso anterior.

El micriservicio que he programado ha sido este:
~~~
#!/usr/bin/env node

var express=require('express');
var app = express();
var port = process.env.PORT || 8080;

const coche_marca = "Opel";

app.get('/marca', function (req, res) {
    res.send( coche_marca );
});

app.get('/', function (req, res) {
    res.send( { Marca: true } );
});

app.listen(port);
console.log('Server running at http://127.0.0.1:'+port+'/');
~~~

Mientras se ejecuta el programa se puede ir a localhost:8080/marca y se ve el valor de la constante `coche_marca`.

### 4. Crear pruebas para las diferentes rutas de la aplicación.

He usado la herramienta que aparecía en los ejercicios: `supertest`. He encontrado información de su uso en esta [página](https://www.npmjs.com/package/supertest). Para empezar he instalado la herramienta ejecutando el siguiente comando en una terminar abierta en el directorio de trabajo:

`npm install supertest --save-dev`

Siguiendo las indicaciones de la página anteriormente mencionada y las del repo de ejercicios, he programado los siguientes tests:

~~~
const request = require('supertest');

const app = require('../ejer4.js');

request(app)
  .get('/')
  .expect('Content-Type', /json/)
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
  });
  
request(app)
  .get('/proc')
  .expect('Content-Type', /json/)
  .expect(200)
  .end(function(err, res) {
    if (err) throw err;
  });
~~~

Este programa está almacenano en su correspondiente carpeta de test, comprueba el programa ejer4.js que es el que he usado en el ejercicio 2, el que he sacado del repo de ejercicios. Cuando ejecuto los test en la terminal no me da ninguna indicación porque los esta ejecutando sin encontar ningún error.

### 5. Experimentar con diferentes gestores de procesos y servidores web front-end para un microservicio que se haya hecho con antelación, por ejemplo en la sección anterior.

Para este ejercio he probado `pm2`, la información para su uso la he sacado de [aquí](https://pm2.keymetrics.io/docs/usage/quick-start/). En las siguientes imágenes se puede ver la prueba que hago con el programa que estoy usando, en este caso lo renombro para que sea ejer5.js.

![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio5/pm2_1.png)
![](https://github.com/pabloalfaro/CC-Ejercicios/blob/main/tema5/Capturas/Ejercicio5/pm2_2.png)

### 6. Usar rake, invoke o la herramienta equivalente en tu lenguaje de programación para programar diferentes tareas que se puedan lanzar fácilmente desde la línea de órdenes un microservicio.

La herramienta que estoy usando en mi proyecto es `Makefile`. El fichero make que he creado para ejecutar los microservicios es el siguiente:

~~~
all: 	install run

install:
	npm init;
	npm install express --save;
	sudo npm install pm2 -g;

run:
	pm2 start ejer6.js -i 4;

stop:
	pm2 stop ejer6;
~~~
