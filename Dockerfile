FROM golang:1.25.3-alpine

WORKDIR /build

COPY . ./

RUN apk add gcc musl-dev

RUN CGO_ENABLED=1 go build -v -o /usr/local/bin/rabbitserver ./

WORKDIR /server

RUN rm -rf /build

COPY frontend/templates ./frontend/templates
COPY assets ./assets

CMD ["rabbitserver"]