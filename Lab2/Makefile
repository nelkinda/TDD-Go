.PHONY: all
all: test hello

.PHONY: test
test::
	go test

hello:
	go build -o hello

.PHONY:
clean::
	go clean

-include User.mk
-include ../User.mk
-include ~/User.mk
