.PHONY: generate build

MKDIR_P = mkdir -p

generate_build_folder:
	@${MKDIR_P} builds
	@echo "[OK] Build folder exists"

generate: generate_build_folder
	@go generate ./...
	@echo "[OK] Files added to embed box!"

security:
	@gosec ./...
	@echo "[OK] Go security check was completed!"

build_cross: generate
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/useEmbedded.exe
	@rm *syso
	@echo "[OK] Windows build was created!"
	@env GOOS=darwin GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/useEmbedded_mac
	@echo "[OK] Mac build was created!"
	@env GOOS=linux GOARCH=ppc64le go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/useEmbedded_ppc
	@echo "[OK] Linux(ppc) build was created!"

build: build_cross 
	@env OOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static' -X main.GitCommit=$CI_COMMIT_SHA" -o builds/useEmbedded
	@echo "[OK] Linux build was created!"

run: build
	@./builds/useEmbedded
