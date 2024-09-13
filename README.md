# Docker Status CLI

This project is a CLI tool written in Go that connects to a remote server via SSH, retrieves Docker container information, and displays the results in the `tmux` status bar. The CLI provides an overview of the total containers, active containers (Up), inactive containers (Down), and dead containers (Died).

## Features

- Connects to a remote server via SSH.
- Retrieves Docker container information (total, active, inactive, dead).
- Displays status in the `tmux` status bar.

## Requirements

- Go 1.18+
- Docker installed on the remote server.
- `tmux` installed locally.
- A `.env` file containing remote server credentials.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/docker-status-cli.git
   cd docker-status-cli
   ```

2. Install the Go dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory of the project with the following variables:

   ```dotenv
   DOCKER_USER=user         # SSH username
   REMOTE_SERVER_IP=192.168.1.19   # IP or hostname of the remote server
   DOCKER_PASSWORD=password # SSH user's password
   ```

4. Build the CLI:

   ```bash
   go build -o docker-status
   ```

## Usage

Run the CLI with the credentials specified in the `.env` file. Ensure that `tmux` is running:

```bash
./docker-status
```

### Example of the output displayed in the `tmux` status bar:

```
Total: 5 | Up: 3 | Down: 1 | Died: 1
```

## Code Structure

- **`main.go`**: The main file that contains the logic to connect to the remote server, retrieve Docker information, and display the status in `tmux`.
- **Main Functions**:
  - `getServerFromEnv`: Reads connection credentials from the `.env` file.
  - `connectToServer`: Connects to the remote server via SSH.
  - `getDockerInfo`: Executes the `docker ps -a` command on the remote server.
  - `parseDockerOutput`: Parses the Docker output and counts total containers, active, inactive, and dead containers.
  - `displayToTmux`: Updates the `tmux` status bar with the container information.

## Contributions

Feel free to submit pull requests with improvements, bug fixes, or new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

