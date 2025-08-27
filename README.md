# cloud-status

A command-line tool and MCP server to check the status of Google Cloud services.

This tool fetches data from the public Google Cloud status pages:
- **Atom Feed:** `https://status.cloud.google.com/en/feed.atom`
- **JSON History:** `https://status.cloud.google.com/incidents.json`

## Installation

To install the `cloud-status` CLI, you can use `go install`:

```sh
go install github.com/ghchinoy/cloud-status@latest
```

## CLI Usage

The CLI provides commands to get the current status and historical incident data.

### Get Current Status

To see the most recent updates from the Google Cloud status feed:

```sh
cloud-status current
```

### Get Historical Incidents

To get a list of historical incidents:

```sh
cloud-status history
```

You can filter the history with the following flags:

- `--limit`: Limit the number of incidents returned (e.g., `--limit 5`).
- `--service`: Filter by a specific service name (e.g., `--service "Google Compute Engine"`).
- `--severity`: Filter by severity (e.g., `--severity high`).

**Example:**

```sh
cloud-status history --service "Google Cloud Storage" --limit 3
```

## MCP Server Usage

The application can also run as an MCP server, exposing its functionality as a tool.

### Stdio Mode

To start the server over stdio, use the `--mcp` flag:

```sh
cloud-status --mcp
```

### HTTP Mode

To start the server over HTTP, use the `--mcp-http` flag with a listen address:

```sh
cloud-status --mcp-http :8080
```

### MCP Tool: `get_cloud_status`

The server exposes a single tool with the following details:

- **Name:** `get_cloud_status`
- **Description:** Retrieves status information for Google Cloud services. It can get the current status from the Atom feed or historical incidents from the JSON feed.
- **Parameters:**
    - `source` (string, **required**): Specifies the data to retrieve. Must be either `"current"` or `"history"`.
    - `service_filter` (string, optional): When `source` is `"history"`, this filters incidents by the exact service name (e.g., `"Google Compute Engine"`).
    - `severity_filter` (string, optional): When `source` is `"history"`, this filters incidents by severity (`low`, `medium`, `high`).
    - `limit` (int, optional): When `source` is `"history"`, this limits the number of incidents returned.
