# clipboard-url-saver

Automatically save urls found in clipboard to file.

## Installation

To install the program [Go](https://go.dev) is required.

After installing Go, the program can be installed by running:

```sh
go install github.com/hougesen/clipurl@latest
```

The program can now be run by either calling the program directly (Most likely `$HOME/go/bin/clipurl`) or by setting up a path to the go install path in your `.bashrc` and calling the name of the program (`clipurl`).

```sh
# .bashrc
export PATH=${PATH}:$(go env GOPATH)/bin
```

### Platform Specific Dependencies

-   macOS: require Cgo, no dependency
-   Linux: require X11 dev package. For instance, install `libx11-dev` or `xorg-dev` or `libX11-devel` to access X window system.
-   Windows: no Cgo, no dependency

## Usage

### Tracking clipboard

Start listening to the clipboard by running

```sh
clipurl start
```

#### Disclaimer

The program does not automatically start on boot. If that is desired you can follow this [guide](https://www.howtogeek.com/687970/how-to-run-a-linux-program-at-startup-with-systemd/) to setup startup on boot.

### URL History

To get a list of all found urls run

```sh
clipurl history
```
