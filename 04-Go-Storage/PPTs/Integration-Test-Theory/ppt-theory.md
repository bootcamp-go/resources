_______________________________________________________________
## Slide 01 - Intro

_______________________________________________________________
## Slide 02 - Temario
- Test de integración
- Herramientas
- Docker
- Package go-txdb
- Implementación en Go

```slide
- Integration Test
- Herramientas
- Docker
- Package go-txdb
- Implementación en Go
```

_______________________________________________________________
## Slide 03 - ¿Que es un test de integración?
Un test de integración es una prueba enfocada en verificar la correcta interacción de nuestro sistema con componentes externos, como bases de datos o servicios externos, que representan procesos ajenos a nuestra aplicación. El proceso aplicado por estos componentes está abstraído y queda fuera de nuestro control.

Durante el test, evaluaremos si ante cierto input, el output que nuestro sistema genera después del proceso aplicado, incluyendo la interacción con la data obtenida de estos componentes externos, es el esperado.

Las pruebas de integración son cruciales para asegurar que nuestro sistema funcione sin problemas con procesos externos.

```slide
¿Que es un Integration Test?
Un Integration Test es una prueba enfocada en verificar la correcta interacción de nuestro sistema con componentes externos, como bases de datos o servicios externos, que representan procesos ajenos a nuestra aplicación

Las pruebas de integración son cruciales para asegurar que nuestro sistema funcione sin problemas con procesos externos.
```

_______________________________________________________________
## Slide 04 - ¿Cuales son sus ventajas?
- Detección temprana de problemas: Las pruebas de integración permiten identificar problemas de interacción con componentes externos antes de llegar a pruebas más exhaustivas, lo que facilita su corrección en una etapa temprana del desarrollo.
- Garantiza la interoperabilidad: Asegura que tu sistema funcione de manera efectiva con componentes externos como bases de datos y servicios externos, como APIs, para evitar problemas de compatibilidad.
- Ahorra tiempo y costos a largo plazo: Identificar y solucionar problemas de integración desde el principio ahorra tiempo y recursos que de otra manera se gastarían en correcciones costosas después del lanzamiento.
- Aumenta la confiabilidad: Contribuye a la confiabilidad y estabilidad del sistema al validar su comportamiento en situaciones del mundo real.
- Facilita la escalabilidad: Al garantizar una correcta integración con componentes externos, se facilita la escalabilidad del sistema a medida que se agregan más funcionalidades o se actualizan los servicios externos.

```slide
¿Cuales son sus ventajas?
- Detección temprana de problemas: Permiten identificar problemas de interacción con componentes externos antes de llegar a pruebas más exhaustivas.
- Garantiza la interoperabilidad: Se evitan problemas de incompatibilidad con componentes externos
- Ahorra tiempo y costos a largo plazo
- Aumenta la confiabilidad
- Facilita la escalabilidad
```

_______________________________________________________________
## Slide 05 - Herramientas
Una forma en la que podemos realizar pruebas de integración son los containers. Estos permiten empaquetar una aplicación con todas sus dependencias en un entorno aislado, lo que facilita su ejecución en cualquier sistema operativo.

Algunos software que nos permiten trabajar con containers son:
- Docker: Docker es una plataforma de contenedores que permite empaquetar aplicaciones y sus dependencias en contenedores. Estos contenedores son entornos aislados que pueden ejecutarse de manera consistente en diferentes sistemas operativos. Docker proporciona herramientas para crear, distribuir y administrar contenedores de aplicaciones.

- Containerd: Es un tiempo de ejecución de contenedores de alto rendimiento y código abierto que se utiliza en la infraestructura de contenedores de Docker. Proporciona funcionalidades esenciales para la ejecución de contenedores, como la creación, ejecución y administración de contenedores.

- Podman: Es una herramienta de administración de contenedores de código abierto que proporciona una experiencia similar a Docker. Permite a los usuarios crear, ejecutar y administrar contenedores sin requerir un demonio central. Podman es una alternativa a Docker y se utiliza para trabajar con contenedores de manera segura y eficiente.

