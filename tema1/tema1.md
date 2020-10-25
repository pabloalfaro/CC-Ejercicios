# Autoevaluación semana 1

### Ejercicio 1. Buscar una aplicación de ejemplo, preferiblemente propia, y deducir qué patrón es el que usa. ¿Qué habría que hacer para evolucionar a un patrón tipo microservicios?

La aplicación que voy a usar es la que realicé en mi TFG. Es una aplicación que, usando datos de pacientes del Hospital de Navarra, da una valoración del rendimiento del servicio de Traumatología. En mi caso la aplicación tendría una estructura de cliente servidor. Hay un fontend dónde puedes elegir una tecnica de las implementadas para comprobar el rendimiento y un backend dónde se llevan a cabo las operaciones. Para transformarlo en un patrón de microservicios habría que tener un servicio encargado de mostrar la interface y luego un servicio por cada una de las técnicas con las que se pueda llevar realizar la estimación del rendimiento.

### Ejercicio 2. En la aplicación que se ha usado como ejemplo en el ejercicio anterior, ¿podría usar diferentes lenguajes? ¿Qué almacenes de datos serían los más convenientes?

En mi caso toda la aplicación la implementé en Python. En el caso de la interface se podria hacer en cualquier lenguaje con el que se puedan implementar interfaces de manera sencilla o que las interfaces resultantes tengan un buen acabado, en mi caso se me ocurre java. Para la parte de las tecnicas de predicción se podría usar cualquier lenguaje de machine learning, por ejemplo R. 

Para el almacenamiento utilizaría MongoDB, en mi caso la aplicación tomaba los datos de un fichero pero utilizando esta base de datos se podría trabajar de una manera más sencilla.
