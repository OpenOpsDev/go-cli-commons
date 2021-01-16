DOCKER_TARGET?=ci

test:
	cd ${module} && go test -v github.com/openopsdev/go-cli-commons/...

lint:
	golangci-lint run

local-module:
	echo "replace github.com/openopsdev/go-cli-commons/$(module) => ../go-cli-commons/$(module)" >> go.mod

remote-module:
	sed -i -e "/replace\ github\.com\/openopsdev\/go-cli-commons\/$(module)\ =>/d" go.mod