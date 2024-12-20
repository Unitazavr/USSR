FROM alpine
WORKDIR ./app
COPY  ./app .
CMD ["./counter_front"]

