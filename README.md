# Task Management App

Aplicación full stack para gestión de tareas con:

- Backend en Go + Gin
- Frontend en Nuxt + Vue
- Base de datos MongoDB
- Despliegue con Docker Compose

## Requisitos

### Para ejecutar con Docker (recomendado)

- Docker
- Docker Compose

### Para ejecutar sin Docker (modo local)

- Go (versión compatible con `backend/go.mod`)
- Node.js 22+
- npm
- MongoDB (local o en contenedor)

## Estructura del proyecto

```text
task-management-app/
	backend/
	frontend/
	mongo-init/
	docker-compose.yml
```

## Opción 1: Ejecutar todo con Docker Compose

Desde la raíz del proyecto: (`task-management-app/`)

```bash
docker compose up --build
```

Esto levanta:

- `mongodb` en `localhost:27017`
- `backend` en `localhost:8080`
- `frontend` en `localhost:3000`

## Configuración de MongoDB

### En Docker Compose

Ya está configurado en `docker-compose.yml`:

- URI usada por backend: `mongodb://mongodb:27017`
- Base de datos: `taskdb`

Además, se monta `mongo-init/init-db.js` en `/docker-entrypoint-initdb.d`, por lo que inserta datos de ejemplo al iniciar un volumen nuevo.

Nota: el script de inicialización se ejecuta solo cuando la data del contenedor está vacía (primer arranque del volumen).

### En local (sin Docker)

Si levantas MongoDB fuera de Docker, asegúrate de tener una instancia en:

- `mongodb://localhost:27017`

Luego usa esas variables para el backend (ver sección siguiente).

## Opción 2: Ejecutar backend manualmente

### 1) Configurar variables de entorno

En la carpeta `backend/`, crea un archivo `.env` con:
(puedes copiar el archivo .env.example)

```env
GO_REST_ENV=dev

GIN_MODE=debug
MONGO_DB_NAME=task-db
MONGO_URI=mongodb://localhost:27017
```

### 2) Instalar dependencias y ejecutar

```bash
cd backend
go mod download
go run ./cmd
```

Backend disponible en:

- `http://localhost:8080`
- API base: `http://localhost:8080/api`

## Opción 3: Ejecutar frontend manualmente

### 1) Configurar variables de entorno

En la carpeta `frontend/`, crea un archivo `.env` con:
(puedes copiar el archivo .env.example)

```env
NUXT_PUBLIC_API_BASE_URL=http://localhost:8080/api
NUXT_PUBLIC_API_TIMEOUT=10000
```

### 2) Instalar dependencias y ejecutar

```bash
cd frontend
npm install
npm run dev
```

Frontend disponible en:

- `http://localhost:3000`

## Endpoints principales del backend

- `POST /api/tasks`
- `GET /api/tasks`
- `GET /api/tasks/filter?page=1&limit=10`
- `GET /api/tasks/:id`
- `PUT /api/tasks/:id/complete`
- `DELETE /api/tasks/:id`
- `POST /api/tasks/:id/tags`

## Problemas comunes

- Error `MONGO_URI no está definido`:
	- Falta configurar variables de entorno en backend.

- Frontend falla con `API base URL is not configured`:
	- Falta `NUXT_PUBLIC_API_BASE_URL` en `frontend/.env`.

- No se cargan datos de ejemplo en MongoDB:
	- El volumen ya existía. Ejecuta `docker compose down -v` y vuelve a levantar.