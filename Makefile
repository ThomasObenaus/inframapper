.DEFAULT_GOAL				:= all
name 								:= "inframapper"

all: vendor build tools cover finish

.PHONY: test
test: generate.mocks
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests"
	@go test ./tfstate ./trace ./aws ./mappedInfra ./terraform -v

.PHONY: cover
cover: generate.mocks
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests + coverage"
	@go test ./tfstate ./trace ./aws ./mappedInfra ./terraform -v -covermode=count -coverprofile=coverage.out

cover.upload:
	# for this to get working you have to export the repo_token for your repo at coveralls.io
	# i.e. export INFRA_MAPPER_COVERALLS_REPO_TOKEN=<your token>
	@${GOPATH}/bin/goveralls -coverprofile=coverage.out -service=circleci -repotoken=${INFRA_MAPPER_COVERALLS_REPO_TOKEN}
	

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

#------------------
#-- Tools
#------------------
tools:
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls	


#------------------
#-- Generate
#------------------
generate:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> generate String() for enums (golang.org/x/tools/cmd/stringer is required for this)"
	@go get golang.org/x/tools/cmd/stringer
	@stringer -type=ResourceType terraform/resource_type.go
	@stringer -type=ResourceType aws/resource_type.go
	@stringer -type=ResourceType mappedInfra/resource_type.go

generate.mocks:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> generate mocks (github.com/golang/mock/gomock is required for this)"
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@mockgen -source=aws/iface/aws_IF.go -destination test/mock_aws_iface/mock_aws_IF.go 
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