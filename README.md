# Auxilify

## Project Description

Auxilify is a desktop application built with Wails, Go, and SvelteKit, designed to provide utilities for image and video processing. It leverages Go for backend logic (e.g., image and video conversion) and SvelteKit for a modern, responsive frontend user interface.

## Setup and Installation

To get this project up and running on your local machine, follow these steps:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/Auxilify.git
    cd Auxilify
    ```

2.  **Install Go dependencies:**
    Ensure you have Go installed (version 1.18 or higher recommended).
    ```bash
    go mod tidy
    ```

3.  **Install Frontend dependencies:**
    Navigate to the `frontend` directory and install Node.js dependencies using npm. Ensure you have Node.js and npm installed.
    ```bash
    cd frontend
    npm install
    cd ..
    ```

4.  **Install Wails CLI:**
    If you don't have the Wails CLI installed, you can install it using:
    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
    For more details on Wails installation, refer to the [Wails documentation](https://wails.io/docs/gettingstarted/installation).

## Running the Application

### Development Mode

To run the application in live development mode with hot-reloading for frontend changes:

```bash
wails dev
```

This will start a Vite development server for the frontend and a Wails dev server. You can access the frontend in your browser at `http://localhost:34115` to interact with your Go methods directly from browser devtools.

### Building for Production

To build a redistributable, production-ready package of the application:

```bash
wails build
```

This command will compile the Go backend and bundle the SvelteKit frontend into a single executable. The output will be located in the `build/bin` directory.
