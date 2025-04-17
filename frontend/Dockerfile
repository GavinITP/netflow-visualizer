FROM node:18-alpine AS builder

WORKDIR /app

COPY frontend/package.json frontend/pnpm-lock.yaml ./

RUN npm install -g pnpm

RUN pnpm install --frozen-lockfile

COPY frontend .

RUN pnpm build

FROM node:18-alpine

WORKDIR /app

COPY --from=builder /app/build .
COPY --from=builder /app/package.json .
COPY --from=builder /app/pnpm-lock.yaml .
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 5173

CMD ["node", "index.js"]