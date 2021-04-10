FROM golang:1.16-alpine AS base
RUN apk --no-cache add make
WORKDIR /opt/web-app
COPY . .
EXPOSE 8080

FROM base AS development
ENTRYPOINT [ "make" ]

FROM base AS compiler
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:latest AS production
COPY --from=compiler /opt/web-app/server .
EXPOSE 8080
ENTRYPOINT [ "./server" ]
