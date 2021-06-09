FROM golang:latest AS go-builder
WORKDIR /build
COPY cmd cmd
COPY internal internal
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./slb-webserver ./cmd/server/main.go

FROM node:14.17 AS frontend-builder
ENV PATH /app/node_modules/.bin:$PATH
WORKDIR /app
COPY frontend/. .
RUN yarn install
RUN yarn prodbuild

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /build/slb-webserver .
COPY --from=frontend-builder /app/build ./frontend_build
EXPOSE 8080/tcp
ENTRYPOINT ["./slb-webserver"]
