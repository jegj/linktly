# Local Development

## Linters and tools

- [golangci-lint](https://golangci-lint.run/welcome/install/)
- [air](https://github.com/air-verse/air)

## Initial steps

```sh
ln -s configs/.env.example config/.env
```

## Initial DB Setup

```sh
ALTER ROLE linktly_user WITH LOGIN PASSWORD 'dev123';
```
