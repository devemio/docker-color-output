# Docker color output

This package allows you to run Docker commands and get color output.

## Requirements

- **php** >= 7.0

## Installation

Clone the repo and add the following lines to `~/.bash_aliases` (`~/.bashrc`, `~/.zshrc`) file.

```bash
DOCKER_COLOR_OUTPUT_PATH='/absolute-path-to-cloned-repo'
alias di="$DOCKER_COLOR_OUTPUT_PATH/bin/docker-images"
alias dps="$DOCKER_COLOR_OUTPUT_PATH/bin/docker-ps"
alias dcps="$DOCKER_COLOR_OUTPUT_PATH/bin/docker-compose-ps"
```

**Note:** change `DOCKER_COLOR_OUTPUT_PATH` to your absolute path where you cloned the repository.

## Usage

You can also pass all arguments as you pass to the command.

### Docker images

```bash
# The 'docker images' command will be called.
di
```

![docker images](https://user-images.githubusercontent.com/5787193/54311065-e5192080-45e4-11e9-8973-25ae5b12bea6.png)

### Docker ps

```bash
# The 'docker ps' command will be called.
dps
```

```bash
# The 'docker ps -a' command will be called.
dps -a
```

![docker ps](https://user-images.githubusercontent.com/5787193/54311067-e5192080-45e4-11e9-8fbb-6d30662656d4.png)

### Docker compose

```bash
# The 'docker-compose ps' command will be called.
dcps
```

![docker-compose ps](https://user-images.githubusercontent.com/5787193/54311063-e4808a00-45e4-11e9-8554-9704207a0db0.png)
