SERVICE := go-hotreload
bold := $(shell tput bold)
sgr0 := $(shell tput sgr0)

.PHONY: help

help:
	@printf "\
	$(SERVICE)\n\
	\n\
	$(bold)SYNOPSIS$(sgr0)\n\
    make [COMMANDS]\n\
	\n\
	$(bold)COMMANDS$(sgr0)\n\
		$(bold)help$(sgr0)\n\
		  Display help\n\n\
		$(bold)dev$(sgr0)\n\
		  Build images and start service for development\n\n\
		$(bold)dev-build$(sgr0)\n\
		  Force build images and start service for development\n\n\
	"
	
dev:
	docker-compose up

dev-build:
	docker-compose up --build