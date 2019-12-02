FROM node:latest

RUN npm install webpack -g

WORKDIR /usr/src/app
COPY . /usr/src/app/

# RUN apk add --no-cache --update make gcc libc-dev libpng-dev automake autoconf libtool nasm
RUN npm install
RUN npm build

EXPOSE 3000

CMD ["npm", "start"]