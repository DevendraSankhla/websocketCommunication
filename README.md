# GoChat Web Application

This is a simple chat application written in Go (Golang) that allows multiple clients to connect to a single WebSocket. The application broadcasts messages from one client to all connected clients and also displays the chat history on the front end.

## Features

- WebSocket Communication: Utilizes the Gorilla WebSocket library for handling WebSocket connections.
- Broadcasting: Messages sent by one client are broadcasted to all connected clients.
- Chat History: Maintains a chat history that is displayed on the front end.

## Usage

### Prerequisites
- Go (Golang) installed on your machine.

## Installation

1. Clone the repository.
2. Run the application.
3. Open your web browser and go to http://localhost:8080 to access the chat application.

## How to Use

- Open the chat application in multiple browser tabs or devices.
- Enter a message in the input field and click the "Send" button.
- The message will be broadcasted to all connected clients, and you will see it in the chat history.

## Code Structure

- `main.go`: Contains the main application logic, WebSocket handling, and HTTP server setup.
- `static/index.html`: HTML file for the chat application front end.

## Front End

The front end is a simple HTML file (`static/index.html`) with JavaScript to handle WebSocket communication. The chat history is displayed, and users can input messages to send to the server.

## Dependencies

- Gorilla WebSocket: A WebSocket implementation for Go.

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
