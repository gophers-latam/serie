# Bienvenido

Bienvenido al proyecto Serie

## Objetivo
El objetivo de este proyecto es solucionar la siguiente serie:

  ````
  2, 4, 7, 28, 33, 198, ...
  
  1    2    4    7    28   33   198 ...
    +1   x2   +3   x4   +5    x6
   
   Si x = 1, y = 2
   Si x = 2, y = 4
   Si x = 3, y = 7
   Si x = 4, y = 28
   Si x = 5, y = 33
  
  ````
Encontrarás también lo siguiente:
* Cómo crear parametrizar tu aplicación mediante el uso de flags y proporcionar ayuda semejante a las man pages de Linux.
* Cómo desacoplar loggers que envíen sus resultados a los streams estándares de Linux.
* Cómo escribir los resultados de los streams estándares en un archivo.
* Cómo asignar estos loggers a una estructura para que puedas aplicar inyección de dependencias.
* Cómo crear una tabla de testing para evitar la creación de pruebas innecesarias.

## Indice
1. Requisitos.
2. Cómo ejecutar.

### 1. Requisitos.
* Tener instalado [Golang](https://golang.org/dl/) 1.15.* ó superior.

### 2. Cómo ejecutar.
* Para descargar este módulo ejecuta:
````
   go get github.com/gophers-latam/serie
  
  ````  
* Para mostrar la ayuda ejecuta el script help.sh.
* Para ejecutar los tests ejecuta el script run_tests.sh.
* Para ejecutar el módulo ejecuta el script run.sh.
