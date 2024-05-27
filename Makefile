.PHONY: build run clean_server stop

build:
	docker-compose build

run:
	docker-compose up

build_and_run: 
	build run

clean_server:
	rm -rf./server

stop:
	docker-compose down
