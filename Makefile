.PHONY: format-lint format lint install

format-lint: format lint				##@Style: run all the style steps

format:													##@Style: format the code style
	@buf format -w moego
	@echo "Format done!"

lint: 													##@Style: lint the code style
	@buf lint moego
	@echo "Lint done!"

install:												##@Tools: link buf modules to local, you should run this again after add new dependencies in buf.yaml
	@rm -rf third_party
	@mkdir -p third_party
	@yq -r '.deps[]' moego/buf.yaml | xargs -I % buf export % -o third_party
	@echo "Generate deps directory done!"