# Start from golang v1.16 base image to have access to go modules
FROM golang:1.16

# create a working directory
WORKDIR /app

# we will be expecting to get API_PORT as arguments
ARG API_PORT

# Fetch dependencies on separate layer as they are less likely to
# change on every build and will therefore be cached for speeding
# up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# copy source from the host to the working directory inside
# the container
COPY . .

RUN chmod +x ./scripts/entrypoint.sh

# This container exposes API_PORT from .env to the outside world
EXPOSE ${API_PORT}

ENTRYPOINT ["./scripts/entrypoint.sh"]