- LXC (Linux Containers): LXC, o Linux Containers, es una tecnología de virtualización a nivel de sistema operativo que permite crear y administrar entornos de contenedores ligeros en sistemas Linux. LXC proporciona un enfoque más directo para trabajar con contenedores a nivel del sistema operativo, lo que puede ser útil en escenarios donde se requiere un mayor control sobre los contenedores.

Estos containers luego se pueden orquestar con herramientas como:
- Docker Compose
- Kubernetes
- OpenShift

```slide
Herramientas

Podemos hacer uso de los containers. Estos permiten empaquetar una aplicación con todas sus dependencias en un entorno aislado, lo que facilita su ejecución en cualquier sistema operativo.

Software de containers:
- Docker: Es una plataforma de containers que permite empaquetar aplicaciones y sus dependencias en containers
- Podman: Es una herramienta de administración de contenedores de código abierto que proporciona una experiencia similar a Docker
- LXC (Linux Containers): Es una tecnología de virtualización a nivel de sistema operativo que permite crear y administrar entornos de contenedores ligeros en 
sistemas Linux

Software de orquestación:
- Docker Compose
- Kubernetes
- OpenShift
```

_______________________________________________________________
## Slide 06 - Docker - Crear una imagen
En docker tenemos 2 conceptos.
Por un lado tenemos las imágenes, que son plantillas de solo lectura que contienen las instrucciones para crear un contenedor. Estas imágenes se pueden descargar de un repositorio público o privado, o se pueden crear a partir de un archivo llamado Dockerfile.

Por otro lado tenemos los contenedores, que son instancias de una imagen que se pueden ejecutar en un sistema operativo. Estos contenedores son entornos aislados que contienen todo lo necesario para ejecutar una aplicación, incluidas las dependencias y las variables de entorno.

En este ejemplo vamos a crear una imagen de MySQL con un Dockerfile, el cual importa una imagen de mysql desde el hub oficial de docker (similar a los repositorios de github pero de imagenes). Luego copia un script sql que se va a ejecutar al levantar el contenedor. Por ultimo setea algunas variables de entorno. El script lo tendremos que armar localmente y cuando creemos la imagen se va a copiar a la imagen.

```slide
Docker - Levantar un container

En docker tenemos 2 conceptos:
- images: Son plantillas de solo lectura que contienen las instrucciones para crear un contenedor. Estas imágenes se pueden descargar de un repositorio público o privado, o se pueden crear a partir de un archivo llamado Dockerfile.
- containers: Son instancias de una imagen que se pueden ejecutar en un sistema operativo. Estos contenedores son entornos aislados que contienen todo lo necesario para ejecutar una aplicación, incluidas las dependencias y las variables de entorno.

Ejemplo
- Crear una imagen de MySQL con un Dockerfile
> dockerfile | > script.sql
- Construir la imagen
> docker build -t test-db .
```

```dockerfile
# Dockerfile
FROM mysql:8.0.26
# Copy sql script
COPY ./init.sql /docker-entrypoint-initdb.d/
# Environment variables
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=test
```

```sql
-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS test_db;
-- Usar la base de datos
USE test_db;
-- Crear una tabla de ejemplo
CREATE TABLE IF NOT EXISTS example_table (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);
```

- Construir la imagen
```bash
# Construir la imagen
docker build -t test-db .
```

_______________________________________________________________
## Slide 07 - Docker - Levantar Container

Para levantar el container, ejecutamos el siguiente comando:

```bash
# Levantar el contenedor
docker run -d --name test-db -p 3306:3306 test-db
```
-d: Corre el contenedor en segundo plano, caso contrario la terminal queda bloqueada
--name: Le asigna un nombre al contenedor
-p: Mapea el puerto 3306 del contenedor al puerto 3306 de la maquina host. El primero mapea el puerto interno del contenedor, exponiendo el servicio.
test-db: Es el nombre de la imagen que se va a correr

```slide
Para levantar el contenedor, ejecutamos el siguiente comando:

> bash
-d: Corre el container en detached mode
--name: Asigna un nombre al container
-p: Mapea el puerto del container al puerto del host.
test-db: Es el nombre de la imagen que se va a correr
```

_______________________________________________________________
## Slide 08 - Docker - Otros comandos
Docker tambien nos permite listar los contenedores que tenemos corriendo e inspeccionarlos.

```bash
# Ver los contenedores corriendo
docker ps
```

