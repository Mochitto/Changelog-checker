DIST = ./dist/main
MAIN_FILE = ./cmd/changelog-checker/main.go


start:
	@echo -e "⏳ Building your go app..."
	@go mod tidy
	@go build -o ${DIST} ${MAIN_FILE}
	@echo -e "✅ Done building!\n\n"
	@${DIST}
