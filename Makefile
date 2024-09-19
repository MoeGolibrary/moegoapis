.PHONY: deps pb

format-lint: format lint				##@Style: run all the style steps

format:													##@Style: format the code style
	@buf format -w .
	@echo "Format done!"

lint: 													##@Style: lint the code style
	@buf lint .
	@buf breaking --against '.git#branch=origin/production,subdir=moego' moego
	@echo "Lint done!"

deps:												##@Tools: link buf modules to local, you should run this again after add new dependencies in buf.yaml
	@rm -rf deps
	@mkdir -p deps
	@yq -r '.deps[]' buf.yaml | xargs -I % buf export % -o deps
	@echo "Generate deps directory done!"

pb:														##@Build steps: generate protobuf
	@buf generate
	@rm -rf genproto
	@mv out/go/github.com/MoeGolibrary/moegoapis/genproto genproto
	@rm -rf out
	@buf build -o genproto/moego.bin --as-file-descriptor-set
	@echo "Generate protobuf files done!"
