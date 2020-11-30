FROM node:12.4-alpine AS builder
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY . .
RUN npm install
RUN npm run build

FROM pierrezemb/gostatic
COPY --from=builder /app/dist/. /srv/http/.
EXPOSE 8043

# Basically the same as GoStatic
#docker run -d --name spacexlaunchbotsite --restart unless-stopped \
#    -p 80:8043 \
#    psidex/spacexlaunchbotsite
