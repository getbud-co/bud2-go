# Bud 2.0 Web Application

Frontend em Next.js 15 para o Bud 2.0.

## Requisitos

- Node 24.x
- Yarn 1.22.x

## Ambiente

Copie `frontend/.env.example` para `frontend/.env` e ajuste:

```bash
BUD_API_URL=http://localhost:8080
NEXT_PUBLIC_CHECKIN_SYNC_URL=
```

`BUD_API_URL` é usado pelas rotas server-side do frontend para autenticar e consumir a API Go.

## Desenvolvimento

```bash
yarn install
yarn lint
yarn build
```

## Autenticação

- O frontend usa o fluxo de JWT do backend.
- Login chama `POST /auth/login`.
- Sessão corrente usa `GET /auth/session`.
- Troca de organização ativa usa `PUT /auth/session`.
- O token fica em cookie `HttpOnly` gerenciado pelo frontend.

## Docker

```bash
docker build -t bud-web .
docker run -p 3000:3000 --env-file .env bud-web
```
