#FROM scratch
#
#ADD ./api-gateway/index /
#ADD ./api-gateway/routes.json /
#
#ADD ./assets/ca-certificates.crt /etc/ssl/certs/
#
#CMD ["./index"]
#
#EXPOSE 3000


FROM node:8.12.0-alpine

WORKDIR /app

COPY ./api-gateway/package.json ./

RUN yarn install

COPY ./api-gateway/index.js ./
COPY ./api-gateway/routes.json ./

CMD ["node", "index.js"]

EXPOSE 3000
