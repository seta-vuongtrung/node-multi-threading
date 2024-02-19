MAKEFILE_LIST=Makefile

.PHONY: get-heavy
get-heavy:
	curl http://localhost:3000/heavy

.PHONY: load-test
load-test:
	loadtest -n 1200 -c 200 -k http://localhost:3000/heavy


