# URL Shortener Service

A simple URL shortener service built with Go, Gin, and GORM.

## Features
- Create a new short URL
- Retrieve the original URL from a short URL
- Update an existing short URL
- Delete an existing short URL
- Track the number of times a short URL is accessed

## Setup
1. Clone the repository.
2. Create a `.env` file with the following variables:

```env
SERVER_PORT=8080
ENV=development
PostgresDSN=host=localhost user=postgres password=postgres dbname=url_shortener port=5432 sslmode=disable
```
3. Install dependencies:

```bash
make install
```

4. Start the server.

```bash
make run
```

## Endpoints
- `POST /api/urls` - Create a short URL
- `GET /api/urls/:shortCode` - Retrieve the original URL and increment access count
- `PUT /api/urls/:shortCode` - Update a short URL
- `DELETE /api/urls/:shortCode` - Delete a short URL


https://roadmap.sh/projects/url-shortening-service
