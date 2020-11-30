FROM node:12.4-alpine AS builder
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY frontend .
RUN npm install
RUN npm run build

# TODO:
#   - Build server
#   - Move frontend to where server can serve
#   - Run server
#   - How to integrate with seperate db container?
