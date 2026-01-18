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

Docker Color Output is a lightweight CLI that adds readable, customizable colors to Docker command output. It reads from stdin, applies a YAML ruleset with per-column logic, and writes clean, aligned tables to stdout.

## Features

- Cross-platform: macOS, Linux, Windows.
- YAML rules: per-column colors, matches, and thresholds.
- Pipeline-friendly: works as a simple pipe in shell workflows.

## Installation

Binaries are available on the [releases](https://github.com/devemio/docker-color-output/releases) page.

#### Mac

```bash
brew install dldash/core/docker-color-output
```

#### Linux

```bash
sudo add-apt-repository ppa:dldash/core
sudo apt update
sudo apt install docker-color-output
```

#### Windows

Download the Windows build from the [releases](https://github.com/devemio/docker-color-output/releases) page.

## Quick Start

```bash
docker ps | docker-color-output
```

```bash
docker images [--format] 2>/dev/null | docker-color-output
```

> [!NOTE]
> On Docker 29+, `docker images` writes extra output to stderr, so redirect it (`2>/dev/null`) when piping into `docker-color-output`.

## Configuration

Pass a custom YAML config with `-c`:

```bash
docker-color-output -c ~/.config/docker-color-output/config.yaml
```

See `internal/config/default.yaml` for the full default configuration, including all commands and rules.

Example override:

```yaml
layout:
  headerColor: yellow
rules:
  commands:
    dockerPs:
      rules:
        "STATUS":
          pipeline:
            - type: match
              when:
                - contains: "Exited"
                  color: red
              default:
                color: brown
```

## Usage

If you want shortcuts, use the bash aliases in `bash/aliases.sh`.

Tested with Docker `29.1.3`.

### docker images / docker image ls

```bash
di # alias
```

```bash
docker images [--format] 2>/dev/null | docker-color-output
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

### docker ps

```bash
dps # alias
```

```bash
docker ps [-a] [--format] | docker-color-output
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

### docker compose ps

```bash
dcps # alias
```

```bash
docker compose ps | docker-color-output
```

![docker compose ps](https://user-images.githubusercontent.com/5787193/93630916-7267dd00-f9f3-11ea-9521-e69152fa86f1.png)

### docker stats

```bash
ds # alias
```

```bash
docker stats [--no-stream] | docker-color-output
```

![docker stats](https://github.com/devemio/docker-color-output/assets/5787193/a3134ae9-707b-4ad7-a5ea-765150d535e8)

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
