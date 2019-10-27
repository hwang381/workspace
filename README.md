# workspace
Multi-repo context switcher

`workspace` manages your multi-repo project and makes it easy to context switch them for different ongoing work items while respecting inter-repo dependencies

## Prerequisites
* `Go 1.11` or above (for module support)
* Add `$GOBIN` to your `$PATH`

## Install
```bash
go install
```

## Use
```bash
workspace
```

## Configure
```bash
mkdir -p ~/.workspace
cp config.example.json ~/.workspace/config.json
```
Then edit `~/.workspace/config.json`

## Caveat
* In order to pull latest code for a branch, it by default reads `~/.ssh/id_rsa` as SSH key for your git repos
