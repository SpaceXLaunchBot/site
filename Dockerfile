FROM node:12.4-alpine AS builder
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY frontend .
RUN npm install
RUN npm run build

#EXPOSE 8080

# TODO:
#   - Build server
#   - Move frontend to static dir that server can serve
#   - Run server
#   - How to integrate with seperate db container?
