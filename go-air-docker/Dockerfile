# 빌드 스테이지
FROM golang:1.23-alpine AS builder

WORKDIR /app

# air live-reload 설치
RUN go install github.com/air-verse/air@latest

# 의존성 파일 복사 및 다운로드
COPY go.mod ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 접근 가능 포트 설정
EXPOSE 8080

# air live-reload 실행
CMD ["air", "-c", ".air.toml"]
