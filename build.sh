set -x
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
docker build -f ./Dockerfile . -t auth:`git describe --always`
rm auth