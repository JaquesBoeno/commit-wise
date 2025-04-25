# Final binary name
APP_NAME := commitwise

# Output directory
BIN_DIR := bin

# Targets
.PHONY: build clean install uninstall

build:
	@echo "🔨 Building CLI..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) .

clean:
	@echo "🧹 Cleaning up binary..."
	rm -rf $(BIN_DIR)/$(APP_NAME)

install: build
	@echo "📦 Installing $(APP_NAME) to /usr/local/bin..."
	sudo cp $(BIN_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)

uninstall:
	@echo "🧹 uninstalling $(APP_NAME) from /usr/local/bin"
	sudo rm /usr/local/bin/$(APP_NAME)