# 🐳 Docker Color Output

This package allows you to run Docker commands and get color output.

## Requirements

- **php** >= 7.0

## ⚡️ Installation

Clone the repo and add the following lines to `~/.bash_aliases` (`~/.bashrc`, `~/.zshrc`) file.

```bash
DOCKER_COLOR_OUTPUT_PATH='/absolute-path-to-cloned-repo'
alias di="$DOCKER_COLOR_OUTPUT_PATH/bin/docker-images"
alias dps="$DOCKER_COLOR_OUTPUT_PATH/bin/docker-ps"
alias dcps="$DOCKER_COLOR_OUTPUT_PATH/bin/docker-compose-ps"
```

**Note:** change `DOCKER_COLOR_OUTPUT_PATH` to your absolute path where you cloned the repository.

## 💥 Usage

You can also pass all arguments as you pass to the command.

### 💡 Docker images

```bash
# The 'docker images' command will be called.
di
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

### 💡 Docker ps

```bash
# The 'docker ps' command will be called.
dps
```

```bash
# The 'docker ps -a' command will be called.
dps -a
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

### 💡 Docker compose

```bash
# The 'docker-compose ps' command will be called.
dcps
```

![docker-compose ps](https://user-images.githubusercontent.com/5787193/93581165-6fe09580-f9a9-11ea-914c-1081eab52a79.png)
