# CarbinX Device Maintenance Hub
Welcome to CarbinX's technician dashboard web app!

The application was developed over Ubuntu 20.04 using Gin Web Framework with Go 1.19.5, Vue.js 2.6 with Vuetify 2.6, and SQLite3.

NOTE: Login functionality temporarily disabled.

## Requirements
Before setting up and running the application, please make sure you have the following software installed:
- Node.js (https://nodejs.org/en/download/)
- Go (https://golang.org/dl/)

## Development Setup
To get started, clone the repository:

```sh
$ git clone github.com/elgizabbasov/technician-dashboard-demo
```

Next, install the required dependencies (Note: You may need to install `npm` if it wasn't installed with Node.js). For the Vue.js frontend, navigate to the `app` directory and run:

```sh
$ cd app
$ (sudo) npm install
```

For the Go backend, navigate to the root directory and run:

```sh
$ go mod download
```

## Run Development Server
To run the development server for the Go backend, navigate to the root directory and run:

```sh
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
