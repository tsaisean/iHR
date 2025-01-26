docker_run_staging:
	docker build -t ihr .
	docker run -p 8080:8080 ihr

docker_run_local:
	docker-compose up

docker_rerun_local:
	docker-compose up --build

docker_clean:
	docker system prune -f

docker_images:
	docker images

docker_rm:
	docker image rm ihr

docker_vm_ls:
	docker volume ls

docker_vm_rm_redis:
	docker volume rm ihr_redis-data

build:
	go build ./...

test:
	go test ./...

gen_secret:
	openssl rand -base64 32