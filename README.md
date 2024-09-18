<img src="https://github.com/user-attachments/assets/9fb07638-5907-4915-b9bf-1ca89255a93d" alt="drawing" style="width:100px;"/>

# Tmux Docker Monitor

<p align="center">
  <img width="256" height="256" src="https://github.com/user-attachments/assets/6ede0ea1-383d-4d7b-914d-d56bca4fab4c" />
</p>

#### Under Construction ####

[![codecov](https://codecov.io/github/waldirborbajr/tmux-docker/graph/badge.svg?token=F5A3EQ6RW5)](https://codecov.io/github/waldirborbajr/tmux-docker)

This Go application monitors Docker containers on a remote server and displays their status in the Tmux status bar.

## Recent Changes

1. The configuration file has been renamed from `.env` to `.tmux-docker-env`.
2. The `.tmux-docker-env` file is now located in the user's home directory (`~/`).
3. The binary is now installed in `~/.local/bin/`.
4. All comments in the code have been translated from Portuguese to English.

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/tmux-docker-monitor.git
   cd tmux-docker-monitor
   ```
3. Build the application:
   ```
   go build -o ~/.local/bin/tmux-docker-monitor
   ```

## Configuration

1. Create a `.tmux-docker-env` file in your home directory:
   ```
   touch ~/.tmux-docker-env
   ```

2. Add the following content to the `.tmux-docker-env` file, replacing the values with your actual server details:
   ```
   DOCKER_USER=your_username
   REMOTE_SERVER_IP=your_server_ip
   DOCKER_PASSWORD=your_password
   ```

## Usage

1. Ensure Tmux is running.
2. Execute the Tmux Docker Monitor:
   ```
   ~/.local/bin/tmux-docker-monitor
   ```

The Tmux status bar will now display information about your Docker containers in the following format:

```
Total: X | Up: Y | Down: Z | Died: W
```

Where:
- X is the total number of containers
- Y is the number of running containers
- Z is the number of stopped containers
- W is the number of containers in a "dead" state

## Automating Execution

To have the Tmux Docker Monitor run automatically when you start Tmux, add the following line to your `~/.tmux.conf` file:

```
set-option -g status-interval 60
run-shell "~/.local/bin/tmux-docker-monitor"
```

This will update the status every 60 seconds. Adjust the interval as needed.

## Troubleshooting

If you encounter any issues:

1. Ensure the `.tmux-docker-env` file is correctly placed and formatted.
2. Check that the binary is correctly installed in `~/.local/bin/`.
3. Verify that you have SSH access to the remote server.
4. Ensure the remote server has Docker installed and that your user has permissions to run Docker commands.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
