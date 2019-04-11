Goker (Dylan Avery)
=========

Summary
---
This application provides a user with the ability to select a hand of 5 playing cards and get its rank in a game of poker. The server is written in Golang. 

### Server

 * For development
    ```
    #!sh
    go build cmd/main.go
    ./main
    ```

 * For tests
    ```
    #!sh
    cd pkg/poker
    go test
    ```


### Frontend Setup and Run
 * Install dependencies
    ```
    #!sh
    cd web/goker-frontend && yarn
    ```

* Run tests
    ```
    #!sh
    cd web/goker-frontend && yarn test
    ```

 * Run application
    ```
    #!sh
    cd web/goker-frontend && yarn start
    ```


