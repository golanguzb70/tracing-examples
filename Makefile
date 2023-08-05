jaeger-up: jaeger-down
	cd ./jaeger && docker-compose up -d

jaeger-down: 
	cd ./jaeger && docker-compose down
