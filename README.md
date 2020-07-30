# cmkit
FOR LEARNING ONLY

# START SERVER
cd cmd
go build -ldflags "-w -s -v -X main.VERSION=1.0.0 -X 'main.BUILD_TIME=$(date)' -X 'main.GO_VERSION=$(go version)'"
./cmkit

# START WEB
cd web
npm start