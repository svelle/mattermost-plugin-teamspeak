# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Mattermost plugin that integrates TeamSpeak 3 server information, showing channels and online users within Mattermost. The plugin consists of:
- Go backend server that connects to TeamSpeak's WebQuery API
- React/TypeScript frontend that displays TeamSpeak data in Mattermost's right-hand sidebar

## Common Development Commands

### Build Commands
- `make all` - Runs style checks, tests, and builds the complete plugin bundle
- `make dist` - Builds and bundles the plugin for distribution
- `make server` - Builds only the server component
- `make webapp` - Builds only the webapp component
- `make watch` - Builds and auto-reloads webapp changes during development

### Testing & Linting
- `make test` - Runs all tests (server and webapp)
- `make check-style` - Runs linting for both Go and TypeScript/React code
- `cd webapp && npm run lint` - Run only webapp linting
- `cd webapp && npm run check-types` - TypeScript type checking
- `cd webapp && npm run fix` - Auto-fix linting issues in webapp

### Plugin Management
- `make deploy` - Builds and deploys plugin to a Mattermost server
- `make enable` - Enables the plugin on the server
- `make disable` - Disables the plugin
- `make reset` - Disables and re-enables the plugin

### Version Management
- `make patch` - Bump patch version (e.g., 1.0.0 → 1.0.1)
- `make minor` - Bump minor version (e.g., 1.0.0 → 1.1.0)
- `make major` - Bump major version (e.g., 1.0.0 → 2.0.0)

## Architecture Overview

### Server Component (Go)
The server communicates with TeamSpeak's WebQuery API and consists of:
- **plugin.go**: Main plugin implementation with background worker that updates channel/client lists every 30 seconds
- **configuration.go**: Plugin settings management (WebQuery URL, API Key, Server ID)
- **HTTP endpoint**: `/plugins/{id}/info` serves TeamSpeak data to the webapp

### Webapp Component (React/TypeScript)
The frontend displays TeamSpeak information in Mattermost:
- **index.tsx**: Plugin entry point, registers UI components with Mattermost
- **channel_header_button**: Toggle button in channel header to show/hide TeamSpeak sidebar
- **ts3sidebar**: Components for displaying channels and clients
  - ChannelList.jsx: Main container that fetches and displays data
  - Channel.jsx: Individual channel display
  - ClientList.jsx: List of clients in a channel
  - Client.jsx: Individual client display

### Plugin Configuration
The plugin requires configuration through Mattermost's plugin settings:
1. **WebQuery URL**: TeamSpeak server's WebQuery API endpoint (default: http://127.0.0.1:10080)
2. **API Key**: Authentication key for WebQuery API
3. **Server ID**: TeamSpeak virtual server ID (default: 1)

### Build System
- Uses Mattermost's plugin build framework with Makefile
- Supports multiple architectures (linux/darwin/windows, amd64/arm64)
- Webpack for webapp bundling
- Automated manifest generation and validation

## Key Technical Details

- TeamSpeak API integration uses WebQuery (REST API)
- Data updates occur every 30 seconds via background goroutine
- Channel ordering is maintained using TeamSpeak's channel_order field
- The plugin supports TeamSpeak 3 features including channel flags, client status, and permissions
- All TeamSpeak data types use custom JSON unmarshaling due to string-encoded numbers in the API