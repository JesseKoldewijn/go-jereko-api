<h1 align="center">Welcome to Go Jereko Api 👋</h1>
<p>
  <a href="https://github.com/JesseKoldewijn/go-jereko-api/actions/workflows/ci.yml"><img src="https://github.com/JesseKoldewijn/go-jereko-api/actions/workflows/ci.yml/badge.svg" alt="Go"></a>
  <a href="https://github.com/JesseKoldewijn/go-jereko-api/blob/main/LICENCE" target="_blank"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" /></a>
</p>

> A simple GoLang service using Gin to serve rest endpoints in a efficient way.

## Install

You can either use the docker image or you local environment to run the service.

### Local

If you're familiar with GoLang's Air cli you can just run mod download and run the service using Air.

```sh
go mod download && go run main.go
```

### Docker

You can also run the service using docker.

```sh
docker build -t go-jereko-api .
docker run -p 8080:8080 go-jereko-api
```

## Author

👤 **Jesse Koldewijn**

-   Website: https://jereko.dev
-   Github: [@JesseKoldewijn](https://github.com/JesseKoldewijn)

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2024 [Jesse Koldewijn](https://github.com/JesseKoldewijn).<br />
This project is [MIT](https://github.com/JesseKoldewijn/go-jereko-api/blob/main/LICENCE) licensed.
