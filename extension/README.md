# cloud-status Gemini CLI Extension

This directory contains a Gemini CLI extension for the `cloud-status` tool.

## Installation

To install this extension, run the following command:

```
gemini extensions install extension/
```

## Prerequisites

This extension assumes that you have the `cloud-status` binary in your system's `PATH`.

## Usage

Once installed, the `cloud-status` MCP server will be started automatically when you run `gemini`.

You can then interact with the `cloud-status` tool using the Gemini CLI.

```
What's the service health of google cloud?
```


## Custom Commands

This extension also adds the following custom commands:

*   `/gcp:current`: Gets the current status of Google Cloud services.
*   `/gcp:history`: Gets the historical status of Google Cloud services.
*   `/gcp-status`: An alias for `/gcp:current`.

These commands are defined in the `commands/` directory. For more information on custom commands, see the [Custom Commands](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/commands.md#custom-commands) and [Extensions](https://github.com/google-gemini/gemini-cli/blob/main/docs/extension.md) documentation.
