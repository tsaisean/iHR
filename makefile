docker_run:
	docker build -t ihr .
	docker run -p 8080:8080 ihr

docker_clean:
	docker system prune -f

docker_rm:
	docker image rm ihr

gen_secret:
	openssl rand -base64 32