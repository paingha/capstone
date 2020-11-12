export ENV_RABBITMQ_HOST=amqp://localhost/
export ENV_TWILIO_ACCOUNT_SID=AC31c081baf99da0cfe90b22efa35092a5
export ENV_TWILIO_AUTH_TOKEN=18763b42ef34f119d4729235dac5ce6c
export ENV_SENDER_PHONE=+14155238886
go mod vendor
go build -o ./smsservice/main.exe  ./smsservice
cd smsservice
main.exe

pause