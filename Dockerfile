# Start from golang v1.16 base image to have access to go modules
FROM golang:1.16

# create a working directory
WORKDIR /app

# Fetch dependencies on separate layer as they are less likely to
# change on every build and will therefore be cached for speeding
# up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy source from the host to the working directory inside
# the container
COPY . .

# Build
RUN go build -o main cmd/api/main.go

# Assign executive mode for entrypoint
RUN chmod +x ./scripts/entrypoint.sh

# This container exposes port to the outside world
EXPOSE 8080

ENTRYPOINT ["./scripts/entrypoint.sh"]

