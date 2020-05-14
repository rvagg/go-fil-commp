# portions from https://github.com/filecoin-project/lotus/blob/master/Makefile

UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
all: commp commp_lambda
else
all: commp
endif
.PHONY: all

CLEAN=commp commp_lambda commp_lambda.zip

FFI_PATH:=extern/filecoin-ffi/
FFI_DEPS:=libfilecoin.a filecoin.pc filecoin.h
FFI_DEPS:=$(addprefix $(FFI_PATH),$(FFI_DEPS))

$(FFI_DEPS): .filecoin-install ;

.filecoin-install: $(FFI_PATH)
	$(MAKE) -C $(FFI_PATH) $(FFI_DEPS:$(FFI_PATH)%=%)
	@touch $@

MODULES+=$(FFI_PATH)
BUILD_DEPS+=.filecoin-install
CLEAN+=.filecoin-install

$(MODULES): .update-modules ;

.update-modules:
	git submodule update --init --recursive
	touch $@

CLEAN+=.update-modules

deps: $(BUILD_DEPS)
.PHONY: deps

commp: $(BUILD_DEPS)
	rm -f commp
	go build $(GOFLAGS) -o commp main.go padreader.go
.PHONY: commp

docker:
	cd docker
	docker build -t commp_lambda_build .

ifeq ($(UNAME_S),Linux)
commp_lambda: $(BUILD_DEPS) docker
	rm -f commp_lambda commp_lambda.zip lib
	go build $(GOFLAGS) -o commp_lambda commp_lambda.go
	mkdir lib
	cp -a /usr/lib64/libOpenCL* lib
	zip -r commp_lambda.zip commp_lambda lib
	rm -rf lib
else
commp_lambda:
	$(error can only run commp_lambda on Linux)
endif
.PHONY: commp_lambda

clean:
	rm -rf $(CLEAN) $(BINS)
	-$(MAKE) -C $(FFI_PATH) clean
.PHONY: clean
