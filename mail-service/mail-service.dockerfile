FROM alpine:latest

RUN mkdir /app

# mailServiceApp is exe made in Makefile
COPY mailServiceApp /app
COPY templates /templates

CMD ["/app/mailServiceApp"]