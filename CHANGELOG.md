# Changelog

All notable changes to the Mattermost TeamSpeak Plugin will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 0.2.0 - 2025-06-19

### Added
- Added CLAUDE.md documentation for AI-assisted development guidance
- Added comprehensive README.md with installation, configuration, and usage instructions
- Added GitHub Actions CI workflow replacing CircleCI
- Added `.nvmrc` file for Node version management
- Added VSCode settings for consistent development environment
- Added TeamSpeak 5 SVG icon (`public/teamspeak5.svg`)
- Added improved logging functionality in build tools (`build/pluginctl/logs.go`)
- Added tests for logging functionality (`build/pluginctl/logs_test.go`)

### Changed
- **BREAKING**: Fixed ServerID configuration field - changed from `ServerId` to `ServerID` to match Go struct
- **BREAKING**: Fixed ServerID default value from string `"1"` to number `1`
- Shortened plugin display name from "Mattermost Teamspeak Plugin" to "Teamspeak"
- Updated to latest Mattermost plugin template structure
- Modernized build system and Makefile
- Updated all npm dependencies to latest versions
- Improved ESLint configuration for better code quality
- Enhanced webpack configuration for better bundling
- Updated Go module dependencies
- Improved plugin manifest handling in build tools

### Fixed
- Fixed configuration loading error due to JSON type mismatch
- Fixed workflow permission issues in CI
- Fixed build issues in Makefile
- Fixed macOS compatibility issues (multiple architecture support)
- Fixed integer width handling for better cross-platform compatibility
- Fixed channel sorting to use TeamSpeak's channel order correctly
- Fixed residual data cleanup issues
- Fixed TypeScript type definitions and imports

### Removed
- Removed CircleCI configuration (replaced with GitHub Actions)
- Removed auto-generated manifest files from version control
- Removed deprecated build sync tooling
- Removed unnecessary build dependencies

### Technical Details
- Updated minimum Mattermost server version requirement to 6.2.1
- Migrated from CircleCI to GitHub Actions for CI/CD
- Improved error handling in server plugin
- Better type safety in TypeScript components
- Reduced bundle size through dependency updates

## Previous Releases

For changes prior to commit d9f3a4722b726e87c0e26cf488278f9761054490, please refer to the git history.
