FROM node:12.4-alpine AS frontend-builder
ENV PATH /app/node_modules/.bin:$PATH
WORKDIR /app
RUN COPY frontend/. .
RUN yarn install
RUN yarn build

FROM golang:latest AS backend-builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./slb-webserver ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=frontend-builder /app/static static
COPY --from=backend-builder /app/slb-webserver .
EXPOSE 8080/tcp
CMD ["./slb-webserver"]
