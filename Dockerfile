FROM golang AS builder
WORKDIR /app
COPY . .
ARG USE_EXAMPLE

RUN go mod download  
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot ./cmd
RUN if [ "$USE_EXAMPLE" = "true" ]; \
        then cp /app/config.yaml.example /app/config.yaml; \
    else \
        if [ ! -f /app/config.yaml ]; \
    # default UpdateHandlerConfig will be used
            then touch /app/config.yaml; \ 
        fi; \
    fi;

FROM alpine:latest
WORKDIR /app 
COPY --from=builder /app/bot .
COPY --from=builder /app/config.yaml . 
RUN apk --no-cache add ca-certificates  
CMD ["./bot"]