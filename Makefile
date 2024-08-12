.PHONY: clean all

wc:
	go build ./cmd/wc

all: wc

clean:
	rm -f wc
