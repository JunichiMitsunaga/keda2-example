VERSION := "0.0.1"

go-build:
	@go build -o ./components-mock ./example/main.go

run:
	@./components-mock

docker-build:
	@docker build -t cabi99/components-mock:$(VERSION) .

deploy-app:
	@kubectl apply -f ./deploy/demo/app/

deploy-scaledobject:
	@kubectl apply -f ./deploy/demo/scaledobject/metrics-api.yaml
