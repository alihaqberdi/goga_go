
create-network:
	@docker network create my-network || true


run_db: create-network
	@docker run -d --name postgres-container \
		--network my-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-p 5432:5432 \
		postgres

run:
	docker build -t my-app .
	docker run -d \
      --name my-app \
      --restart=always \
      --network my-network \
      -p 8888:8888 \
      my-app



