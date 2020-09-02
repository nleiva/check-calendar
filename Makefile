.DEFAULT_GOAL := help
TARGETS = linux darwin windows

.EXPORT_ALL_VARIABLES:

.PHONY:

all:

compile: ## Compile Go code
	@for target in $(TARGETS); do env GOOS=$$target go build -o roles/go-role/library/calendar_$$target *.go; done

test-manual: ## Test module with arguments file
	go run *.go args.json

test-ansible: ## Test module with ansible
	ansible-playbook test-module.yml

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'