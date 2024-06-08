[![skynx.com](https://github.com/skynx-io/assets/blob/HEAD/images/logo/skynx-logo_black_180x45.png)](https://skynx.com)

[![Discord](https://img.shields.io/discord/654291649572241408?color=%236d82cb&style=flat&logo=discord&logoColor=%23ffffff&label=Chat)](https://skynx.com/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/skynx-io/discussions)
[![X](https://img.shields.io/badge/Follow_on_X-000000?style=flat&logo=x&logoColor=white)](https://x.com/skynxHQ)
[![Mastodon](https://img.shields.io/badge/Follow_on_Mastodon-2f0c7a?style=flat&logo=mastodon&logoColor=white)](https://mastodon.social/@skynx)

Open source projects from [skynx.com](https://skynx.com).

# skynx-cli

[![Go Report Card](https://goreportcard.com/badge/skynx.io/s-cli)](https://goreportcard.com/report/skynx.io/s-cli)
[![Release](https://img.shields.io/github/v/release/skynx-io/s-cli?display_name=tag&style=flat)](https://github.com/skynx-io/s-cli/releases/latest)
[![GitHub](https://img.shields.io/github/license/skynx-io/s-cli?style=flat)](/LICENSE)

This repository contains `skynxctl`, a tool for managing the [skynx](https://skynx.com) SASE platform from the command line.

`skynxctl` is available for a variety of Linux platforms, macOS and Windows.

## Minimun Requirements

`skynxctl` has the same [minimum requirements](https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements) as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

## Getting Started

See [Quick Start](https://skynx.com/docs/platform/getting-started/quickstart/) to learn how to start building your skynx cloud-agnostic architecture.

## Documentation

For the complete skynx platform documentation visit [skynx.com/docs](https://skynx.com/docs/).

## Installation

### Binary Downloads

Linux, macOS and Windows binary downloads are available from the [Releases](https://github.com/skynx-io/s-cli/releases) page.

You can download the pre-compiled binaries and install them with the appropriate tools.

### Linux Installation

#### Linux binary installation with curl

1. Download the latest release.

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/linux/amd64/skynxctl"
    ```

2. Validate the binary (optional).

    Download the skynxctl checksum file:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/linux/amd64/skynxctl_checksum.sha256"
    ```

    Validate the skynxctl binary against the checksum file:

    ```bash
    sha256sum --check < skynxctl_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    skynxctl: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    skynxctl: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install `skynxctl`.

    ```shell
    sudo install -o root -g root -m 0755 skynxctl /usr/local/bin/skynxctl
    ```

    > **Note**:
    > If you do not have root access on the target system, you can still install skynxctl to the `~/.local/bin` directory:
    >
    > ```shell
    > chmod +x skynxctl
    > mkdir -p ~/.local/bin
    > mv ./skynxctl ~/.local/bin/skynxctl
    > # and then append (or prepend) ~/.local/bin to $PATH
    > ```

#### Package Repository

skynx provides a package repository that contains both DEB and RPM downloads.

For DEB-based platforms (e.g. Ubuntu and Debian) run the following to set up a new APT sources.list entry and install `skynx-cli`:

```shell
echo 'deb [trusted=yes] https://repo.skynx.com/apt/ /' | sudo tee /etc/apt/sources.list.d/skynx.list
sudo apt update
sudo apt install skynx-cli
```

For RPM-based platforms (e.g. RHEL, CentOS) use the following to create a repo file and install `skynx-cli`:

```shell
cat <<EOF | sudo tee /etc/yum.repos.d/skynx.repo
[skynx]
name=skynx Repository - Stable
baseurl=https://repo.skynx.com/yum
enabled=1
gpgcheck=0
EOF
sudo yum install skynx-cli
```

#### Homebrew installation on Linux

If you are on Linux and using [Homebrew](https://docs.brew.sh/Homebrew-on-Linux) package manager, you can install the skynx CLI with Homebrew.

1. Run the installation command:

    ```shell
    brew install skynx-io/tap/skynx-cli
    ```

2. Test to ensure the version you installed is up-to-date:

    ```shell
    skynxctl version show
    ```

### macOS Installation

#### macOS binary installation with curl

1. Download the latest release.

    **Intel**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/amd64/skynxctl"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/arm64/skynxctl"
    ```

2. Validate the binary (optional).

    Download the skynxctl checksum file:

    **Intel**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/amd64/skynxctl_checksum.sha256"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/arm64/skynxctl_checksum.sha256"
    ```

    Validate the skynxctl binary against the checksum file:

    ```console
    shasum --algorithm 256 --check skynxctl_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    skynxctl: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    skynxctl: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Make the skynxctl binary executable.

    ```shell
    chmod +x skynxctl
    ```

4. Move the skynxctl binary to a file location on your system `PATH`.

    ```shell
    sudo mkdir -p /usr/local/bin
    sudo mv skynxctl /usr/local/bin/skynxctl
    sudo chown root: /usr/local/bin/skynxctl
    ```

    > **Note**: Make sure `/usr/local/bin` is in your `PATH` environment variable.

#### Homebrew installation on macOS

If you are on macOS and using [Homebrew](https://brew.sh/) package manager, you can install the skynx CLI with Homebrew.

1. Run the installation command:

    ```shell
    brew install skynx-io/tap/skynx-cli
    ```

2. Test to ensure the version you installed is up-to-date:

    ```shell
    skynxctl version show
    ```

### Windows Installation

#### Windows binary installation with curl

1. Open the Command Prompt as Administrator and create a folder for skynx.

    ```shell
    mkdir 'C:\Program Files\skynx'
    ```

2. Download the latest release into the skynx folder.

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/windows/amd64/skynxctl.exe"
    ```

3. Validate the binary (optional).

    Download the skynxctl.exe checksum file:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/windows/amd64/skynxctl.exe_checksum.sha256"
    ```

    Validate the skynxctl.exe binary against the checksum file:

    - Using Command Prompt to manually compare CertUtil's output to the checksum file downloaded:

         ```shell
         CertUtil -hashfile skynxctl.exe SHA256
         type skynxctl.exe_checksum.sha256
         ```

    - Using PowerShell to automate the verification using the -eq operator to get a `True` or `False` result:

         ```powershell
         $($(CertUtil -hashfile skynxctl.exe SHA256)[1] -replace " ", "") -eq $(type skynxctl.exe_checksum.sha256).split(" ")[0]
         ```

4. Append or prepend the folder `C:\Program Files\skynx` to your `PATH` environment variable.
5. Test to ensure the version of skynxctl is the same as downloaded.

    ```shell
    skynxctl version show
    ```

## Artifacts Verification

### Binaries

All artifacts are checksummed and the checksum file is signed with [cosign](https://github.com/sigstore/cosign).

1. Download the files you want and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [Releases](https://github.com/skynx-io/s-cli/releases) page:

2. Verify the signature:

    ```shell
    cosign verify-blob \
        --cert checksums.txt.pem \
        --signature checksums.txt.sig \
        checksums.txt
    ```

3. If the signature is valid, you can then verify the SHA256 sums match the downloaded binary:

    ```shell
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker Images

Our Docker images are signed with [cosign](https://github.com/sigstore/cosign).

Verify the signatures:

```console
COSIGN_EXPERIMENTAL=1 cosign verify skynx/skynxctl
```

## Configuration

The first time you run `skynxctl`, you will be assisted to generate your `skynxctl.yml`. This config file will be located by default at the `$HOME/.skynx` directory.

See the [skynxctl configuration reference](https://skynx.com/docs/platform/reference/skynxctl.yml/) to find all the configuration options.

## Usage

See usage with:

```shell
skynxctl help
```

## Running with Docker

You can also run `skynxctl` as a Docker container. See the example below.

Registries:

- `skynx/skynxctl`
- `ghcr.io/skynx-io/skynxctl`

Example usage:

```shell
docker run --rm -ti -v $HOME/.skynx:/root/.skynx:ro skynx/skynxctl help
```

## Community

Have questions, need support and or just want to talk?

Get in touch with the skynx community!

[![Discord](https://img.shields.io/badge/Join_us_on_Discord-5865F2?style=flat&logo=discord&logoColor=white)](https://skynx.com/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/skynx-io/discussions)
[![X](https://img.shields.io/badge/Follow_on_X-000000?style=flat&logo=x&logoColor=white)](https://x.com/skynxHQ)
[![Mastodon](https://img.shields.io/badge/Follow_on_Mastodon-2f0c7a?style=flat&logo=mastodon&logoColor=white)](https://mastodon.social/@skynx)

## Code of Conduct

Participation in the skynx community is governed by the Contributor Covenant [Code of Conduct](https://github.com/skynx-io/.github/blob/HEAD/CODE_OF_CONDUCT.md). Please make sure to read and observe this document.

Please make sure to read and observe this document. By participating, you are expected to uphold this code.

## License

The skynx open source projects are licensed under the [Apache 2.0 License](/LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fskynx-io%2Fs-cli.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fskynx-io%2Fs-cli?ref=badge_large)
