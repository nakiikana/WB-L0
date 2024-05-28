# WB-L0
## Overview
This server application is designed to handle various operations including connecting to a PostgreSQL database, managing a cache, initializing a NATS streaming subscription, and publishing messages at regular intervals. It leverages several custom packages for configuration management, caching, repository interactions, service orchestration, and HTTP request handling. Additionally, it utilizes the NATS Streaming client library for message passing and the Ticker package for scheduling periodic tasks.

## Configuration
Configuration settings are loaded from a YAML file located at ../config/config.yaml. Ensure this file is correctly configured with your database connection details, cache settings, and any other relevant configurations.

## Key Features
- Database Operations: Connects to a PostgreSQL database for storing and retrieving data.
- Caching: Implements a caching mechanism to improve performance by reducing database access.
- NATS Streaming Integration: Subscribes to a NATS streaming topic and publishes messages periodically.
- Error Handling and Logging: Utilizes the Logrus library for comprehensive logging and error reporting.

## Error Handling
The server employs robust error handling throughout its execution flow. It logs fatal errors using Logrus, ensuring that the application stops gracefully upon encountering unrecoverable issues.

## Shutdown Process
Upon receiving an interrupt signal (e.g., Ctrl+C), the server initiates a graceful shutdown process. It waits for ongoing operations to complete and notifies subscribers before terminating.