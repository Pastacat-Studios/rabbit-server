FROM golang:1.25.5

WORKDIR /build

COPY . ./

RUN CGO_ENABLED=1 go build -v -o /usr/local/bin/rabbitserver ./

WORKDIR /server

RUN rm -rf /build

COPY frontend/templates ./frontend/templates

CMD ["rabbitserver"]