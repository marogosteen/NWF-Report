FROM golang:1.17

RUN mkdir /src
WORKDIR /src

COPY . ./
RUN go build

EXPOSE 3000:3000

CMD ["./nwf-report"]
