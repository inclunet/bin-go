FROM ubuntu:latest
LABEL maintainer="inclunet"
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /bingo
COPY ./bingo/build /bingo
COPY ./bin-go /usr/bin/bin-go
EXPOSE 80
CMD bin-go
