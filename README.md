# Activator

![release](https://github.com/tomshaw/activator/actions/workflows/release.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/tomshaw/activator?status.svg)](https://godoc.org/github.com/tomshaw/activator)
![Supported Version](https://img.shields.io/badge/go%20version-%3E%3D1.19-turquoise)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomshaw/activator)](https://goreportcard.com/report/github.com/tomshaw/activator)
[![License](https://img.shields.io/github/license/tomshaw/activator)](https://github.com/tomshaw/activator/blob/master/LICENSE)

`activator` is a command line font management tool used with [Fontastic](https://github.com/tomshaw/fontastic) Electron Font Manager.

## Installation

```sh
go install github.com/tomshaw/activator@latest
```

Activator has been tested and supports Go versions >=1.19.x.

## Usage

Run `activator -h` to print help instructions.

---

## Font Installation

> The following commands must be used with elevated or administrative privileges.

### Installing Fonts

```sh
activator install "Font-Name.otf"
``` 

### Uninstalling Fonts

```sh
activator uninstall "Font-Name.otf"
``` 

### Multiple Fonts

Installing multiple fonts is supported by separating paths by a space.

```sh
activator install "Font-Name.otf" "Font-Name Bold.otf" "Font-Name Bold Italic.otf"
``` 

### Temporary Installation

Windows supports installing fonts temporarily. Fonts are removed after a system reboot. 

```sh
activator install --temporary=true "Font-Name.otf"
``` 

### Temp Uninstallation 

```sh
activator uninstall --temporary=true "Font-Name.otf"
``` 

---

## Font Management

Included utility commands to help manage system fonts.

### Finding Fonts

Finds and prints fonts including sub folders.

```sh
activator fonts find --root "C:\Fonts"
``` 

### Copying Fonts

Copies fonts from source to destination folder.

```sh
activator fonts copy --source "C:\Fonts" --destination "C:\Dest"
``` 

---

## Contributions

1. Fork the repo
2. Clone the fork (`git clone git@github.com:YOUR_USERNAME/activator.git && cd activator`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Make changes and add them (`git add --all`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create a pull request

## License

See [LICENSE](https://github.com/tomshaw/activator/blob/master/LICENSE).
