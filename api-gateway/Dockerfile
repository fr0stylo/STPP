FROM node:8.12.0-alpine

WORKDIR /app

COPY package*.json ./
COPY yarn.lock ./

RUN yarn install

COPY . .

CMD ["yarn", "start"]

EXPOSE 3000
