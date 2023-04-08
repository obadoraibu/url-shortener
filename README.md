# URL-shortener REST-API
Укорачиватель ссылок

### Для запуска приложения:

```
make build && make run
```

## API Methods

**Default port `8080`, Endpoint `http://localhost:8080/`**

### `/add` - создание короткой ссылки

* Method: `POST`

#### Default request:

```json5
{
  "long_url": "example.com",
  "short_url": "short"
}
```

Response:

```json5
{
  "delete_key": "b36316d7-ffbf-4300-b6c3-9f73ed02ce6e", 
  "short_url": "short"
}
```

### `/delete` - создание короткой ссылки

* Method: `POST`

#### Default request:

```json5
{
  "delete_key": "b36316d7-ffbf-4300-b6c3-9f73ed02ce6e",
  "short_url": "short"
}
```

Response:

```json5
HTTP 200
```

### `/{short_url}` - создание короткой ссылки

* Method: `GET`

Response:

```json5
Redirect
```