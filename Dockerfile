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
COPY oldfrontend/. .
# Write short hash of current commit to version file
COPY .git/refs/heads/master .
RUN head -c 7 ./master > ./public/version && rm ./master
RUN yarn install
RUN yarn prodbuild

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /build/slb-webserver .
COPY --from=frontend-builder /app/build ./frontend_build
EXPOSE 8080/tcp
ENTRYPOINT ["./slb-webserver"]
