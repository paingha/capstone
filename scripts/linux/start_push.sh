export ENV_RABBITMQ_HOST=amqp://localhost/
export ENV_ONE_SIGNAL_APP_KEY=YWRhYTdjNDEtOTNlYS00YWM2LTgzMDktOWZjZjAzMDYzMzll
export ENV_ONE_SIGNAL_APP_ID=2efaa337-277b-4654-810f-83aa8f702717
go mod vendor
del ./pushservice/main.exe
go build -o ./pushservice/main.exe  ./pushservice
cd pushservice
main.exe

pause