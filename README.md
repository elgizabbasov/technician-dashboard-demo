# CarbinX Device Maintenance Hub
Welcome to CarbinX's technician dashboard web app!

The application was developed over Ubuntu 20.04 using Gin Web Framework with Go 1.19.5, Vue.js 2.6 with Vuetify 2.6, and SQLite3.

## Requirements
Before setting up and running the application, please make sure you have the following software installed:
- Node.js (https://nodejs.org/en/download/)
- Go (https://golang.org/dl/)

## Development Setup
To get started, clone the repository:

```sh
$ git clone https://github.com/CleanO2tech/clean-technician-web.git
```

Next, install the required dependencies. For the Vue.js frontend, navigate to the `app` directory and run:

```sh
$ cd app
$ (sudo) npm install
```

For the Go backend, navigate to the `api` directory and run:

```sh
$ cd api
$ go mod download
```

## Run Development Server
To run the development server for the Go backend, navigate to the `api` directory and run:

```sh
$ cd api
$ go run main.go
```

To run the development server for the Vue.js frontend, navigate to the `app` directory and run:

```sh
$ cd app
$ npm run serve
```

Then, navigate to the host address presented in the terminal (often `http://127.0.0.1:8000/`).

### List of Contributors
- Elgiz Abbasov @elgizabbasov
- Zane Regel @zaneregel
- Jaron Baldazo @bjaron
- Manjot Singh @ManjotSin
- Duncan Grace @duncangrace