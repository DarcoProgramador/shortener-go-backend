# URL Shortener Backend

Este proyecto es el backend de un acortador de URLs. Proporciona una API para acortar URLs largas.

## Características

- Acortar URLs largas.
- Obtener URLs originales.
- Estadísticas de cantidad de visitas.
- Eliminar URLS acortadas.
- Actualizar link acortado por una nueva URL.

## Requisitos

- Go

## Instalación

1. Clona el repositorio:
    ```sh
    git clone https://github.com/tu-usuario/url-shortener-backend.git
    ```
2. Navega al directorio del proyecto:
    ```sh
    cd url-shortener-backend
    ```
3. Crea el archivo .ENV:
    ```sh
    GOOSE_DRIVER=sqlite3
    GOOSE_DBSTRING=./urls.db
    GOOSE_MIGRATION_DIR=./internal/database/migrations
    ```

## Uso

1. Instala los paquetes:
    ```sh
    go install
    ```
2. Contruye el ejecutable para tu maquina:
   ```sh
    go build cmd/main.go
    ```
3. Ejecuta el servidor:
    ```sh
    ./main.exe 
    ```
3. La API estará disponible en `http://localhost:8080`.

## Endpoints

- `POST /shorten`: Acorta una URL larga.
    ```sh
    curl --location 'http://localhost:8080/shorten' \
    --header 'Content-Type: application/json' \
    --data '{
        "url": "https://www.google.com"
    }'
    ```
- `GET /shorten/{short_code}`: Obtiene la URL original.
    ```sh
    curl --location 'http://localhost:8080/shorten/Zl1CY0'
    ```
- `GET /shorten/{short_code}/stats`: Obtiene estadísticas de uso.
    ```sh
    curl --location 'http://localhost:8080/shorten/Zl1CY0/stats'
    ```
- `PUT /shorten/{short_code}`: Actualiza la url del link acortado
    ```sh
    curl --location --request PUT 'http://localhost:8080/shorten/Zl1CY0' \
    --header 'Content-Type: application/json' \
    --data '{
        "url": "https://roadmap.sh/projects/url-shortening-service"
    }'
    ```
- `DELETE /shorten/{short_code}`: Elimina la URL acortada de la base de datos.
    ```sh
    curl --location 'http://localhost:8080/shorten/Zl1CY0'
    ```  

## Licencia
Este proyecto está bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.

****
[Reto de roadmap.sh](https://roadmap.sh/projects/url-shortening-service)