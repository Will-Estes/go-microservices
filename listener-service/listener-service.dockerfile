FROM alpine:latest

RUN mkdir /app

# loggerServiceApp is exe made in Makefile
COPY listenerServiceApp /app

CMD ["/app/listenerServiceApp"]