Para detener un contenedor podemos usar el comando stop

```bash
# Detener el contenedor
docker stop test-db
```

_______________________________________________________________
## Slide 09 - Implementacion en Go - Intro
Nuestra API de ejemplo usando Chi, es sobre Movies. Se compone de las siguientes capas:
- Repository: se encarga de interactuar con la base de datos.
- Handler: se encarga de manejar las peticiones y responderlas.
- Application: se encarga de orquestar las funcionalidades de la app.

En nuestro caso utilizaremos una base de datos compartida.

_______________________________________________________________
## Slide 10 - Implementacion en Go - Repository
```go
package internal

var ErrMovieAlreadyExists = errors.New("movie already exists")

type Movie struct {
    Id          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type MovieRepository interface {
    // Save a movie
    Save(m *Movie) (err error)
}
```

```go
package repository

type MovieMysql struct {
    db *sql.DB
}

func (m *MovieMysql) Save(movie *internal.Movie) (err error) {
    // prepare statement
    stmt, err := m.db.Prepare("INSERT INTO `movies` (`id`, `title`, `description`) VALUES(?, ?, ?)")
    if err != nil {
        return
    }
    defer stmt.Close()
    // execute statement
    _, err = stmt.Exec(movie.Id, movie.Title, movie.Description)
    if err != nil {
        var mysqlErr *mysql.MySQLError
        if errors.As(err, &mysqlErr) {
            switch mysqlErr.Number {
            case 1062:
                err = internal.ErrMovieAlreadyExists
            }
            return
        }
        return
    }
    return
}
```

_______________________________________________________________
## Slide 11 - Implementacion en Go - Handler
```go
package handler

type Movie struct {
    rp internal.MovieRepository
}

func (m *Movie) Create() http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // request
        var movie internal.Movie
        err := request.JSON(r, &movie)
        if err != nil {
            response.Error(w, http.StatusBadRequest, "invalid request body")
            return
        }
        // save movie
        err = m.rp.Save(&movie)
        if err != nil {
            switch {
            case errors.Is(err, internal.ErrMovieAlreadyExists):
                response.Error(w, http.StatusConflict, "movie already exists")
            default:
                response.Error(w, http.StatusInternalServerError, "internal server error")
            }
            return
        }
        // response
        response.JSON(w, http.StatusCreated, movie)
    }
}
```

_______________________________________________________________
## Slide 12 - Implementacion en Go - Test
Vamos a realizar un test de integracion de nuestra api con la base de datos. Para esto vamos a utilizar un contenedor de mysql.
De esta forma nuestros tests se conectaran con la db a través del puerto que hayamos mappeado.

```go
package handler_test

func TestMovieCreate(t *testing.T) {
    // arrange
    // - db: init
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test_db")
    require.NoError(t, err)
    defer db.Close()
    rp := repository.MovieMysql{db: db}
    hd := handler.Movie{rp: &rp}
    hdFunc := hd.Create()
    // act
    request := httptest.NewRequest(http.MethodPost, "/movies", strings.NewReader(
        `{"id": 1, "title": "test", "description": "test"}`
    ))
    response := httptest.NewRecorder()
    hdFunc(response, request)
    // assert
    assert.Equal(t, http.StatusCreated, response.Code)
    assert.JSONEq(t, `{"id":1,"title":"test","description":"test"}`, response.Body.String())
}
```

_______________________________________________________________
## Slide 13 - Package go-txdb

Recordando los principios FIRST, enfocados en Isolated y Repeatable. Los test deben ser independientes entre si, estar aislados y las condiciones iniciales con las que se ejecutan deben ser las mismas, lo que permite que sean repetibles.

Si en nuestros test utilizamos una db compartida, debemos tener precaución con el concepto de Shared Memory.
Pueden ocurrir 2 cosas:
- Condiciones Iniciales: Si un test modifica la db, puede afectar a otro test, donde ya no trabaja sobre un entorno con las mismas condiciones iniciales. Es necesario que posterior a cada test se limpie la db.
- Concurrencia: En caso de los test ejecutarse concurrentemente (se puede con t.Parallel()), pueden ocurrir dataraces, donde si un test modifica la db, mientras otro esta ejecutandose, puede haber inconsistencias. Por eso es necesario que los test se ejecuten de forma secuencial a través de algun mecanismo de bloqueo.

