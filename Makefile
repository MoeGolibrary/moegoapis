.PHONY: deps pb watch install-hooks

format-lint: format lint				##@Style: run all the style steps

format:													##@Style: format the code style
	@buf format -w .
	@echo "Format done!"

lint: 													##@Style: lint the code style
	@buf lint .
	@buf breaking --against '.git#branch=origin/production' .
	@echo "Lint done!"

deps:												##@Tools: link buf modules to local, you should run this again after add new dependencies in buf.yaml
	@rm -rf deps
	@mkdir -p deps
	@yq -r '.deps[]' buf.yaml | xargs -I % buf export % -o deps
	@echo "Generate deps directory done!"

pb:														##@Build steps: generate protobuf
	@buf generate
	@rm -rf genproto docs/openapi
	@mv out/go/github.com/MoeGolibrary/moegoapis/genproto genproto
	@mv out/openapi docs/
	@find docs/openapi -type f ! -name '*_service.swagger.json' -delete || true
	@rm -rf out
	@buf build -o genproto/moego.bin --as-file-descriptor-set
	@echo "Generate protobuf files done!"

watch:										##@Development: watch proto files and run 'make pb' on change (requires entr: brew install entr)
	@command -v entr >/dev/null 2>&1 || { echo "Install entr first: brew install entr"; exit 1; }
	@echo "Watching moego/**/*.proto - press Ctrl+C to stop"
	@find moego -name '*.proto' | entr -c make pb

install-hooks:							##@Development: install git hooks for auto 'make pb' on commit
	@git config core.hooksPath .githooks
	@chmod +x .githooks/pre-commit
	@echo "Git hooks installed. Proto changes will trigger 'make pb' on commit."
