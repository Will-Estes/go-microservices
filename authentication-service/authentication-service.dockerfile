FROM alpine:latest

RUN mkdir /app

# authApp is exe made in Makefile
COPY authApp /app

CMD ["/app/authApp"]