# Delongify

Delongify is a link shortening microservice. It consists of two REST API endpoints.

## Endpoints

### POST
#### `delongify.xyz/api/shorten`

Takes the following json document in the request body.

```json
{
    "Url": "http://www.example.com",
}
```

A unique 6-character slug is produced and used to create a key-value pair with the supplied url. This key-value pair is stored in a MongoDB database. This is a proof of concept, so the key-value pairs are only stored for 24 hours before they are expired.

### GET
#### `delongify.xyz/{slug}`

This is the redirection endpoint. A url that corresponds to the slug given in the path parameter is searched for in the database. If one is found, the user is redirected to that url. Otherwise, the user is redirected to an error page.