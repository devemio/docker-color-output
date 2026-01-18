<p align="center">
  <img src="https://user-images.githubusercontent.com/5787193/161379988-21c856ef-839a-433e-b014-e81042adac6d.png">
</p>

<p align="center">
  <a href="https://github.com/devemio/docker-color-output/actions/workflows/workflow.yml"><img src="https://img.shields.io/github/actions/workflow/status/devemio/docker-color-output/workflow.yml?branch=main"></a>
  <a href="https://codecov.io/gh/devemio/docker-color-output"><img src="https://img.shields.io/codecov/c/gh/devemio/docker-color-output" alt="Coverage"></a>
  <a href="https://github.com/devemio/docker-color-output/releases"><img src="https://img.shields.io/github/downloads/devemio/docker-color-output/total?color=brightgreen" alt="Downloads"></a>
  <a href="https://github.com/devemio/docker-color-output/releases"><img src="https://img.shields.io/github/v/release/devemio/docker-color-output" alt="Release"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-green.svg" alt="License"></a>
</p>

Docker Color Output is a lightweight tool that enhances Docker's command output by adding vibrant, customizable colors. It processes the standard output through a pipeline, making it easier to read and more visually appealing.

## Features 🚀

- **Cross-Platform Support:** Works on macOS, Linux, and Windows.
- **Customizable Color Schemes:** Easily adjust the colors to your preference.
- **Pipeline Integration:** Transforms Docker command outputs into colorful, structured displays.

## Installation 👨‍💻

You can find all compiled binaries on the [releases](https://github.com/devemio/docker-color-output/releases) page.

#### Mac 🍏

```bash
brew install dldash/core/docker-color-output
```

#### Linux 🐧

```bash
sudo add-apt-repository ppa:dldash/core
sudo apt update
sudo apt install docker-color-output
```

#### Windows 🪟

Download the Windows build from the [releases](https://github.com/devemio/docker-color-output/releases) page.

## Configuration ⚙️

Easily tailor the color scheme and rules to match your personal preferences. Run `docker-color-output` with the
`-c` flag and provide the path to your custom YAML configuration file. Any settings you omit will fall back to
the built-in defaults, while the settings you provide will override them.

```shell
docker-color-output -c ~/.config/docker-color-output/config.yaml
```

##### Default Configuration File

See `internal/config/default.yaml` for the full default configuration, including all rules.

##### Example Override

```yaml
layout:
  headerColor: yellow
rules:
  dockerPs:
    columns:
      "STATUS":
        pipeline:
          - type: match
            when:
              - contains: "Exited"
                color: red
            default:
              color: brown
```

### Silent Mode 🔇

Silent Mode ensures a cleaner output by suppressing error messages. When enabled, if an error occurs,
the tool will simply pass through the original Docker output without displaying any error notifications.

```bash
docker ps | docker-color-output -s
```

## Usage 📚

Enhance your Docker workflow with these handy aliases and enjoy vibrant outputs.

### Aliases 🪄

Utilize the bash functions provided in [aliases.sh](bash/aliases.sh) to streamline your commands.

### docker images 💡

```bash
di # alias
```

```bash
docker images [--format] | docker-color-output
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

#### docker ps 💡

```bash
dps # alias
```

```bash
docker ps [-a] [--format] | docker-color-output
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

#### docker compose ps 💡

> [!NOTE]
> The latest version supports docker compose `2.x`.

```bash
dcps # alias
```

```bash
docker compose ps | docker-color-output
```

![docker compose ps](https://user-images.githubusercontent.com/5787193/93630916-7267dd00-f9f3-11ea-9521-e69152fa86f1.png)

#### docker stats 💡

```bash
ds # alias
```

```bash
docker stats [--no-stream] | docker-color-output
```

![docker stats](https://github.com/devemio/docker-color-output/assets/5787193/a3134ae9-707b-4ad7-a5ea-765150d535e8)

## License 📜

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
