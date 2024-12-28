APP=cheetah
LDFLAGS="-s -w -extldflags '-static'"
PLATFORM=$(shell uname -s)
MAIN_PATH="cmd/main.go"

ifeq ($(PLATFORM), linux)
	GOOS=linux
	RM=rm -f
else ifneq ($(findstring MSYS_NT, $(PLATFORM)), )
	GOOS=windows
	RM=del
else
	$(error Unsupported OS: $(PLATFORM))
endif

.PHNOY: all build build-linux clean

all: build

build:
	go env -w GOOS=$(GOOS)
	go build -ldflags=$(LDFLAGS) -o $(APP) $(MAIN_PATH)

build-linux:
	go env -w GOOS=linux
	go build -ldflags=$(LDFLAGS) -o $(APP) $(MAIN_PATH)


clean:
	@RM $(APP)

