FROM golang:1.21.0 as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

# CMD ["go", "run", "/usr/src/app/cmd/read-it", "."]


RUN CGO_ENABLED=0 GOOS=linux go build -o go_project ./cmd/read-it

# Use a smaller image to run the app
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN ls -la

# Copy the pre-built binary file from the previous stage
COPY --from=builder /usr/src/app/go_project .
COPY --from=builder /usr/src/app/pkg/read-it ./migrations

CMD [ "./go_project" ]
