# SubSpace Automator
A Go + Rod based browser automation proof-of-concept demonstrating human-like behavior and stealth techniques for educational evaluation.

## ⚠️ Legal & Ethical Disclaimer
**This project is created strictly for educational and technical evaluation purposes.**
It is **NOT** intended for production use and must **NOT** be used on real LinkedIn accounts.
Automating LinkedIn violates their Terms of Service. The authors are not responsible for any misuse of this software.

## 🎯 Project Objective
This project serves as a technical proof-of-concept to:
- **Demonstrate browser automation** using the Go programming language and the Rod library.
- **Simulate human-like interactions** to understand the mechanics of bot detection evasion.
- **Showcase clean system design** with a modular Go architecture and separation of concerns.
- **Handle automation failures gracefully**, ensuring robustness in uncertain web environments.

## 🧠 System Overview
The system is composed of the following key components:
- **Frontend UI**: A React + Tailwind dashboard for control and monitoring (Optional addition).
- **Go Automation Engine**: The core logic driver handling task scheduling and execution.
- **Stealth Layer**: A dedicated module ensuring interactions appear organic (mouse movements, delays).
- **Rod Controller**: Manages the headless (or headful) Chromium browser instance.
- **State Storage**: SQLite/JSON persistence for resume capability and data logging.
- **Logging**: Structured logging for real-time feedback and debugging.

Detailed architecture is available in [docs/HLD.md](docs/HLD.md).

## ⚙️ Tech Stack
- **Language**: Golang (1.20+)
- **Automation Lib**: Rod (DevTools Protocol)
- **Browser**: Chromium
- **Frontend**: React, Tailwind CSS (Vite)
- **Storage**: SQLite
- **Logging**: Zap / Custom Structured Logger

## 🚀 How to Run

### 1. Backend Setup
```bash
cd backend
# Create a .env file based on example if needed, or defaults will be used
# cp .env.example .env 
go run cmd/server/main.go
```
Server runs on `http://localhost:8080`

### 2. Frontend Setup
```bash
cd frontend
npm install
npm run dev
```
Frontend runs on `http://localhost:5173`

Open your browser to `http://localhost:5173` to interact with the Automator.

## 🔄 Automation Flow
1.  **Initialization**: Configuration is loaded from `.env` (if present) or defaults.
2.  **Launch**: Chromium is launched via Rod with specific flags to mask automation signals.
3.  **Stealth**: The stealth layer initializes, overriding default navigator properties.
4.  **Action**:
    - **Login**: Attempts credential-based login (if provided) or waits for session.
    - **Search**: Navigates to search results based on user query.
5.  **Behavior**: All actions are rate-limited with randomized "human" delays and Bézier curve mouse movements.
6.  **Persistence**: Scraped data and session state are persisted to SQLite.
7.  **Exit**: The system handles stops or failures by cleaning up browser contexts gracefully.

## 🎥 Demo Video
[Link to Demo Video would go here]

## 📚 Additional Documentation
- **High-Level Design**: [docs/HLD.md](docs/HLD.md)
- **Low-Level Design**: [docs/LLD.md](docs/LLD.md)
- **API Documentation**: [docs/API.md](docs/API.md)
- **Stealth Strategy**: [docs/STEALTH.md](docs/STEALTH.md)
