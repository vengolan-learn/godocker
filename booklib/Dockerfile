FROM golang:latest

ENV PORT 8080

WORKDIR /app

ADD . /app/

EXPOSE $PORT


#CGO_ENABLED=0 that the generated build is statically linked(meaning no dynamic linking libraries uses at runtime)
RUN CGO_ENABLED=0 go build -o booklib 


#fail build if the ./booklib executbale uses external libraries ( dynamically linked)
RUN ldd booklib 2>&1 | grep -q "not a dynamic executable"

CMD ["./booklib"]