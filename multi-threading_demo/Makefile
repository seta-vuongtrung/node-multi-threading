MAKEFILE_LIST=Makefile

.PHONY: non-blocking
non-blocking:
	curl --get http://localhost:3000/non-blocking

.PHONY: blocking
blocking:
	curl --get http://localhost:3000/blocking

.PHONY: test_time
test_time:
	time curl --get http://localhost:3000/blocking

