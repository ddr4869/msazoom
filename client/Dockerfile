FROM node:22-alpine

RUN apk add --no-cache python3 make g++

WORKDIR /app
COPY package*.json ./

RUN npm install
COPY . .
COPY .env.local .env.local

RUN npm run build

EXPOSE 3000
CMD ["npm", "start"]