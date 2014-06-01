.PHONY: all

all: c8.js
	make -C c8go --no-print-directory

c8.js: c8.coffee
	coffee -c c8.coffee
