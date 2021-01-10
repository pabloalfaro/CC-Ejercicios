1. Instalar etcd3, averiguar qué bibliotecas funcionan bien con el lenguaje que estemos escribiendo el proyecto (u otro lenguaje), y hacer un pequeño ejemplo de almacenamiento y recuperación de una clave; hacer el almacenamiento desde la línea de órdenes (con etcdctl) y la recuperación desde el mini-programa que hagáis.

Viendo la documentación de [openstack](https://docs.openstack.org/install-guide/environment-etcd-ubuntu.html) hago la instalación usando el comando:

sudo snap install etcd

Siguiendo los pasos de este [post](https://cloud.ibm.com/docs/databases-for-etcd?locale=es) hago lo siguiente para crear la clave usando etcdctl.


En esta [página](https://etcd.io/docs/v3.3.12/integrations/) he buscado librerías para go. La de V3 para Go tenía mal el link, pero buscando a traves de la de V2 he encontrado el (repo)[https://github.com/etcd-io/etcd/tree/master/client/v3].

https://pkg.go.dev/go.etcd.io/etcd/clientv3


Encontré problemas con algunas dependencias:
https://github.com/etcd-io/etcd/issues/11563


2. Realizar una aplicación básica que use algún microframework para devolver alguna estructura de datos del modelo que se viene usando en el curso u otra que se desee. La intención de este ejercicio es simplemente que se vea el funcionamiento básico de un microframework, especialmente si es alguno que, como express, tiene un generador de código. Se puede usar, por otro lado, el lenguaje y microframework que se desee.

Encontré información sobre el archivo package.json [aquí](https://medium.com/noders/t%C3%BA-yo-y-package-json-9553929fb2e3).

3. Programar un microservicio en express (o el lenguaje y marco elegido) que incluya variables como en el caso anterior.


4. Crear pruebas para las diferentes rutas de la aplicación.

https://www.npmjs.com/package/supertest 


5. Experimentar con diferentes gestores de procesos y servidores web front-end para un microservicio que se haya hecho con antelación, por ejemplo en la sección anterior.

https://pm2.keymetrics.io/docs/usage/quick-start/


6. Usar rake, invoke o la herramienta equivalente en tu lenguaje de programación para programar diferentes tareas que se puedan lanzar fácilmente desde la línea de órdenes un microservicio.



