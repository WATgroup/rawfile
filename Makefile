#
.PHONY: dummy

GOLINT:=$(shell command -v revive)


GOLINT_OPTS=-formatter stylish


lint:	tidy
	if [ -n $(GOLINT) ]; then $(GOLINT) $(GOLINT_OPTS); fi

-include /usr/share/go-common/gocommon.mk
