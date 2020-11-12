.PHONY: start-api
start-api:
	if [ -a ./api/api.exe ]; then rm -rf ./api/api.exe; fi;
	@echo "Starting API Service"
ifeq ($(OS),Windows_NT)
	cd ./scripts/win && ./start.cmd
else
	cd ./scripts/linux && ./start.sh
endif

.PHONY: start-export
start-export:
	@echo "Starting Export Service"
	cd ./exportservice && cargo run main.rs

.PHONY: start-mail
start-mail:
	if [ -a ./mailservice/main.exe ]; then rm -rf ./mailservice/main.exe; fi;
	@echo "Starting Mail Queue Service"
ifeq ($(OS),Windows_NT)
	cd ./scripts/win && ./start_mail.cmd
else
	cd ./scripts/linux && ./start_mail.sh
endif

.PHONY: start-sms
start-sms:
	if [ -a ./smsservice/main.exe ]; then rm -rf ./smsservice/main.exe; fi;
	@echo "Starting Sms Queue Service"
ifeq ($(OS),Windows_NT)
	cd ./scripts/win && ./start_sms.cmd
else
	cd ./scripts/linux && ./start_sms.sh
endif

.PHONY: start-push
start-push:
	if [ -a ./pushservice/main.exe ]; then rm -rf ./pushservice/main.exe; fi;
	@echo "Starting Push Queue Service"
ifeq ($(OS),Windows_NT)
	cd ./scripts/win && ./start_push.cmd
else
	cd ./scripts/linux && ./start_push.sh
endif
	
.PHONY: start-webapp
start-webapp:
	@echo "Starting Web App"
	cd webapp && yarn start

.PHONY: build-webapp
build-webapp:
	@echo "Building Web App"
	cd webapp && yarn build

.PHONY: clear-api
clear-api:
	@echo "Clear API exe"
	rm -rf ./api/main.exe

.PHONY: clear-mail
clear-mail:
	@echo "Clear Mail Service exe"
	rm -rf ./mailservice/main.exe

.PHONY: clear-sms
clear-sms:
	@echo "Clear SMS Service exe"
	rm -rf ./smsservice/main.exe

.PHONY: clear-push
clear-push:
	@echo "Clear Push Service exe"
	rm -rf ./pushservice/main.exe

.PHONY: upload-api
upload-api:
	@echo "Uploading api"
	cd ./terraform/api && terraform init && terraform plan && terraform apply

.PHONY: upload-mail
upload-mail:
	@echo "Uploading Mail Service"
	cd ./terraform/mailservice && terraform init && terraform plan && terraform apply

.PHONY: upload-sms
upload-sms:
	@echo "Uploading SMS Service"
	cd ./terraform/api && terraform init && terraform plan && terraform apply

.PHONY: upload-push
upload-push:
	@echo "Uploading Push Service"
	cd ./terraform/api && terraform init && terraform plan && terraform apply

.PHONY: upload-upload
upload-upload:
	@echo "Uploading Upload Service"
	cd ./terraform/api && terraform init && terraform plan && terraform apply

.PHONY: plan-api
plan-api:
	@echo "Terraform Planing API"
	cd ./terraform/api && terraform init && terraform plan

.PHONY: plan-mail
plan-mail:
	@echo "Terraform Planning Mail Service"
	cd ./terraform/mailservice && terraform init && terraform plan

.PHONY: plan-sms
plan-sms:
	@echo "Terraform Planning SMS Service"
	cd ./terraform/smsservice && terraform init && terraform plan

.PHONY: plan-push
plan-push:
	@echo "Terraform Planning Push Service"
	cd ./terraform/pushservice && terraform init && terraform plan

.PHONY: plan-upload
plan-Upload:
	@echo "Terraform Planning Upload Service"
	cd ./terraform/uploadservice && terraform init && terraform plan
