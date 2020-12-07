# This is a duplicate of the Dockerfile in site-frontend but uses these files for backend-builder and git clones the
# frontend code for frontend-builder.

FROM node:12.4-alpine AS frontend-builder
RUN apk --no-cache add git
ENV PATH /app/node_modules/.bin:$PATH
WORKDIR /app
RUN git clone https://github.com/SpaceXLaunchBot/site-frontend.git /app
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
