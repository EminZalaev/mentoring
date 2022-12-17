FROM golang:1.18
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 8080

CMD go run main.go