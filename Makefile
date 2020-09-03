.DEFAULT_GOAL := help
TARGETS = linux darwin windows

.EXPORT_ALL_VARIABLES:

.PHONY:

all:

compile: ## Compile Go code
	@for target in $(TARGETS); do env GOOS=$$target go build -o roles/go-role/library/calendar_$$target *.go; done

test-manual-go: ## Test Go module with arguments file
	@go run *.go test/args.json

test-manual-bash: ## Test Go module with arguments file via Bash module.
	@roles/go_role/library/go_run 'test/args.sh'

test-ansible: ## Test module with Ansible
	@ansible-playbook test-module.yml

test-go: ## Test module with Ansible and Go
	@ansible-playbook test-module.yml --extra-vars "go=true"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'