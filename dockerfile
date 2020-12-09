from alpine:latest

RUN apk add --no-cache bash && rm -f /var/cache/apk/*

COPY ./bin /usr/local/bin
RUN chmod 777 /usr/local/bin
CMD ["/usr/local/bin/server"]
EXPOSE 8080