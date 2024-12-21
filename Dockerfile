FROM alpine
WORKDIR ./app
COPY  ./app .
EXPOSE 8087
CMD ["./counter_front"]

