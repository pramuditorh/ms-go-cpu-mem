FROM registry.access.redhat.com/ubi8/go-toolset:1.19.10-3.1690904352 AS builder

USER 0

WORKDIR /app

COPY go.mod .

COPY *.go ./

RUN go get .

RUN go build -o ./main

FROM registry.access.redhat.com/ubi8/ubi:8.8-1009 AS production

WORKDIR /app

# RUN yum install curl 

COPY --from=builder /app/main .

LABEL syalala="biar beda aja"

EXPOSE 8080

ENTRYPOINT ["./main"]
