#build
FROM golang:1.23-alpine AS builder

#diretorio
WORKDIR /app

#copia os arquivos das dependencias
COPY go.mod ./

RUN go mod download

COPY . .

#compila o binario
RUN CGO_ENABLED=0 GOOS=linux go build -o servidor .

#execução
FROM alpine:latest

WORKDIR /app

#copia so o binario compilado antes
COPY --from=builder /app/servidor .


#container executa na porta 8080
EXPOSE 8080

CMD ["./servidor"]