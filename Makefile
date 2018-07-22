.DEFAULT_GOAL				:= all
name 								:= "terrastate"

all: vendor build cover run finish

.PHONY: test
test: generate
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests"
	@go test ./tfstate ./trace ./aws ./mappedInfra -v

.PHONY: cover
cover: generate
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests + coverage"
	@go test ./tfstate ./trace ./aws ./mappedInfra -cover -v

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
	@echo "--> generate String() for enums"
	@stringer -type=Type terraform
	@stringer -type=Type aws

run: build
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run ${name}"
	@./${name}

finish:
	@echo "=================================================================================="