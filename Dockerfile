# ── Stage 1: build ──────────────────────────────────────────────────────────
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Cache de dependencias (se reconstruye solo cuando cambian go.mod/go.sum)
COPY go.mod go.sum ./
RUN go mod download

# Código fuente
COPY . .

# Compilar binario estático óptimo para producción
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# ── Stage 2: runtime ─────────────────────────────────────────────────────────
FROM alpine:3.21

WORKDIR /app

# Solo el binario — sin código fuente, sin toolchain de Go
COPY --from=builder /app/server /app/server

EXPOSE 8080

CMD ["/app/server"]
