FROM golang AS builder
WORKDIR /app
COPY . .

RUN go mod download  
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot ./cmd

FROM alpine:latest
WORKDIR /app 
COPY --from=builder /app/bot .
COPY --from=builder /app/config.yaml . 
RUN apk --no-cache add ca-certificates  
CMD ["./bot"]