# Color Docker Output

This package allows you to run Docker commands and get a color output.

## Installation

Clone the repo and add the following line to `.bashrc` or `.zshrc` file.

```bash
source /path/to/repo/docker
```

It creates convenient aliases to call commands.

## Usage

You can also pass all arguments as you pass to the command.

### Docker images

```bash
# docker images
di
```

### Docker ps

```bash
# docker ps
dps
```

```bash
# docker ps -a
dps -a
```

### Docker compose

```bash
# docker-compose ps
dcps
```