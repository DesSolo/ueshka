PROJECT_NAME=$(shell basename "$(PWD)")
BINARIES_DIRECTORY="bin"

clean:
		@rm -rf ${BINARIES_DIRECTORY}

build: clean
		@go build -ldflags="-w -s" -o ${BINARIES_DIRECTORY}/${PROJECT_NAME}

run: build
		@${BINARIES_DIRECTORY}/${PROJECT_NAME}

release: clean
		@echo "> build"

		@echo "mipsle"
		@GOOS=linux GOARCH=mipsle CGO_ENABLED=0 go build -ldflags="-w -s" -o ${BINARIES_DIRECTORY}/${PROJECT_NAME}-mipsle
		
		@echo "...done"