# Project Entry for NSU IRB

![Twitter Follow](https://img.shields.io/twitter/follow/painghajnr?style=social)
![GitHub repo size](https://img.shields.io/github/repo-size/paingha/auth-service?style=plastic)
![Coveralls github](https://img.shields.io/coveralls/github/paingha/auth-service)

This is a project entry for the overhaul of the NSU IRB which currently handles communication and applications for research approval through email. This is also my potential capstone project. (Still trying to decide lol 😂).

![](https://raw.githubusercontent.com/paingha/capstone/master/capstone-github-image.PNG?token=AB6SB22F5FXRNOWSR73EUGK72KRP2)

## Installation

Initialize environment variables for API by creating a api.cmd or api.bash file like so:

```bash
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=postgres
set DB_PASS=123456
set DB_DATABASE=databasename
set DB_SSL=disable
set ENV_RABBITMQ_HOST=amqp://localhost/
set ENV_SENDGRID_API_KEY=SendgridApiKeyGoesHere
set ENV_SENDER_EMAIL=mail@senderemail.com
set ENV_BASE_URL=http://localhost:8080/v1
go mod vendor
go build -o ./api/main.exe ./api
cd api
main.exe

pause
```

Initialize environment variables for Mailer Service by creating a mail.cmd or mail.bash file like so:

```bash
set ENV_RABBITMQ_HOST=amqp://localhost/
set ENV_SENDGRID_API_KEY=SendgridApiKeyGoesHere
set ENV_SENDER_EMAIL=mail@senderemail.com
set ENV_BASE_URL=http://localhost:8080/v1
go mod vendor
go build -o ./mailservice/main.exe  ./mailservice
cd mailservice
main.exe

pause
```

## Usage

### Steps:
#### Step 1: Click the mail.cmd file to start the mailer queue service.(Make sure your [RabbitMQ server](https://www.rabbitmq.com) is running)
#### Step 2: Click the api.cmd file to start the API service. (The mail queue service has to be running for the API to work)
#### Swagger UI is available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
#### API is available at [http://localhost:8080/v1](http://localhost:8080/v1)


## Contributing
This project is currently not accepting contributions.

## TODO
The following are todo items for this project
##### - Dockerize API, Mail Queue Service, and SMS Queue Service. - DONE
##### - Setup React Webapp using create-react-app.
##### - Use email verification html files for verification emails.
##### - Add more tests
##### - Setup CI/CD pipeline with Circle CI

## License
[MIT License](https://choosealicense.com/licenses/mit/)

## Authors
##### - Paingha Joe Alagoa - [🌍 Website](http://paingha.me), [🐦 Twitter](https://twitter.com/painghajnr), [💼 Github](https://github.com/paingha)
