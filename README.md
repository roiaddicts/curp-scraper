# CURP Scraper

CURP Scraper is a lightweight Go API server for extracting and validating CURP (Clave Única de Registro de Población) information from official sources. It provides a simple HTTP interface to fetch CURP data, with built-in health checks and CAPTCHA solving support via [Capsolver](https://capsolver.com/).

## Features

- Fetch and validate CURP information via HTTP API
- Health check endpoint for monitoring
- CAPTCHA solving integration (Capsolver)
- Easy to deploy and configure

## Requirements

- Go 1.22+ (recommended Go 1.24.2)
- [Capsolver API key](https://dashboard.capsolver.com/)

## Setup

1. **Clone the repository:**

   ```sh
   git clone <your-repo-url>
   cd bank-scraper
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Configure environment variables:**

   Create a `.env` file in the project root:

   ```
   CAPSOLVER_API_KEY=YOUR_CAPSOLVER_API_KEY
   ```

4. **Run the server:**
   ```sh
   go run cmd/server/main.go
   ```

## API Endpoints

### `GET /curp/{curp}`

Fetch CURP information.

**Path Parameters:**

- `curp` (required): The CURP string to look up.

**Example:**

```
GET http://localhost:8080/curp/YOUR_CURP_HERE
```

**Successful Response:**

```json
{
  "data": {
    "curp": "RODR850715MDFLRS09",
    "nombres": "JUANA MARIA",
    "primerApellido": "RODRIGUEZ",
    "segundoApellido": "LOPEZ",
    "claveGenero": "MUJER",
    "genero": "MUJER",
    "fechaNacimiento": "15/07/1985",
    "diaNacimiento": "15",
    "mesNacimiento": "07",
    "anioNacimiento": "1985",
    "claveEntidadNacimiento": "DF",
    "entidadNacimiento": "DISTRITO FEDERAL"
  }
}
```

**Error Response:**

```json
{
  "error": {
    "code": "CURP_NOT_FOUND",
    "message": "CURP found in cache, but no data available"
  }
}
```

### `GET /health`

Health check endpoint.

**Example:**

```
GET http://localhost:8080/health
```

**Response:**

```json
{
  "status": "healthy",
  "balance": "0.00"
}
```

## License

MIT License

---

_Made with Go._
