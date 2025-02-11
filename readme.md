# 📌 PinKing

A simple application where users can upload images and place pin markings on them. The project consists of a Go server using PocketBase for backend functionalities and an Expo-based frontend for a smooth cross-platform experience.

## 🚀 Features
- Upload pictures
- Add pin markers to images
- View and manage pinned locations
- Backend powered by PocketBase and Go
- Cross-platform frontend built with Expo

## 🛠 Setup & Installation

### Prerequisites
Make sure you have the following installed:
- [Go](https://go.dev/dl/)
- [PocketBase](https://pocketbase.io/docs/)
- [Node.js](https://nodejs.org/)
- [Expo CLI](https://docs.expo.dev/get-started/installation/)

### Backend Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/felixnhs/pinking-go.git
   cd pinking-go
   ```
2. Install dependencies and start PocketBase:
   ```sh
   cd backend
   go mod tidy
   ./pocketbase serve
   ```
3. Run the Go server:
   ```sh
   go run main.go
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```sh
   cd ui
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the Expo development server:
   ```sh
   expo start
   ```

## 📜 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎓 Learning Purpose

This project was created solely for learning and experimenting with Go, PocketBase, and Expo. It does not serve any higher purpose or commercial goal.