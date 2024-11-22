FROM golang:1.19 AS build-stage
WORKDIR ./app
COPY  ./app .
CMD ["./counter"]

FROM build-stage AS run-test-stage
RUN go test -v
