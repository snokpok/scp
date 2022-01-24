REGION_AWS = us-west-2
AWS_ACCOUNT_ID = 678445257595
AWS_ECR_URI = $(AWS_ACCOUNT_ID).dkr.ecr.$(REGION_AWS).amazonaws.com
DOCKER_REPO = scp-go
DOCKER_TAG = latest
DOCKER_REPO_TAG = $(DOCKER_REPO):$(DOCKER_TAG)
PATH_TO_PEM_KEY = main-scp-server-key.pem

all:
	make login-docker-aws-ecr
	cd ./scp-go && make docker-build
	make tag-publish-aws-ecr
	make push-aws-ecr

login-docker-aws-ecr:
	aws ecr get-login-password --region $(REGION_AWS)  | docker login --username AWS --password-stdin $(AWS_ECR_URI)

tag-publish-aws-ecr:
	docker tag $(DOCKER_REPO_TAG) $(AWS_ECR_URI)/$(DOCKER_REPO_TAG)

push-aws-ecr:
	docker push $(AWS_ECR_URI)/$(DOCKER_REPO_TAG)

ssh-ec2-instance:
	ssh -i "$(PATH_TO_PEM_KEY)" ec2-user@ec2-35-160-69-69.us-west-2.compute.amazonaws.com