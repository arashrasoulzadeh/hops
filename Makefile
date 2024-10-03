# Variables
PACKAGE_NAME = hops
VERSION = 0.0.1
BUILD_DIR = ./build/$(PACKAGE_NAME)
DEBIAN_DIR = $(BUILD_DIR)/DEBIAN
BINARY_DIR = $(BUILD_DIR)/usr/local/bin
CONTROL_FILE = $(DEBIAN_DIR)/control

# Architectures
ARCHS = amd64 arm64

# Detect OS
UNAME_S := $(shell uname -s)

# Default target to build everything
all: debs

# Step 1: Build Go binaries for multiple architectures and macOS
build:
	@echo "Compiling Go binaries for architectures..."
	@for arch in $(ARCHS); do \
		GOARCH=$$arch GOOS=linux go build -o build_temp_$$arch; \
	done
	@if [ "$(UNAME_S)" = "Darwin" ]; then \
		echo "Compiling Go binary for macOS..."; \
		GOARCH=amd64 GOOS=darwin go build -o build_temp_mac; \
	fi

# Step 2: Copy binaries and create .deb packages for each architecture
debs: build
	@for arch in $(ARCHS); do \
		$(MAKE) deb ARCH=$$arch; \
	done

# Step 3: Create the control file and .deb package for the specified architecture
deb:
	@echo "Creating .deb package for architecture: $(ARCH)..."

	# Create package structure
	mkdir -p $(DEBIAN_DIR) $(BINARY_DIR)

	# Copy binary to build folder
	cp build_temp_$(ARCH) $(BINARY_DIR)/$(PACKAGE_NAME)

	# Create control file with Version field
	@echo "Creating control file for $(ARCH)..."
	@echo "Package: $(PACKAGE_NAME)\nVersion: $(VERSION)\nSection: $(PACKAGE_NAME)\nPriority: optional\nArchitecture: $(ARCH)\nMaintainer: Arash Rasoulzadeh <arashrasoulzadeh@gmail.com>\nDescription: Helper for devops\nHomepage: http://www.meetarash.ir" > $(CONTROL_FILE)

	# Build .deb package
	@echo "Building .deb package for $(ARCH)..."
	dpkg-deb --build $(BUILD_DIR) $(PACKAGE_NAME)_$(ARCH).deb

# Step 4: Clean up binaries and build folders
clean:
	@echo "Cleaning up..."
	rm -rf build_temp_* $(BUILD_DIR)

.PHONY: all build deb debs clean
