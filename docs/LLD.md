# Low-Level Design (LLD)

## 1. Folder / Package Structure

```
/backend
├── cmd
│   └── server          # Entry point for the API server
├── internal
│   ├── api             # HTTP Handlers and Router configuration
│   ├── automation      # Core logic for task execution (Engine)
│   ├── browser         # Rod wrapper and browser instance management
│   ├── models          # Data structures (Profile, Task, Status)
│   ├── store           # SQLite database implementation
│   └── utils           # Helper functions (Logger, Randomizers)
```

## 2. Module Responsibilities

### `internal/api`
- **Does**: Parse JSON requests, validate input, call Automation methods, return JSON responses.
- **Does NOT**: Contain business logic or browser control code.

### `internal/automation`
- **Does**: Define the steps of the automation (e.g., `RunSearch`). Manage the `State` of the current job.
- **Does NOT**: Directly interact with the implementation details of the database or HTTP server.

### `internal/browser`
- **Does**: Initialize `rod.Browser`, apply `stealth` plugins, handle page navigation, input typing, and element clicking.
- **Does NOT**: Decide *what* to scrape, only *how* to interact.

### `internal/store`
- **Does**: Handle SQL queries, open/close DB connections.
- **Does NOT**: Process business logic.

## 3. Key Structs / Concepts

- **`AutomationService`**: The singleton struct holding the current state, running status, and logs.
- **`Status`**: A struct containing `Running` (bool) and `Logs` ([]string) shared with the frontend.
- **`Profile`**: The domain model representing a scraped candidate.

## 4. Error Handling Strategy
- **No Panics**: The system strictly avoids `panic` in favor of returning `error`.
- **Propagation**: Errors are wrapped with context (e.g., `fmt.Errorf("failed to login: %w", err)`) and bubbled up to the Engine.
- **Graceful Failure**: If a step fails (e.g., selector not found), the automation logs the error, attempts a retry (if configured), or halts the specific task without crashing the server.
- **Cleanup**: `defer` statements ensure the browser and pages are closed even on error.

## 5. State Handling
- **InMemory State**: The `AutomationService` keeps the immediate `running` status and a buffer of `logs` for the UI.
- **Persistent State**: The `SQLite` database stores successful `profiles` to ensure data isn't lost if the app restarts.
- **Resume Capability**: (Future Scope) The DB design supports storing "Last Scraped Page" to resume interrupted tasks.
