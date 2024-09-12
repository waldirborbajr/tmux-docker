# tmux-docker

A tmux plugin that displays the names of running Docker containers on a remote server in the status bar, refreshing every 10 seconds.

## Requirements

- Go 1.16 or higher
- tmux 3.0 or higher
- Docker running on a remote server with API enabled

## Installation

### Manual Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/tmux-docker.git ~/.tmux/plugins/tmux-docker
   ```

2. Build the plugin:

   ```
   cd ~/.tmux/plugins/tmux-docker
   go build -o tmux-docker
   ```

3. Add the following lines to your `~/.tmux.conf` file:

   ```
   # Set the remote Docker host IP and port
   DOCKER_HOST="192.168.1.100:2375"
   
   # Add the plugin to your status line
   set -g status-right '#(~/.tmux/plugins/tmux-docker/tmux-docker $DOCKER_HOST)'
   ```

   Replace `192.168.1.100:2375` with the IP address and port of your remote Docker host.

   You can also add it to the left side of the status bar if preferred:

   ```
   set -g status-left '#(~/.tmux/plugins/tmux-docker/tmux-docker $DOCKER_HOST)'
   ```

4. Reload tmux configuration:

   ```
   tmux source-file ~/.tmux.conf
   ```

## Remote Docker Configuration

To allow the plugin to connect to a remote Docker host, you need to configure the Docker daemon on the remote server to accept API requests:

1. On the remote server, edit the Docker daemon configuration file (usually `/etc/docker/daemon.json`):

   ```json
   {
     "hosts": ["tcp://0.0.0.0:2375", "unix:///var/run/docker.sock"]
   }
   ```

   This allows Docker to accept connections on all network interfaces on port 2375.

2. Restart the Docker service:

   ```
   sudo systemctl restart docker
   ```

**Note:** Exposing the Docker API like this can be a security risk. Ensure you have proper firewall rules in place and consider using TLS for encryption.

## Usage

Once installed and configured, the plugin will automatically start displaying running container names from the remote Docker host in your tmux status bar. The list will refresh every 10 seconds.

## Configuration

### Catppuccin Theme Integration

If you're using the Catppuccin theme, you can integrate the plugin by adding the following to your theme configuration:

```
set -g @catppuccin_status_modules_right "... docker ..."
set -g @catppuccin_docker_color "$thm_cyan"
set -g @catppuccin_docker_icon "🐳"
```

Replace `...` with your other desired modules. Adjust the color and icon as needed.

Then, in your `~/.tmux.conf`, add:

```
set -g status-right '#{docker}'
```

## Building from Source

To build the plugin from source:

1. Navigate to the plugin directory:

   ```
   cd /wks/cavelab/tmux-docker
   ```

2. Build the Go binary:

   ```
   go build -o tmux-docker
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
