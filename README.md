# PGen (building... learning GO Lang...)

Protocol generator API in GO.

The PGen is a microservice created to generate service protocols for any type of services.
These protocols are readable so that people can easily record

The protocol number consists of 16 digits (by default), the first 8 (fixed) being the current date 
and the last (configurable) random:

	Format: 'YYYYMMDD????????' sample: 2021120912345678

# Technologies

- [Gin - Go Web Framework](https://github.com/gin-gonic/gin)
- [Redis - KV NoSQL for Cache](https://github.com/go-redis)

# Instructions to Run

- Clone de project:
  - `git@github.com:gbzarelli/pgen.git`
- Go to `pgen` directory: 
  - `$cd pgen/`
- Start the infrastructure
  - `$docker-compose up -d`
- Run the project
  - `go run main.go`

If needed custom the decimal places value to generate a new protocol (default 8), just 
create na env `PROTOCOL_DECIMAL_PLACES_AFTER_DATE` with the value.

## API

The project starts in `localhost:5000` with a unique endpoint to generate a new protocol:

### Request:
``
curl --request POST --url http://localhost:5000/v1/protocol
``

### Response:

```json
201 {"protocol": "2021121204066844"}
```
