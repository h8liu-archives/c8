.PHONY: all

all: c8.js fs.js

c8.js: c8.coffee
	coffee -c c8.coffee

fs.js: fs.coffee
	coffee -c fs.coffee
