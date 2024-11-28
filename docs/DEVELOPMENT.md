# Local Development

## Docker compose

[compose](https://github.com/docker/compose)

## Linters and tools

- [golangci-lint](https://golangci-lint.run/welcome/install/)
- [air](https://github.com/air-verse/air)
- [pino-pretty](https://github.com/pinojs/pino-pretty)

## Initial steps

```sh
ln -s configs/.env.example config/.env
```

## Initial DB Setup

```sh
ALTER ROLE linktly_user WITH LOGIN PASSWORD 'dev123';
```
