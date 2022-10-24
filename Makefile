image:
	docker build -t task-clean-arch:latest -f Dockerfile .

container:
	docker run --rm -p 8899:8899 --env-file ./local.env --name task-clean-arch task-clean-arch:latest