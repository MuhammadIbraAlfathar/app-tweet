FROM golang:1.22.7 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app

COPY go.mod go.sum ./

# Unduh dan cache dependensi
RUN go mod download

# Copy semua kode proyek ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o server cmd/server/main.go


FROM gcr.io/distroless/base-debian11

# Copy hasil build dari stage sebelumnya ke stage ini
COPY --from=builder /app/server .

# Tentukan perintah untuk menjalankan aplikasi saat container berjalan
CMD ["./server"]