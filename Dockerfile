FROM node:12.4-alpine AS staticbuilder
RUN apk --no-cache add git 
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
RUN git clone https://github.com/SpaceXLaunchBot/site-frontend.git /app
RUN yarn install
RUN yarn build
# Our files are now in /app/static

FROM golang:latest AS serverbuilder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./slbsiteserver ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=staticbuilder /app/static static
COPY --from=serverbuilder /app/slbsiteserver .
EXPOSE 8080/tcp
CMD ["./slbsiteserver"]
