FROM golang:1.12 as build
LABEL Name=betty Version=0.0.1

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
EXPOSE 7000
CMD ["/app", "--port", "7000", "--host", "0.0.0.0"]
