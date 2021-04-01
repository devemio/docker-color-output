# 🐳 Docker Color Output

This package allows you to colorize the docker output.

## ⚡️ Installation

You can download the binary file from the [releases page](../../releases/latest).

### Aliases

For convenience, you can use the aliases from the [bash/aliases.sh](bash/aliases.sh) file.
Update the `DOCKER_COLOR_OUTPUT_PATH` variable and run the `source aliases.sh` command.

## 💥 Usage

### 💡 Docker images

```bash
di
```

```bash
docker images | dco
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

### 💡 Docker ps

```bash
dps [-a]
```

```bash
docker ps [-a] | dco
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

### 💡 Docker compose

```bash
dcps
```

```bash
docker-compose ps | dco
```

![docker-compose ps](https://user-images.githubusercontent.com/5787193/93630916-7267dd00-f9f3-11ea-9521-e69152fa86f1.png)
