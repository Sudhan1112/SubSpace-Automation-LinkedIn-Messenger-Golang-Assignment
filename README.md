# Browser Automation Project

A Proof-of-Concept for human-like browser automation using Go (Rod) and React.

## Features
- **Stealth Automation**: Uses Rod + Stealth to simulate human behavior (mouse movements, delays).
- **Backend**: Go-based controller with REST API and SQLite persistence.
- **Frontend**: Modern React + Tailwind CSS dashboard.
- **Functionality**:
    - Automated Login (Credential-based).
    - Profile Search & Pagination.
    - Connection Requests & Messaging.
    - Real-time logging and activity tracking.

## Prerequisites
- Go 1.20+
- Node.js 16+
- Google Chrome / Chromium installed.

## Setup

1.  **Backend Setup**:
    ```bash
    cd backend
    go mod tidy
    cp .env.example .env
    # Edit .env with your LinkedIn credentials (Use a dummy account!)
    ```

2.  **Frontend Setup**:
    ```bash
    cd frontend
    npm install
    ```

## Usage

1.  **Start Backend**:
    ```bash
    cd backend
    go run cmd/server/main.go
    ```
    Server runs on `http://localhost:8080`.

2.  **Start Frontend**:
    ```bash
    cd frontend
    npm run dev
    ```
    Frontend runs on `http://localhost:5173`.

3.  **Run Automation**:
    - Open `http://localhost:5173`.
    - Click **Start**.
    - Observe the browser window (if HEADLESS=false) and the logs on the dashboard.

## Disclaimer
This software is for **educational purposes only**. Automating interactions on websites like LinkedIn may violate their Terms of Service. The authors are not responsible for any account bans or restrictions. Use at your own risk.

## Architecture
See [docs/HLD.md](docs/HLD.md) and [docs/LLD.md](docs/LLD.md) for detailed design.
