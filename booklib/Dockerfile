FROM golang:latest

ENV PORT 8080

WORKDIR /app
ADD . /app/
RUN go mod tidy 
RUN go mod verify

EXPOSE $PORT

CMD ["go","run","main.go"]