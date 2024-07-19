# 使用官方 Go 映像作為基礎
FROM golang:1.22-alpine AS builder

# 設置工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製源代碼
COPY . .

# 添加調試信息
RUN pwd && ls -la

# 構建應用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用輕量級的 alpine 映像作為最終映像
FROM alpine:latest

# 安裝 CA 證書
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 從構建階段複製二進制文件
COPY --from=builder /app/main .

# 複製可能需要的其他文件（如配置文件）
# COPY --from=builder /app/config.yaml .

# 暴露應用端口
EXPOSE 8080

# 運行應用
CMD ["./main"]
