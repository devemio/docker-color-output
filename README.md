<p align="center">
  <img src="https://user-images.githubusercontent.com/5787193/161379988-21c856ef-839a-433e-b014-e81042adac6d.png">
</p>

<p align="center">
  <a href="https://github.com/devemio/docker-color-output/actions/workflows/go.yml"><img src="https://img.shields.io/github/actions/workflow/status/devemio/docker-color-output/go.yml?branch=main"></a>
  <a href="https://codecov.io/gh/devemio/docker-color-output"><img src="https://img.shields.io/codecov/c/gh/devemio/docker-color-output" alt="Coverage"></a>
  <a href="https://github.com/devemio/docker-color-output/releases"><img src="https://img.shields.io/github/downloads/devemio/docker-color-output/total?color=brightgreen" alt="Downloads"></a>
  <a href="https://github.com/devemio/docker-color-output/releases"><img src="https://img.shields.io/github/v/release/devemio/docker-color-output" alt="Release"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-green.svg" alt="License"></a>
</p>


## 👨‍💻 Installation

This package uses a pipeline and allows you to color the **Docker** output.

You can find all compiled binaries on the [releases](https://github.com/devemio/docker-color-output/releases) page.

#### 🍏 Mac

```bash
brew install dldash/core/docker-color-output
```

#### 🏝 Linux

```bash
sudo add-apt-repository ppa:dldash/core
sudo apt update
sudo apt install docker-color-output
```

#### 🏔️ Windows

You can download the Windows build from the [releases](https://github.com/devemio/docker-color-output/releases) page.

## 😺 Configuration

You can customize the color palette to your liking, to do this you need to run `docker-color-output` with the `-c`
flag and specify the path to the configuration file. You can overwrite only the colors you want, the rest will be
used by default.

```shell
docker-color-output -c ~/.config/docker-color-output/config.json
```

Default configuration file:

```json
{
  "colors": {
    "reset": "\u001b[0m",
    "black": "\u001b[0;30m",
    "darkGray": "\u001b[1;30m",
    "red": "\u001b[0;31m",
    "lightRed": "\u001b[1;31m",
    "green": "\u001b[0;32m",
    "lightGreen": "\u001b[1;32m",
    "brown": "\u001b[0;33m",
    "yellow": "\u001b[1;33m",
    "blue": "\u001b[0;34m",
    "lightBlue": "\u001b[1;34m",
    "purple": "\u001b[0;35m",
    "lightPurple": "\u001b[1;35m",
    "cyan": "\u001b[0;36m",
    "lightCyan": "\u001b[1;36m",
    "lightGray": "\u001b[0;37m",
    "white": "\u001b[1;37m"
  }
}
```

## 📚 Usage

### 🪄 Aliases

Use these bash [functions](bash/aliases.sh) to improve the user experience.

### 💡 docker images

```bash
di # alias
```

```bash
docker images [--format] | docker-color-output
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

#### 💡 docker ps

```bash
dps # alias
```

```bash
docker ps [-a] [--format] | docker-color-output
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

#### 💡 docker compose ps

⚠️ The latest version works with docker-compose `2.x`.

```bash
dcps # alias
```

```bash
docker compose ps | docker-color-output
```

![docker compose ps](https://user-images.githubusercontent.com/5787193/93630916-7267dd00-f9f3-11ea-9521-e69152fa86f1.png)
