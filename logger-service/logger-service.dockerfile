FROM alpine:latest

RUN mkdir /app

# loggerServiceApp is exe made in Makefile
COPY loggerServiceApp /app

CMD ["/app/loggerServiceApp"]