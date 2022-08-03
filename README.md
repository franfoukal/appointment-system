# Appointment System

## Initialize project

* Configure environment variable `SCOPE=<SCOPE>` outside `.env` file
  * Example in sh
  
    ```sh
    SCOPE=local \
    go run ./cmd/api/main.go
    ```
* Copy `local.env.example` file and change name for `<SCOPE>.env` and add environment variables
  * Scopes: `TEST`, `LOCAL` or `PRODUCTION`

* Create DB schema `appointments` before start the project


## Considerations

* When adding a new environment variable to `.env` file, add the same to the `.env.example` file with a dummy value.

## VSCode Debug

* add into file `./.vscode/launch.sh` :

    ```json
    {
        // Use IntelliSense to learn about possible attributes.
        // Hover to view descriptions of existing attributes.
        // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch Package",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "program": "${workspaceFolder}/cmd/api/main.go",
                "env": {
                    "ENVIRONMENT": "local"
                }
            }
        ]
    }
    ```