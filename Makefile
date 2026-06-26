# Settings
SPEC_URL     := https://api.clo.ru/openapi.json
API_URL      := https://api.clo.ru/
SPEC_FIXED   := openapi.final.json
JQ_SCRIPT    := spec/fix.jq
CONFIG       := spec/config.yaml
SDK_OUT      := clo_gen.go
MAJOR        := v3

# Tools
JQ           := jq
OAPI_CODEGEN := oapi-codegen
CURL         := curl -sSfL

.PHONY: all
all: generate clean-spec ## Full cycle (generate + cleanup)

.PHONY: fetch-and-fix
fetch-and-fix: ## 1. Download and transform the spec
	@echo "--- [1/2] Fetching and fixing spec ---"
	@$(CURL) $(SPEC_URL) | $(JQ) --arg api_url "$(API_URL)" -f $(JQ_SCRIPT) > $(SPEC_FIXED)

.PHONY: generate
generate: fetch-and-fix ## 2. Generate the SDK and tidy
	@echo "--- [2/2] Generating Go SDK ---"
	@$(OAPI_CODEGEN) -config $(CONFIG) $(SPEC_FIXED)
	@sed -i.bak 's/WithHTTPClient/WithHTTPClientDoer/g' $(SDK_OUT) && rm -f $(SDK_OUT).bak
	@go mod tidy
	@echo "Success: $(SDK_OUT) generated."

.PHONY: release
release: generate ## Tag the next independent-semver patch off the latest $(MAJOR).* tag
	@echo "--- Determining next version ---"
	@LATEST_TAG=$$(git tag -l "$(MAJOR).*" | sort -V | tail -n1); \
	if [ -z "$$LATEST_TAG" ]; then \
		NEW_TAG="$(MAJOR).0.0"; \
		echo "No existing $(MAJOR).* tags. Starting with $$NEW_TAG"; \
	else \
		echo "Found latest tag: $$LATEST_TAG"; \
		MAJ=$$(echo $$LATEST_TAG | cut -d. -f1); \
		MIN=$$(echo $$LATEST_TAG | cut -d. -f2); \
		PATCH=$$(echo $$LATEST_TAG | cut -d. -f3); \
		NEW_TAG="$$MAJ.$$MIN.$$(($$PATCH + 1))"; \
		echo "Incrementing to $$NEW_TAG"; \
	fi; \
	git tag -a "$$NEW_TAG" -m "Auto-release $$NEW_TAG"; \
	git push origin "$$NEW_TAG"

.PHONY: clean-spec
clean-spec: ## Remove the temporary spec file
	@rm -f $(SPEC_FIXED)

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'