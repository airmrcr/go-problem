MODULE_DIRS = . ./zap

all: download tidy format build test bench

download:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod download) &&) true

tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy) &&) true

format:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go fmt) &&) true

build:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go build -v ./...) &&) true

test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race -v ./...) &&) true

bench:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -run=XXX -bench=. -v ./...) &&) true

update:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go get -u all) &&) true
