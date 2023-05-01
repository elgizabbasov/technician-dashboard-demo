# CarbinX Technician Maintenance Hub
Welcome to CarbinX's technician maintenance web app!
The CarbinX [Technician Maintenance Hub](https://technician-dashboard-demo.vercel.app/) is a user-friendly web application designed to help CleanO2 technicians to set up and maintain CarbinX devices worldwide. The application also allows technicians to easily monitor device & sensor health via dashboards and send regular maintenance reports.

CleanO2 Carbon Capture Technologies' CarbinX device captures carbon dioxide from heating-appliance flue gas in a one-step, zero-waste reaction. While the existing solution for the maintenance of the devices is sufficient, it needs certain features and is not easy to scale. In particular, it is challenging for technicians to view real-time device updates during bi-weekly maintenance inspections. The CarbinX Technician Dashboard application features an intuitive interface that allows users to overview and examine charts of various components of the CarbinX device, such as temperature sensor data for a device. 

The CarbinX Technician Maintenace Hub is an innovative software solution for the CleanO2 technician staff looking to get informed on the state of health of the device and its components or gather broad device and sensor information without needing physical work or expertise.

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
