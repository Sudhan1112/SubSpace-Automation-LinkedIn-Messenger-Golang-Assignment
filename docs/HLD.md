# High-Level Design (HLD)

## 1. Purpose
The **SubSpace Automator** is a modular browser automation system designed to simulate human interactions on professional networking platforms. Its primary purpose is to demonstrate how to perform robust, stealthy automation using Go and Rod, while decoupling the control logic from the browser execution details.

## 2. Architecture Diagram

```mermaid
graph TD
    User[User] -->|Interacts| Frontend[Frontend UI (React)]
    Frontend -->|HTTP REST| API[Go API Server]
    
    subgraph "Go Automation Ecosystem"
        API -->|Controls| Engine[Automation Engine]
        
        Engine -->|Uses| Stealth[Stealth Layer]
        Engine -->|Persists| Storage[(SQLite DB)]
        Engine -->|Logs| Logger[Logger Service]
        
        Stealth -->|Wraps| Rod[Rod Controller]
    end
    
    Rod -->|CDP Protocol| Browser[Chromium Browser]
    Browser -->|Renders| Target[Target Website]
```

## 3. Component Responsibilities

### Frontend UI
- Provides a user-friendly interface for configuration and monitoring.
- Displays real-time logs and scraped data.
- Sends start/stop commands to the backend.

### Go API Server
- Exposes REST endpoints (`/start`, `/stop`, `/status`).
- Bridges the communication between the UI and the Automation Engine.

### Automation Engine
- The "Brain" of the system.
- Manages the lifecycle of the automation task.
- Orchestrates the sequence of actions (Login -> Search -> Scrape).
- Handles error recovery and flow control.

### Stealth Layer
- Intercepts browser commands to inject human-like behaviors.
- Modifies `navigator` properties to mask WebDriver signals.
- Implements randomized delays and non-linear mouse movements.

### Rod Controller
- Wraps the Rod library to provide a high-level abstraction for browser actions.
- Manages pages, contexts, and elements.

### Storage & Logger
- **Storage**: Persists profile data and operational state (SQLite).
- **Logger**: Captures detailed execution logs for debugging and UI streaming.

## 4. Data Flow
1.  **Input**: User sends a command (e.g., "Search for Engineer") via the UI.
2.  **Processing**: The Engine accepts the command, initializes a new browser context.
3.  **Execution**:
    - The Engine instructs the Rod Controller to navigate.
    - The Stealth Layer ensures the navigation looks organic.
    - Data is extracted from the DOM.
4.  **Persistence**: Extracted data is sent to the Storage layer.
5.  **Feedback**: Operational logs are streamed back to the UI in real-time.
6.  **Output**: Final results are displayed in the Dashboard.
