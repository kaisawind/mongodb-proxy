FROM alpine:3.8

LABEL MAINTAINER="kaisawind <wind.kaisa@gmail.com>"

ADD bin/cmd-server /

RUN chmod 755 /cmd-server

ENTRYPOINT ["/cmd-server"]