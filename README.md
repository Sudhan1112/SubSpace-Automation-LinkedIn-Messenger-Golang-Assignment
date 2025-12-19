# SubSpace Automator (LinkedIn Messenger – Golang)

A **Go + Rod–based browser automation proof-of-concept** demonstrating controlled, human-like browser behavior and automation orchestration through a minimal dashboard interface.

> ⚠️ **Educational Proof-of-Concept Only**  
> This project is NOT intended for production usage and MUST NOT be used on real LinkedIn accounts.

---

## ⚠️ Legal & Ethical Disclaimer

This project is created **strictly for educational and technical evaluation purposes**.

- Automating LinkedIn violates LinkedIn’s Terms of Service.
- This repository exists to demonstrate **browser automation mechanics, system design, and UI-to-backend state control**.
- The author assumes **no responsibility for misuse** of this code.

---

## 🎯 Project Objective

This project serves as a technical proof-of-concept to:

- Demonstrate browser automation using **Golang + Rod**
- Simulate **human-like behavior** (delays, cursor movement, pacing)
- Showcase **clean backend architecture**
- Provide **clear automation state visibility** through UI
- Handle failures and shutdowns gracefully

---

## 🧠 System Overview

The system is composed of the following components:

- **Frontend UI** – React + Tailwind dashboard for control & monitoring
- **Go Automation Engine** – Core automation logic
- **Stealth Layer** – Human-like behavior simulation
- **Rod Controller** – Chromium browser control
- **State Storage** – SQLite / JSON persistence
- **Logging Layer** – Real-time structured logs

📘 Detailed architecture: `docs/HLD.md`

---

## 📂 Project Structure

```text
/backend
 ├── cmd
 │   └── server
 │       └── main.go
 ├── internal
 │   ├── api
 │   │   └── handler.go
 │   ├── automation
 │   │   ├── auth.go
 │   │   ├── browser.go
 │   │   ├── connect.go
 │   │   ├── message.go
 │   │   └── search.go
 │   ├── models
 │   │   ├── config.go
 │   │   └── data.go
 │   ├── store
 │   │   └── sqlite.go
 │   └── utils
 │       └── random.go
 ├── go.mod
 └── go.sum

/frontend
 ├── src
 │   ├── assets
 │   ├── App.jsx
 │   ├── App.css
 │   ├── index.css
 │   └── main.jsx
 ├── package.json
 └── vite.config.js

/docs
 ├── HLD.md
 ├── LLD.md
 ├── API.md
 └── STEALTH.md

README.md
````

---

## ⚙️ Tech Stack

* **Language**: Golang (1.20+)
* **Automation**: Rod (Chrome DevTools Protocol)
* **Browser**: Chromium
* **Frontend**: React, Tailwind CSS, Vite
* **Storage**: SQLite
* **Logging**: Structured logging (Zap / custom)

---

## 🚀 How to Run

### Backend

```bash
cd backend
go run cmd/server/main.go
```

Server runs on:

```
http://localhost:8080
```

---

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend runs on:

```
http://localhost:5173
```

---

## 🖥️ Dashboard UI States & Automation Flow

The dashboard is **state-driven**.
Each screenshot below represents a **distinct automation lifecycle stage**.

---

### 🟢 Idle / Ready State

![Idle State](https://github.com/user-attachments/assets/a6c87771-250f-4d00-97d5-f9a3d0628bd9)

**Description**

* Initial state after startup
* System is ready but inactive

**UI Elements**

* Custom Credentials (Optional Input)
* Disabled Search Bar
* Automation Button (▶ Start)
* Console:

  ```
  Ready to start.
  ```

**Backend State**

* Server running
* No browser instance
* No credentials loaded

---

### 🔐 Credentials Expanded State

![Credentials Expanded](https://github.com/user-attachments/assets/87d20856-ee2c-46fa-b582-945906eec77d)

**Description**

* User expands optional credential input

**UI Elements**

* Email input
* Password input (masked)
* Hide Credentials button
* Automation Button (▶ Start)

**Backend State**

* Credentials stored **in-memory only**
* No disk or log persistence
* Supports credential or session-based login

---

### ▶️ Automation Running State

![Automation Running](https://github.com/user-attachments/assets/023fe7b5-e596-45ad-9e09-1c84ea8e7919)

**Description**

* Active automation execution

**UI Elements**

* Enabled Search Bar (e.g. `Software developer`)
* Stop System button (■)
* Live console logs

**Observed Console Output**

```
Ready to start.
Initializing browser...
Logging in...
Login successful.
Searching for: Software developer
Found 0 profiles.
Automation flow finished.
```

**Backend State**

* Chromium launched via Rod
* Stealth configuration applied
* Search & pagination executed
* Graceful shutdown on completion

---

### 🔴 Stop / Interrupt Handling

**Behavior**

* Triggered via Stop System button
* Cancels automation safely
* Closes browser context
* Frees resources
* Returns system to Idle state

Prevents orphaned Chromium processes and partial execution.

---

## 🔄 UI State Summary

| State                | Action Button | Credentials | Search   | Browser     |
| -------------------- | ------------- | ----------- | -------- | ----------- |
| Idle                 | ▶ Start       | Hidden      | Disabled | Not Running |
| Credentials Expanded | ▶ Start       | Visible     | Disabled | Not Running |
| Running              | ■ Stop        | Locked      | Enabled  | Running     |
| Finished             | ▶ Start       | Hidden      | Disabled | Closed      |

---

## 🔄 Automation Flow (High-Level)

1. Load configuration
2. Launch Chromium via Rod
3. Apply stealth & human-like behavior
4. Authenticate (credentials or session)
5. Execute search workflow
6. Process results
7. Persist state/logs
8. Shutdown gracefully

---

## 🎥 Demo Video

📎 Demo Video:
[https://drive.google.com/file/d/1cVEbIblPFTai8m1Yp8zfXk7aJgMFQidG/view](https://drive.google.com/file/d/1cVEbIblPFTai8m1Yp8zfXk7aJgMFQidG/view)

---

## 📚 Documentation

* **HLD**: `docs/HLD.md`
* **LLD**: `docs/LLD.md`
* **API Docs**: `docs/API.md`
* **Stealth Strategy**: `docs/STEALTH.md`

---
