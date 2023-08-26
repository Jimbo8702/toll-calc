obu: 
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver: 
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator: 
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

kafkaup:
	@docker-compose up -d
	
kafkadown:
	@docker-compose down 

.PHONY: obu