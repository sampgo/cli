.PHONY: clean
default: build

clean:
	rm -f ~/bin/sampgo

build:
	go build -o ~/bin/sampgo