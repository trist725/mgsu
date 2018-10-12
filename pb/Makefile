PACKAGE_PATH=github.com/trist725/mgsu/pb
WORK_DIR=$(shell pwd)
########################################################################################################################
.PHONY: all \
		make_go \
		gen_static_data_go \
		protoc-gen-enum-go \
		protoc-gen-mgo-go \
########################################################################################################################
all: make_go \
	gen_static_data_go \
	protoc-gen-enum-go \
	protoc-gen-mgo-go \
########################################################################################################################
clean:
	rm -rf $(WORK_DIR)/pkg
	rm -f $(GOPATH)/bin/make_go.sh
	rm -f $(GOPATH)/bin/gen_static_data_go
	rm -f $(GOPATH)/bin/protoc-gen-enum-go
	rm -f $(GOPATH)/bin/protoc-gen-mgo-go
########################################################################################################################
make_go:
	cp ./make_go.sh $(GOPATH)/bin
	chmod +x $(GOPATH)/bin/make_go.sh

gen_static_data_go:
	@echo $(shell date "+%F %R:%S")
	go install $(PACKAGE_PATH)/gen_static_data_go

protoc-gen-enum-go:
	@echo $(shell date "+%F %R:%S")
	go install $(PACKAGE_PATH)/protoc-gen-enum-go
	cd $(WORK_DIR)/protoc-gen-enum-go; go generate ./enum; go test

protoc-gen-rpc-go:
	@echo $(shell date "+%F %R:%S")
	go install $(PACKAGE_PATH)/protoc-gen-rpc-go
	cd $(WORK_DIR)/protoc-gen-rpc-go; go generate ./rpc; go test

protoc-gen-mgo-go:
	@echo $(shell date "+%F %R:%S")
	go install $(PACKAGE_PATH)/protoc-gen-mgo-go
	cd $(WORK_DIR)/protoc-gen-mgo-go; go generate ./model; go test
########################################################################################################################
