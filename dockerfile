FROM ubuntu:latest
LABEL maintainer="inclunet"
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /bingo
COPY ./bingo/build /accessbot
COPY ./bingo /usr/bin/bingo
CMD bingo
