# portions from https://github.com/filecoin-project/lotus/blob/master/Makefile

all: commp
.PHONY: all

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
	go build $(GOFLAGS) -o commp .
.PHONY: commp

clean:
	rm -rf $(CLEAN) $(BINS)
	-$(MAKE) -C $(FFI_PATH) clean
.PHONY: clean