_______________________________________________________________
## Slide 14 - Package go-txdb

Para solucionar estos potenciales conflictos, en Go tenemos el package go-txdb que abstrae el driver mysql y nos permite abrir una conexión con la db, iniciando una transacción aislada. Al cerrarse aplica rollback.

Aplica:
- Rollback: Cuando se abre una conexion, al cerrarse aplica un rollback, por lo que no se persisten los cambios.
- Isoation: Cada transaccion se ejecuta de forma aislada, por lo que no se ven afectadas por otras transacciones.

Al abrir una conexión, se inicia una transacción aislada, lo que evita verse afectada por otras transacciones. Al cerrarse se aplica un rollback, por lo que no se persisten los cambios.

¡Aclaración! Algunos puntos a tener en cuenta.

DBMS como MySQL y PostgreSQL no permiten rollback:
- La estructura de la base de datos no se puede revertir. Como crear una tabla.
- Autoincrement: el valor de un campo autoincremental no se puede revertir.

_______________________________________________________________
## Slide 15 - Package go-txdb - implementacion
En nuestro test anterior, primero debemos registrar el driver txdb

En go la funcion init es una función especial que se utiliza para inicializar paquetes dentro de un programa. La función init se ejecuta automáticamente antes de que se inicie la función main en un programa Go. Cada paquete en Go puede tener una o varias funciones init, y todas se ejecutarán en el orden en que se importan los paquetes. En nuestro caso se ejecutara previo a los tests para registrar el driver de txdb sobre el de mysql. Una vez hecho esto, las conexiones que abramos con sql.Open() inician una transaccion aislada y con rollback, por lo que no se guardaran los cambios en la base de datos.

```go
import (
    "database/sql"
    "github.com/DATA-DOG/go-txdb"
    "github.com/go-sql-driver/mysql"
)

func init() {
    cfg := mysql.Config{
        User:                 "root",
        Net:                  "tcp",
        Addr:                 "127.0.0.1:3306",
        DBName:               "test_db",
    }
    txdb.Register("txdb", "mysql", cfg.FormatDSN())
}
```

_______________________________________________________________
## Slide 16 - Package go-txdb - implementacion

Finalmente debemos abrir la conexión con el driver txdb. En identifier no es necesario poner el DSN, ya que el driver txdb ya lo tiene registrado.
```go
package handler_test

func TestMovieCreate(t *testing.T) {
    // arrange
    // - db: init
    db, err := sql.Open("txdb", "")
    require.NoError(t, err)
    defer db.Close()
    // ...
}
```

_______________________________________________________________
## Slide 17 - Conclusion

Hemos aprendido sobre las pruebas de integración y sus ventajas en el desarrollo de software. Las pruebas de integración se centran en verificar la interacción correcta de un sistema con componentes externos, como bases de datos o servicios externos.

Algunas de las ventajas clave de las pruebas de integración incluyen la detección temprana de problemas, garantizar la interoperabilidad, ahorrar tiempo y costos a largo plazo, aumentar la confiabilidad y facilitar la escalabilidad del sistema.

Vimos herramientas como Docker, Containerd, Podman y LXC para realizar pruebas de integración utilizando containers

Luego aprendimos cómo crear una imagen de MySQL con Docker y cómo realizar pruebas de integración en una API utilizando un contenedor de MySQL. Ademas hablamos de los principios FIRST y las posibles desventajas de utilizar una base de datos compartida en pruebas de integración, así como la solución mediante el paquete go-txdb.

En resumen, las pruebas de integración son esenciales para garantizar que un sistema funcione de manera efectiva con componentes externos, y existen herramientas y enfoques para llevar a cabo estas pruebas de manera efectiva, manteniendo la independencia y el aislamiento de las pruebas.

```slide
Hemos aprendido sobre las pruebas de integración y sus ventajas en el desarrollo de software, entre ellas la detección temprana de problemas y garantizar la interoperabilidad.

Vimos herramientas de containers, como Docker para realizar pruebas de integración.

Luego aprendimos cómo levantar un container de MySQL con Docker.

Finalmente implementamos una prueba de integracion con el package go-txdb manteniendo los principios Isolated y Repeatable.
```