GITHUB_OS_LIST := ubuntu-latest

include ./.bootstrap/makesystem.mk

ifeq ($(MAKESYSTEM_FOUND),1)
include $(MAKESYSTEM_BASE_DIR)/go.mk
endif
