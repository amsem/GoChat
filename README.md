# GoChat

GoChat is a real-time chat application that allows users to communicate instantly. This README provides instructions on how to run the application locally.

## Installation

### Clone the Repository

```bash
git clone https://github.com/amsem/GoChat
```

### Backend Setup

Navigate to the project directory

```bash
cd GoChat/backend
```

Install dependencies

```bash
go mod tidy
```

Start the server

```bash
go run .
```

### Frontend Setup

Navigate to the project directory

```bash
cd GoChat/frontend
```

Install dependencies

```bash
npm install
```

Start the server

```bash
npm start
```

## Docker Deployment

The application also includes a Dockerfile for running both the backend and frontend apps separately. Please refer to the Docker documentation for deployment instructions.

## About

This real-time chat application is built with Go for the backend and React for the frontend. It provides a seamless communication experience for users, allowing them to connect and chat in real-time.

Feel free to contribute to the project by submitting bug reports or feature requests!
