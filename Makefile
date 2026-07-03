SRC := src

.PHONY: all run 

all: run

run:
	cd $(SRC) && go run .

