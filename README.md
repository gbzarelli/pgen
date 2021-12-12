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

[... TODO]
