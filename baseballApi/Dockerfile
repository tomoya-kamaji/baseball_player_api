FROM golang:1.19-alpine

WORKDIR /app/baseballApi

COPY . .

# RUN apk add --update --no-cache \
#       libc-dev \
#       gcc \
#       git && \
#       GOOS=linux go build -a -o server


RUN apk --no-cache add ca-certificates tzdata
RUN ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

CMD ["go", "run", "cmd/main.go"]
