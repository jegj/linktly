# Key Generations

## Generation

### Private Key

```sh
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
```

### Public Key

```sh
openssl rsa -pubout -in private.pem -out public.pem
```

## Verify

```sh

openssl rsa -in private.pem -check
openssl rsa -pubin -in public.pem -text -noout
```

## Base64

```sh
base64 private.pem -w 0

```
