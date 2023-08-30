# Dockerfile

# node ,go 등 runtime 및 버전 설정
FROM golang:1.19-alpine AS builder


# move to working dir (/build)
WORKDIR /build

# copy 하고 dependency 를 go mod를 이용해 다운로드
COPY go.mod go.sum ./
RUN go mod download 

# 전체 코드를 (.) 카피해서 container(.) 안으로
COPY . .

# Set necessary environment variables needed for our image 
# CGO_ENABLED와 -ldflags="-s -w"로 빌드해서 바이너리 사이즈를 줄인다
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# 그리고 api 서버 빌드
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch

# Copy binary and config files from /build 
# to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Export necessary port.
EXPOSE 5000

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]