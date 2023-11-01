__________________________________________________________________________________
## PPT 01 - Intro

__________________________________________________________________________________
## PPT 02 - ¿Que es CURL? - 01
Es una herramienta de línea de comandos y una biblioteca de software utilizada para transferir datos a través de varios protocolos de red. Curl es muy versátil y admite una amplia gama de protocolos de transferencia, lo que lo convierte en una herramienta valiosa para realizar solicitudes y transferir datos desde y hacia servidores web y otros recursos de red. Algunos de los protocolos compatibles incluyen HTTP, HTTPS, FTP, FTPS, SCP, SFTP, LDAP, y muchos más.

En este caso lo utilizaremos para realizar pruebas de humo sobre una API REST con el protocolo HTTP.

```slide
Es una herramienta de línea de comandos y una biblioteca de software utilizada para transferir datos a través de varios protocolos de red.

Nos permite realizar solicitudes y transferir datos desde y hacia servidores web y otros recursos de red.

En este caso lo utilizaremos para realizar Smoke Test sobre una API REST con el protocolo HTTP.
```

__________________________________________________________________________________
## PPT 03 - CURL - Info - 02
A continuación, los flags más importantes de la sintaxis de curl:
- X: Especifica el método HTTP que se utilizará para realizar la solicitud. Por defecto, curl utiliza el método GET.
```bash
curl -X GET https://api.example.com
```
- H: Permite especificar encabezados HTTP personalizados para la solicitud.
```bash
curl -X GET https://api.example.com -H "Content-Type: application/json"
```
- D: Permite especificar datos para enviar en una solicitud. Por ejemplo para una solicitud POST o PUT.
```bash
curl -X POST https://api.example.com -H "Content-Type: application/json" -d '{"username": "admin", "password": "1234"}'
```

Otros flags:
- i: Muestra los encabezados de respuesta HTTP en la salida.
```bash
curl -X GET https://api.example.com -i
```
- o: Permite especificar un archivo de salida para almacenar la respuesta HTTP.
```bash
curl -X GET https://api.example.com -o response.json
```

En caso de necesitar mas información, pueden consultar la documentación oficial de CURL en el siguiente enlace: https://curl.se/docs/manpage.html o con el comando
```bash
curl --help
```

__________________________________________________________________________________
## PPT 04 - Implementacion CURL - 01 (envío de petición)
A modo de ejemplo, contamos con una API REST que nos permite realizar operaciones CRUD sobre una entidad llamada "movies". Intentaremos crear una nueva película utilizando el método POST.

Para ello, utilizaremos el siguiente comando:
```bash
curl -X POST https://api.example.com/movies -H "Content-Type: application/json" -d '{"title": "The Matrix", "year": 1999, "director": "Lana Wachowski, Lilly Wachowski"}'
```

__________________________________________________________________________________
## PPT 05 - Implementacion CURL - 02 (validación de respuesta)
Finalmente validamos la respuesta, que sea lo que esperamos. En este caso, esperamos que la respuesta sea un JSON con el siguiente formato:
```json
{
    "message": "Movie created successfully",
    "data": {
        "id": 1,
        "title": "The Matrix",
        "year": 1999,
        "director": "Lana Wachowski, Lilly Wachowski"
    }
}
```

De esta forma aplicamos un caso simple de prueba de humo sobre una API REST utilizando CURL

__________________________________________________________________________________
## PPT 06 - Cierre