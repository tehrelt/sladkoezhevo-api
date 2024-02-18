# SladkoEzhevo API

## Tech-stack
- Database - PostgreSQL

## Config

API config location is `config/.yaml`

- `port` - Port which the server starts on
- `env` - Enviroment config (dev, prod)
- `db` - Database config
  - `host` - host for postgreSQL db
  - `port` - port for postgreSQL db
  - `user` - postgreSQL user
  - `pass` - postgreSQL user's password
  
**Example:**
```yaml
// config/.yaml
port: 8080
env: "dev"
db:
  host: "127.0.0.1"
  port: 5432
  user: "postgres"
  pass: "1337"
```

## Dependencies
1. [Fiber](https://github.com/gofiber/fiber) - web framework
2. [CleanEnv](https://github.com/ilyakaznacheev/cleanenv) - config parsing