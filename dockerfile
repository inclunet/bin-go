FROM ubuntu:latest
LABEL maintainer="inclunet"
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /bingo
COPY ./bingo/build /bingo
COPY ./bingo-go /usr/bin/bin-go
CMD bin-go
