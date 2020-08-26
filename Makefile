.PHONY: generate build

MKDIR_P = mkdir -p

generate_build_folder:
	@${MKDIR_P} builds
	@echo "[OK] Build folder exists"

generate:
	@go generate ./...
	@echo "[OK] Files added to embed box!"

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

build: generate generate_build_folder
	#@env OOS=linux GOARCH=arm go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/isv_embed_arm
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/isv_embed.exe
	@rm *syso
	@env OOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/isv_embed 
	#@env GOOS=darwin GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/isv_embed_mac
	@echo "[OK] App binary was created!"

run: build
	@./builds/isv_embed
