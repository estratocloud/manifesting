.PHONY: up

help:
	@echo "Available options:"
	@echo "    make up"

up:
	docker rm -f manifesting-docs > /dev/null 2>&1
	docker build . -t manifesting-docs-image
	docker run --interactive --volume $(shell pwd):/app --entrypoint="" manifesting-docs-image bundle install
	docker run --interactive --volume $(shell pwd):/app --detach -p"4014:4000" --name manifesting-docs manifesting-docs-image
	xdg-open http://localhost:4014/manifesting/
