# workspace
Multi-repo context switcher

`workspace` manages your multi-repo project and makes it easy to context switch them for different ongoing work items while respecting inter-repo dependencies

## Prerequisites
* `git` is install and is available under your `$PATH`

## Install
Get the binaries from [releases](https://github.com/hwang381/workspace/releases), then rename it to `workspace` and put it somewhere under your `$PATH`.

## Configure
```bash
mkdir -p $HOME/.workspace
cp default.workspace.json $HOME/.workspace/default.workspace.json
```

Then edit `$HOME/.workspace/default.workspace.json`

You can also have multiple workspaces. Just create new configuration files under `$HOME/.workspace` whose filenames end with `.workspace.json`

For example, if you have those two files under `$HOME/.workspace`

```
default.workspace.json
pet_project.workspace.json
```

Then you have two workspaces, one called `default` and another called `pet_project`

By default, actions execute on the `default` workspace. Add the `-w <name of workspace>` flag to execute actions on an alternative workspace

## Use
Type `workkspace` for help info

## Develop

### Prerequisites
* `Go 1.11` or above (for module support)

### Test build
Run `go build`, and then find the `workspace` binary under the repo root
