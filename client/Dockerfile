FROM node:14-alpine as build
WORKDIR /frontend
COPY . .
RUN apk add --no-cache python make g++
RUN yarn install

ENV HOST 0.0.0.0
EXPOSE 8080