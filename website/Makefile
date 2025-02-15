###############################################################################
# Configuration and Variables                                                 #
###############################################################################
APP_TARGET = ./cmd/website/main.go

GO_PROD_FLAGS = -ldflags="-s -w"
GO_DEV_FLAGS = 

HUGO_BUILD_FLAGS = --gc --cleanDestinationDir
HUGO_SERVER_FLAGS = $(HUGO_BUILD_FLAGS) --disableFastRender --watch

GO_CLEAN_TARGETS = website
HUGO_CLEAN_TARGETS = public/ resources/ .hugo_build.lock hugo_stats.json

.PHONY: all dev clean ui-build ui-devbuild ui-server app-build app-server

###############################################################################
# Default Targets                                                             #
###############################################################################

all: clean ui-build app-build

dev: clean ui-devbuild app-server

clean:
	rm -rf $(GO_CLEAN_TARGETS) $(HUGO_CLEAN_TARGETS)

###############################################################################
# UI Build and Serve                                                          #
###############################################################################

ui-build:
	HUGO_ENVIRONMENT=production hugo build $(HUGO_BUILD_FLAGS)

ui-devbuild:
	HUGO_ENVIRONMENT=development hugo build $(HUGO_BUILD_FLAGS)

ui-server:
	hugo server $(HUGO_SERVER_FLAGS)

###############################################################################
# App Build and Server                                                        #
###############################################################################

app-build:
	go build -o website $(GO_PROD_FLAGS) $(APP_TARGET)

app-server:
	go run $(GO_DEV_FLAGS) $(APP_TARGET)