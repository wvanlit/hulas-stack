# Clean up previous build
rm -rf ./bin
mkdir bin

# Clean up previous docker image HULAS
docker stop hulas-server
docker rm hulas-server
docker rmi -f hulas-server

# Build go binary
go build -o ./bin/ ./server/server.go

# Build docker image
docker build -t hulas-server .