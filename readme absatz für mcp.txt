MCP Server
The mcp-dredger MCP server exposes loaded API specifications as context for AI assistants via the Model Context Protocol (stdio). The server reads the provided OpenAPI and AsyncAPI specifications and makes their contents (endpoints, schemas, channels, etc.) available to the AI assistant.

Build

cd mcp-dredger
go build -o mcp-dredger .
Flags
-f [path] Path to the OpenAPI specification file to load.
-a [path] Path to the AsyncAPI specification file to load.
-examples [path] Path to a directory with additional specifications the server can reference.
IDE Integration
Pre-configured MCP server settings are included in the repository for VS Code (.vscode/mcp.json) and Zed (.zed/settings.json).