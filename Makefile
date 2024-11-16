
reload: git_reload

git_reload:
	git pull
	$(MAKE) docker_reload

docker_reload:
	go build cmd/main.go
	docker stop my-app
	docker rm my-app
	docker rmi my-app
	docker build -t my-app .
	docker run -d \
      --name my-app \
      --restart=always \
      --network my-network \
      -v /my-docker-data/my-uploads:/app/uploads \
      -p 8080:8080 \
      my-app





