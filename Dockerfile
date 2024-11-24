FROM golang:1.22-alpine3.19 as builder

WORKDIR /app


COPY . .

RUN go build -o acervo .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/acervo /app/acervo

CMD ["/app/acervo"]


