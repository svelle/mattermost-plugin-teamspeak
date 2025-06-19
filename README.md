# Mattermost TeamSpeak Plugin

A Mattermost plugin that integrates TeamSpeak 3 server information directly into your Mattermost workspace. View online users and channels from your TeamSpeak server without leaving Mattermost.

## Features

- 📊 Real-time display of TeamSpeak channels and online users
- 🔄 Automatic updates every 30 seconds
- 👥 See user details including away status, idle time, and client information
- 📱 Integrated into Mattermost's right-hand sidebar
- 🔐 Secure connection using TeamSpeak's WebQuery API

## Requirements

- Mattermost Server v6.2.1+
- TeamSpeak 3 Server with WebQuery enabled
- Go 1.22+ (for building from source)
- Node.js and npm (for building from source)

## Installation

### From Release

1. Download the latest release from the [releases page](https://github.com/svelle/mattermost-plugin-teamspeak/releases)
2. In Mattermost, go to **System Console > Plugins > Plugin Management**
3. Click **Upload Plugin** and select the downloaded file
4. Enable the plugin

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/svelle/mattermost-plugin-teamspeak.git
   cd mattermost-plugin-teamspeak
   ```

2. Build the plugin:
   ```bash
   make dist
   ```

3. Upload the generated `dist/*.tar.gz` file to your Mattermost server

## Configuration

After installing the plugin, configure it in **System Console > Plugins > TeamSpeak**:

### Settings

| Setting | Description | Default |
|---------|-------------|---------|
| **WebQuery URL** | The URL of your TeamSpeak server's WebQuery API | `http://127.0.0.1:10080` |
| **API Key** | Your TeamSpeak WebQuery API key | *(required)* |
| **Server ID** | TeamSpeak virtual server ID (if running multiple servers) | `1` |

### Enabling TeamSpeak WebQuery

1. Connect to your TeamSpeak server as ServerQuery Admin
2. Enable the WebQuery module:
   ```
   serveredit virtualserver_weblist_enabled=1
   ```
3. Generate an API key:
   ```
   apikeyadd scope=manage lifetime=0
   ```
4. Note the generated API key for plugin configuration

For detailed WebQuery setup, see the [TeamSpeak WebQuery documentation](https://community.teamspeak.com/t/webquery-discussion-help-3-12-0-onwards/7184).

## Usage

1. Click the TeamSpeak icon in the channel header
2. The right-hand sidebar will open showing:
   - All TeamSpeak channels in hierarchical order
   - Online users in each channel
   - User status indicators (away, muted, etc.)

The display updates automatically every 30 seconds to reflect the current server state.

## Development

### Prerequisites

- Go 1.22+
- Node.js 16+
- Make

### Building

```bash
# Install dependencies
make webapp/node_modules

# Run tests
make test

# Check code style
make check-style

# Build for development
make dist

# Watch mode for webapp development
make watch
```

### Project Structure

```
├── server/          # Go backend
│   ├── plugin.go    # Main plugin logic
│   └── configuration.go # Settings management
├── webapp/          # React frontend
│   └── src/
│       ├── components/
│       │   ├── channel_header_button/
│       │   └── ts3sidebar/
│       └── index.tsx # Plugin entry point
└── build/           # Build scripts and tools
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/svelle/mattermost-plugin-teamspeak/issues)
- **Mattermost Plugin Documentation**: [developers.mattermost.com](https://developers.mattermost.com/extend/plugins/)

## Acknowledgments

- TeamSpeak is a registered trademark of TeamSpeak Systems GmbH
- Built using the [Mattermost Plugin Starter Template](https://github.com/mattermost/mattermost-plugin-starter-template)