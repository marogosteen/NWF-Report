FROM golang:1.17

RUN mkdir /src
WORKDIR /src

COPY . ./
RUN go build

EXPOSE 8080

CMD ["./nwf-report"]
