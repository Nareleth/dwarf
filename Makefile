TARGET 	:= dwarf
SRC 	:= src
BUILD 	:= $(CURDIR)/build

.PHONY: all build cleanrun 

all: run

build: clean
	mkdir $(BUILD)
	cd $(SRC) && go build -o $(BUILD)/$(TARGET) .

run:
	cd $(SRC) && go run .

clean:
	rm -rf $(BUILD) || true
