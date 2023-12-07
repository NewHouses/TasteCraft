## build: builds all binaries
build: clean build_back
	@echo All binaries built!

## clean: cleans all binaries and runs go clean
clean:
	@echo Cleaning...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned and deleted binaries

## build_back: builds the back end
build_back:
	@echo Building back end...
	@go build -o dist/go_tastecraft_api.exe .\cmd\api
	@echo Back end built!

## start: starts front and back end
start: start_back

## start_back: starts the back end
start_back: build_back
	@echo Starting the back end...
	set && start /B .\dist\go_tastecraft_api.exe
	@echo Back end running!

## stop: stops the front and back end
stop: stop_back
	@echo All applications stopped

## stop_back: stops the back end
stop_back:
	@echo Stopping the back end...
	@taskkill /IM go_tastecraft_api.exe /F
	@echo Stopped back end