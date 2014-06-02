.PHONY: all deploy

all: www/c8.js
	make -C c8go --no-print-directory

www/c8.js: c8.coffee
	coffee -c -o www c8.coffee

deploy:
	rm ~/projects/github.io/c8/*
	cp www/* ~/projects/github.io/c8/.
