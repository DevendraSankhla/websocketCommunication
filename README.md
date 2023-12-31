GoChat Web Application

This is a simple chat application written in Go (Golang) that allows multiple clients to connect to a single WebSocket. The application broadcasts messages from one client to all connected clients and also displays the chat history on the front end.

-Features:
  1. WebSocket Communication: Utilizes the Gorilla WebSocket library for handling WebSocket connections.
  2. Broadcasting: Messages sent by one client are broadcasted to all connected clients.
  3. Chat History: Maintains a chat history that is displayed on the front end.

-Usage:
  1. Prerequisites
    a. Go (Golang) installed on your machine.

-Installation:
  1. Clone the repository: (Run following cmd)
    a. git clone https://github.com/your/repo.git
    b .cd repo

  2. Run the application: (Run following cmd)
    a. go run main.go

  3. Open your web browser and go to http://localhost:8080 to access the chat application.

-How to Use:
  1. Open the chat application in multiple browser tabs or devices.
  2. Enter a message in the input field and click the "Send" button.
  3. The message will be broadcasted to all connected clients, and you will see it in the chat history.

-Code Structure:
  1. main.go: Contains the main application logic, WebSocket handling, and HTTP server setup.
  2. static/index.html: HTML file for the chat application front end.

-Front End:
  The front end is a simple HTML file (static/index.html) with JavaScript to handle WebSocket communication. The chat history is displayed, and users can input messages to send to the server.

-Dependencies:
  Gorilla WebSocket: A WebSocket implementation for Go.

-Contributing:
  Feel free to contribute to the project by opening issues or submitting pull requests.

-License:
  This project is licensed under the MIT License - see the LICENSE file for details.
