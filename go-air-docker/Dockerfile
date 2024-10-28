# 빌드 스테이지
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 의존성 파일 복사 및 다운로드
COPY go.mod ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 접근 가능 포트 설정
EXPOSE 8080

# 실행 스테이지
FROM alpine:latest

WORKDIR /app

# 빌더에서 실행 파일만 복사
COPY --from=builder /app/main .

# 실행
CMD ["./main"]
