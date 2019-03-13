# Docker color output

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

![docker images](https://user-images.githubusercontent.com/5787193/54311065-e5192080-45e4-11e9-8973-25ae5b12bea6.png)

### Docker ps

```bash
# docker ps
dps
```

```bash
# docker ps -a
dps -a
```

![docker ps](https://user-images.githubusercontent.com/5787193/54311067-e5192080-45e4-11e9-8fbb-6d30662656d4.png)

### Docker compose

```bash
# docker-compose ps
dcps
```

![docker-compose ps](https://user-images.githubusercontent.com/5787193/54311063-e4808a00-45e4-11e9-8554-9704207a0db0.png)