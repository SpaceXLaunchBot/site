FROM node:12.4-alpine AS builder
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY . .
RUN npm install
RUN npm build

FROM pierrezemb/gostatic
COPY --from=builder /app/dist/. /srv/http/.

# From GoStatic docs
EXPOSE 8043
