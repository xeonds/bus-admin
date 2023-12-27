PROJECT_NAME=bus-admin
CONFIG_FILE	=config.yaml
BINARY_DIR	=bus-admin

all: web server
	if [ -d "build" ];then rm -rf build; fi \
	&& mkdir build && cp -r web/dist build/ && cp server/server build/ && cp -r server/resource build/resource 

web:
	@cd web/ && if [ -d "dist" ];then rm -rf dist; fi \
	&& npm set registry http://mirrors.cloud.tencent.com/npm/ && npm install && npm build

server:
	@cd server/ && if [ -f "server" ];then rm server; fi \
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy \
	&& go build -ldflags "-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n') -X main.Version=${TAGS_OPT}" -v

run: server
	cd $(BINARY_DIR) && ./$(PROJECT_NAME) -c $(CONFIG_FILE)