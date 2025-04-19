https://github.com/k5prrr/fastStartGO

# Install
1. micro .gitignore
2. sh ./scripts/installModules.sh


# Start
go fmt ./... && clear && go run ./cmd/App/main.go


# Docker
GOOS=linux go build -o build/appGO cmd/App/main.go
docker build -t NAME1:v1 .
docker run -it --rm NAME1:v1 ls -l /build
5


sh scripts/start.sh

http://localhost:8081/testSpeed




