NAME=bus-admin
BINARY_DIR=build
VERSION=1.0.0
BUILDTIME=$(shell date -u)
GOBUILD=cd server && go build -ldflags '-X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'
FRONTBUILD=cd web && pnpm i && vite build --outDir=../$(BINARY_DIR)/dist

all: linux-amd64 windows-amd64 frontend

linux-amd64: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../$(BINARY_DIR)/$(NAME)-$@-$(VERSION)

linux-arm64: 
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o ../$(BINARY_DIR)/$(NAME)-$@-$(VERSION)

windows-amd64: 
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../$(BINARY_DIR)/$(NAME)-$@-$(VERSION).exe

windows-amd64-v3: 
	GOOS=windows GOARCH=amd64 GOAMD64=v3 $(GOBUILD) -o ../$(BINARY_DIR)/$(NAME)-$@-$(VERSION).exe

frontend:
	$(FRONTBUILD)

clean:
	rm -rf $(BINARY_DIR)/bus-admin*
	rm -rf $(BINARY_DIR)/dist

run: linux-amd64
	cd $(BINARY_DIR) && ./$(NAME)-linux-amd64-$(VERSION)
