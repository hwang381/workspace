# workspace
Multi-repo context switcher

`workspace` manages your multi-repo project and makes it easy to context switch them for different ongoing work items while respecting inter-repo dependencies

## Install
Get the binaries from [releases](https://github.com/hwang381/workspace/releases), then rename it to `workspace` and put it somewhere under your `$PATH`.

## Configure
```bash
mkdir -p $HOME/.workspace
cp config.example.json $HOME/.workspace/config.json
```

Then edit `$HOME/.workspace/config.json`

## Use
Type `workkspace` for help info

## Develop

### Prerequisites
* `Go 1.11` or above (for module support)

### Test build
Run `go build`, and then find the `workspace` binary under the repo root
