.PHONY: all
all: test

.PHONY: test
test::
	go test --cover

.PHONY:
clean::
	go clean

-include User.mk
-include ../User.mk
-include ~/User.mk
