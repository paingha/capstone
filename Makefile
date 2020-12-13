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

