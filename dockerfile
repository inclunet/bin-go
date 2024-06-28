FROM golang:1.21-alpine as build-backend
WORKDIR /app
COPY ./ .
RUN go build -o inclugames main.go

FROM node:18 as build-frontend
WORKDIR /app
COPY ./bingo .
RUN npm install
RUN npm run build

FROM ubuntu:latest
LABEL maintainer="inclunet"
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /app
COPY --from=build-backend /app/inclugames /usr/bin/inclugames
COPY --from=build-backend /app/classes.json /app
COPY --from=build-frontend /app/build /app
EXPOSE 80
CMD inclugames
