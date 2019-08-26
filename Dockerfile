FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app/src
ENTRYPOINT ["./picacomic"]

EXPOSE 9002