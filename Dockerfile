FROM golang:1.26-trixie

WORKDIR /app

RUN apt-get update && apt-get install -y curl \
    && curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

# Copy go mod files first for better caching
COPY go.mod go.sum* ./
RUN go mod download

CMD ["tail", "-f", "/dev/null"]
