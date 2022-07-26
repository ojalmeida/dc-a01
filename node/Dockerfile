FROM node:14-alpine

RUN mkdir -p /etc
RUN mkdir -p /usr/src/node-api

COPY *.js package.json node_modules /usr/src/node-api/
COPY config.yaml /etc/

WORKDIR /usr/src/node-api

RUN npm install

EXPOSE 8081


ENTRYPOINT node app.js