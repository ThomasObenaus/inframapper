.DEFAULT_GOAL				:= all
name 								:= "terrastate"

all: vendor build cover finish

.PHONY: test
test: generate
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests"
	@go test ./tfstate ./trace ./aws ./mappedInfra ./terraform -v

.PHONY: cover
cover: generate
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests + coverage"
	@go test ./tfstate ./trace ./aws ./mappedInfra ./terraform -cover -v

#-----------------
#-- build
#-----------------
build: generate
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Build the $(name)"
	@go build -o $(name) .

#------------------
#-- dependencies
#------------------
depend.update:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> updating dependencies from Gopkg.lock"
	@dep ensure -update

depend.install:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> install dependencies as listed in Gopkg.toml"
	@dep ensure

vendor: depend.install

generate:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> generate String() for enums (golang.org/x/tools/cmd/stringer is required for this)"
	@go get golang.org/x/tools/cmd/stringer
	@stringer -type=Type terraform
	@stringer -type=Type aws
	@stringer -type=Type mappedInfra
	
	@echo "--> generate mocks (github.com/golang/mock/gomock is required for this)"
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@mockgen -source=aws/infra_loader.go -destination test/mock_aws/mock_infra_loader.go
	@mockgen -source=aws/infra.go -destination test/mock_aws/mock_infra.go
	@mockgen -source=aws/resource.go -destination test/mock_aws/mock_resource.go
	@mockgen -source=terraform/resource.go -destination test/mock_terraform/mock_resource.go
	@mockgen -source=terraform/infra.go -destination test/mock_terraform/mock_infra.go
	@mockgen -source=mappedInfra/infra.go -destination test/mock_mappedInfra/mock_infra.go
	@mockgen -source=mappedInfra/mapper.go -destination test/mock_mappedInfra/mock_mapper.go
	@mockgen -source=mappedInfra/resource.go -destination test/mock_mappedInfra/mock_resource.go
	@mockgen -source=tfstate/iface/s3_downloader.go -destination test/mock_tfstate_iface/mock_s3_downloader.go
	@mockgen -source=tfstate/iface/state_loader.go -destination test/mock_tfstate_iface/mock_state_loader.go

run: build
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run ${name}"
	@./${name}

finish:
	@echo "=================================================================================="