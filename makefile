.PHONY: all

all: www/c8.js
	make -C c8go --no-print-directory

www/c8.js: c8.coffee
	coffee -c -o www c8.coffee
