# school

## Project Setup

To initially start the school, execute the following steps

- Clone the project from GitLab
- Download and install Go (1.25.5) <https://go.dev/doc/install>
- Verify the golang version with ``go version``
- Install Project Dependencies with ``go get ./...`` or ``go mod tidy``


## Run

To Start the Backend Server follow the steps below

- Add the ``.env`` File from MOFIS Wiki into the ``.\cmd\`` Directory
- From `.\cmd\` Directory execute ``go run main.go``

## Build

The following commands can be used to build the application
 
   - ``go build -v -o cmd/school.exe cmd/main.go``

 go build with version example:
    
   - `GOOS=windows GOARCH=amd64 go build -v -ldflags="-s -w -X 'school/internal/build.gitVersion=v.0.0.1' -X 'school/internal/build.gitCommit=dev' -X 'school/internal/build.date=2024-01-01T00:00:00Z'" -o cmd/school.exe cmd/main.go`

For Production the gitversion must be set 

   - Default: `gitVersion=development` to run the application in development mode

## Test

1. Install [TDM GCC Compiler (tdm64-gcc-10.3.0-2)](https://jmeubank.github.io/tdm-gcc/articles/2021-05/10.3.0-release) and add it to your system's environment variables.
2. Enable the use of the CGo tool for SQLite tests with: ``   go env -w CGO_ENABLED=1``
3. Tests can be executed using the commands ``go test ./... -v`` and ``go clean -testcache``
4. Note the use of the CGo tool on the deployment server must be disabled: ``go env -w CGO_ENABLED=0``

## API Documentation

- Make sure swagger is installed: `go install github.com/swaggo/swag/cmd/swag@latest`
- Add User Environment Variable: `Path=C:\Users\{user}\go\bin` and restart vscode

To create the Api Specification with Swagger execute following commands

   - generate api docs: `swag init -g .\cmd\main.go`  

The API Specification can be tried out at the following URL

   - ` http://localhost:8083/api/swagger/index.html`

## Development

As Development Environment we recommend VSCode with the ``Id: golang.go`` Extension

Debug Config

```
    - .vscode\launch.json config

    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch Package",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "program": "${workspaceFolder}/cmd"
            }
        ]
    }
```

