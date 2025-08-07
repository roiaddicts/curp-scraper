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
GET /curp/YOUR_CURP_HERE HTTP/1.1
Host: localhost:8080
```

**Curl example (success):**

```sh
curl -i http://localhost:8080/curp/YOUR_CURP_HERE
```

**Successful Response:**

```
HTTP/1.1 200 OK
Content-Type: application/json

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

**Curl example (error):**

```sh
curl -i http://localhost:8080/curp/INVALIDCURP1234567
```

**Error Response:**

```
HTTP/1.1 404 Not Found
Content-Type: application/json

{
  "error": {
    "code": "CURP_NOT_FOUND",
    "message": "CURP found in cache, but no data available"
  }
}
```

The `error.code` field in the response can have one of the following values:

| Code                   | HTTP Status Code          | Description                                                                                          |
| ---------------------- | ------------------------- | ---------------------------------------------------------------------------------------------------- |
| CAPTCHA_FAILURE        | 500 Internal Server Error | Failed to solve or validate the CAPTCHA.                                                             |
| ENCODING_ERROR         | 500 Internal Server Error | Failed to encode request data.                                                                       |
| REQUEST_CREATION_ERROR | 500 Internal Server Error | Failed to create the HTTP request.                                                                   |
| RENAPO_HTTP_ERROR      | 500 Internal Server Error | Error response from RENAPO service.                                                                  |
| RENAPO_RATE_LIMITED    | 429 Too Many Requests     | Rate limited by RENAPO. Please try again later. Look for `Retry-After` header to know when to retry. |
| DECODE_ERROR           | 500 Internal Server Error | Failed to decode the response from RENAPO.                                                           |
| CURP_NOT_FOUND         | 404 Not Found             | No records found for the provided CURP.                                                              |
| INVALID_CURP           | 400 Bad Request           | The provided CURP is not 18 characters long and/or does not match required format rules.             |

### `GET /health`

Health check endpoint.

**Example:**

```
GET /health HTTP/1.1
Host: localhost:8080
```

**Response:**

```json
{
  "status": "healthy",
  "balance": "0.00"
}
```

## Note

This is API version 1 (v1) and is subject to change.

## License

MIT License

---

_Made with Go._
