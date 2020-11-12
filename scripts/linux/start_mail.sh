export ENV_RABBITMQ_HOST=amqp://localhost/
export ENV_SENDGRID_API_KEY=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
export ENV_SENDER_EMAIL=info@paingha.tech
export ENV_BASE_URL=http://localhost:8080/v1
go mod vendor
go build -o ./mailservice/main.exe  ./mailservice
cd mailservice
main.exe

pause