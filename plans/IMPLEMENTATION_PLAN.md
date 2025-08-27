# Implementation Plan: `cloud-status`

This document outlines the plan to create the `cloud-status` CLI and MCP server.

## Lessons Learned from `calctl` and `drivectl`

This project will adopt several successful patterns from the `calctl` and `drivectl` projects:

*   **Project Structure:** The `cmd/`, `internal/`, `mcp/`, and `plans/` directory structure will be used for consistency.
*   **Configuration:** `github.com/spf13/viper` will be used for managing configuration, such as the MCP server address.
*   **Testing:** A dedicated `TEST_PLAN.md` will be created to ensure robust testing of both CLI and MCP functionality.
*   **MCP Server:** The explicit tool registration pattern from `calctl` will be used for clarity and simplicity.
*   **Simplification:** Unlike the other tools, `cloud-status` accesses public URLs, so it does **not** require the complex OAuth2 authentication logic.

---

## Phase 1: Project Scaffolding & Setup

1.  **Create Project Structure:** Ensure the following directory structure exists:
    *   `cmd/`, `internal/fetcher`, `internal/parser`, `internal/types`, `mcp/`, `plans/`
2.  **Dependencies:** Add the necessary libraries to `go.mod`:
    *   `github.com/spf13/cobra`
    *   `github.com/spf13/viper`
    *   `github.com/gemini-cli/mcp/go-sdk` (with a `replace` directive to the local SDK path)
    *   A suitable XML parsing library (e.g., `github.com/antchfx/xmlquery` or similar).

## Phase 2: Core Logic Implementation (`internal/`)

1.  **Define Data Structures (`internal/types/`):** Create Go `structs` to unmarshal the data from both the Atom (`feed.atom`) and JSON (`incidents.json`) feeds.
2.  **Implement Data Fetcher (`internal/fetcher/`):** Write functions to download the content from the two Google Cloud status URLs using standard HTTP calls.
3.  **Implement Parsers (`internal/parser/`):**
    *   **Atom Parser:** Create a function to parse the Atom XML feed into the defined Go structs.
    *   **JSON Parser:** Create a function to parse the `incidents.json` array into a slice of incident structs.

## Phase 3: CLI Implementation (`cmd/`)

1.  **`root` Command:** Set up the root command with persistent flags managed by `viper` (e.g., for the MCP server address).
2.  **`current` Command:** This command will use the fetcher and Atom parser to display the most recent status update in a clean, readable format.
3.  **`history` Command:** This command will use the fetcher and JSON parser to display a list of historical incidents. It will include flags for filtering by service, severity, and limiting the number of results (e.g., `cloud-status history --service "Google Compute Engine" --limit 10`).

## Phase 4: MCP Server Implementation (`mcp/`)

1.  **Define Tool:** Define an MCP tool named `get_cloud_status` with parameters:
    *   `source`: (enum: "current", "history")
    *   `service_filter`: (string, optional)
    *   `limit`: (int, optional)
2.  **Implement MCP Server:** In `mcp/server.go`, implement the server using the explicit registration pattern. The tool handler will reuse the core logic from the `internal/` packages.

## Phase 5: Documentation

1.  **`README.md`:** Create a comprehensive `README.md` with a project overview, build/installation instructions, and usage examples for all CLI commands.
2.  **`plans/TEST_PLAN.md`:** Create the initial test plan document.

## Phase 6: Testing

1.  **Flesh out `TEST_PLAN.md`:** Detail specific manual test cases for:
    *   `cloud-status current`
    *   `cloud-status history` (with and without filters)
    *   Starting the MCP server.
    *   Calling the `get_cloud_status` tool via an MCP client with various parameters.
2.  **Execute Test Plan:** Manually execute all tests outlined in the plan.
3.  **Commit Changes:** After all tests pass successfully, commit the code changes to the Git repository with a descriptive message.