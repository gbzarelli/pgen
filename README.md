![Issues](https://img.shields.io/github/issues/gbzarelli/pgen.svg)
![Forks](https://img.shields.io/github/forks/gbzarelli/pgen.svg)
![Stars](https://img.shields.io/github/stars/gbzarelli/pgen.svg)
![Release Version](https://img.shields.io/github/release/gbzarelli/pgen.svg)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/326bae7f15244e488de6d8aaa94ae17b)](https://www.codacy.com/gh/gbzarelli/pgen/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gbzarelli/pgen&amp;utm_campaign=Badge_Grade)

# PGen (building... learning GO Lang...)

Unrepeatable Protocol generator API in GO.

The PGen is a microservice created to generate service protocols for any type of services.
These protocols are readable so that people can easily record

The protocol number consists of 16 digits (by default), the first 8 (fixed) being the current date
and the last (configurable) random:

	Format: 'YYYYMMDD????????' sample: 2021120912345678

# Technologies

- [Gin - Go Web Framework](https://github.com/gin-gonic/gin)
- [Redis - KV NoSQL for Cache](https://github.com/go-redis)
- [Testify - Asserts and Mocks](https://github.com/stretchr/testify)
- [Metrics - GinProm by Depado](https://github.com/Depado/ginprom)

# Instructions to Run

### Prepare

- Clone de project:
  - `git@github.com:gbzarelli/pgen.git`
- Go to `pgen` directory:
  - `$cd pgen/`
- Build project
  - Prepare dependencies
    - `$make dependencies`
  - Build project
    - `$make build`

### Run with Docker

- Generate Dockerfile
  - `$make build-docker`
- Run full stack:
  - `$make run-docker-compose`

### Run in project

- Run the dependencies (infra / redis):
  - `$make run-docker-stack`
- Run project
  - `$go run main.go`

### Envs

Custom the decimal places value to generate a new protocol (default 8):
- `PROTOCOL_DECIMAL_PLACES_AFTER_DATE`
  - Default in Project and Dockerfile: 8

Custom Redis address:
- `REDIS_ADDRESS`
  - Default in Project: localhost:6379
  - Default in Dockerfile: redis:6379

## API

The project starts in `localhost:5000` with a unique endpoint to generate a new protocol:

## Endpoint: v1/protocol

#### Request:

```shell
curl --request POST --url http://localhost:5000/v1/protocol
```

#### Response:

```json
201 {"protocol": "2021121204066844"}
```

## Metrics by prometheus

#### View metrics by app

```shell
curl --request GET --url http://localhost:5000/metrics
```

#### Prometheus

If you run by docker just open prometheus in: `http://localhost:9090/`
