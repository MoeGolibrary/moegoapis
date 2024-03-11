.PHONY: format-lint format lint install pb

format-lint: format lint				##@Style: run all the style steps

format:													##@Style: format the code style
	@buf format -w moego
	@echo "Format done!"

lint: 													##@Style: lint the code style
	@buf lint moego
	@buf breaking --against '.git#branch=origin/production,subdir=moego' moego
	@echo "Lint done!"

install:												##@Tools: link buf modules to local, you should run this again after add new dependencies in buf.yaml
	@rm -rf third_party
	@mkdir -p third_party
	@yq -r '.deps[]' buf.yaml | xargs -I % buf export % -o third_party
	@echo "Generate deps directory done!"

pb:														##@Build steps: generate protobuf
	@buf generate
	@buf build -o tmp/moego.bin --as-file-descriptor-set
	@rm -rf out
	@mv tmp/go/github.com/MoeGolibrary/moegoapis/genproto out
	@mv tmp/moego.bin out/moego.bin
	@rm -rf tmp
	@echo "Generate protobuf files done!"