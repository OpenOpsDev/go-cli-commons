DOCKER_TARGET?=ci

test:
	cd ${module} && go test -v github.com/openopsdev/go-cli-commons/...

lint:
	go get -u golang.org/x/lint/golint
	test -z "$$(golint ./...)"

local-module:
	echo "replace github.com/openopsdev/go-cli-commons/$(module) => ../go-cli-commons/$(module)" >> go.mod

remote-module:
	sed -i -e "/replace\ github\.com\/openopsdev\/go-cli-commons\/$(module)\ =>/d" go.mod