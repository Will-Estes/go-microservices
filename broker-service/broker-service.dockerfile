FROM alpine:latest

RUN mkdir /app

# brokerApp is exe made in Makefile
COPY brokerApp /app

CMD ["/app/brokerApp"]