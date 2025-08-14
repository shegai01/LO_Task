FROM golang:1.21-alpine as builder
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app
EXPOSE 8000
RUN GOOS=linux GOARCH=amd64 go build -o taskManager

FROM scratch
COPY --from=builder /app/taskManager /usr/bin/taskManager

ENTRYPOINT [ "/usr/bin/taskManager" ]