# Low Level Design (LLD)

## Backend Structure (`/backend`)

### Packages
*   **`cmd/server`**: Application entry point. Loads config, initializes Store, BrowserManager, and API Handler, then starts the HTTP server.
*   **`internal/automation`**:
    *   `BrowserManager`: Struct holding `*rod.Browser` and `*rod.Page`. Methods: `Start()`, `Stop()`, `Login()`, `SearchProfiles()`, `SendConnectionRequest()`, `SendMessage()`.
    *   `Stealth`: Implements `stealth` library usage and `HumanMove` (random mouse simulation).
*   **`internal/api`**:
    *   `Handler`: Struct holding references to `BrowserManager` and `Store`.
    *   `ServeHTTP`: Handles routing.
    *   `Endpoints`:
        *   `StartAutomation`: Triggers the main automation goroutine.
        *   `StopAutomation`: Signals stop.
        *   `GetStatus`: Returns current logs (in-memory buffer) and running state.
        *   `GetData`: Returns records from SQLite.
*   **`internal/store`**:
    *   `Store`: Wraps `gorm.DB`.
    *   `SaveActivity`: Inserts `ProfileActivity`.
    *   `GetActivities`: Selects all activities.
*   **`internal/models`**:
    *   `Config`: Env vars.
    *   `Profile`: Struct for scraped profile data.
    *   `ProfileActivity`: Database model.

## Frontend Structure (`/frontend`)

### Components
*   **`App.jsx`**: Main Dashboard component.
    *   **State**: `status` (running, logs), `activities` (history).
    *   **Effects**: Polls `/api/status` every 2s.
    *   **Render**: Header (Start/Stop), Logs Panel (scrollable text), Activity Table (rows of data), Stats Panel (counts).

## API Contract
| Endpoint | Method | Params | Response | Description |
| :--- | :--- | :--- | :--- | :--- |
| `/api/start` | POST | None | `{message: string}` | Starts automation task. |
| `/api/stop` | POST | None | `{message: string}` | Stops automation. |
| `/api/status`| GET | None | `{running: bool, logs: []string}` | Returns live status. |
| `/api/data` | GET | None | `[{id, profile_url, action, timestamp}]` | Returns history. |

## Database Schema (SQLite)
Table `profile_activities`:
*   `id` (PK, Auto Inc)
*   `profile_url` (Text)
*   `action` (Text) - e.g., "SEARCH_FOUND", "CONNECT"
*   `metadata` (Text) - e.g., Name
*   `timestamp` (DateTime)
