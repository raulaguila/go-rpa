FROM golang:alpine

RUN apk add gcompat

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy folders and files to container
COPY . .

# Run binary
CMD ["./gorpa"]
