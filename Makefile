# $(API_DOCKER_DIR)
API_DOCKER_DIR = ./baseballApi/
NGINX_DOCKER_DIR = ./k8s/baseball/api/docker/nginx/
K8S_DIR = ./k8s/baseball/api/base

.PHONY: docker-build k8s-apply

docker-build:
	eval $$(minikube docker-env) ;\
	docker build -t baseball-api:latest  $(API_DOCKER_DIR)
	docker build -t my-nginx:latest $(NGINX_DOCKER_DIR)

minikube-apply: docker-build
	kubectl apply -k $(K8S_DIR)
	minikube service baseball-api