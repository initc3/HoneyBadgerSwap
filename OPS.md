# Dev Ops

## git submodule reminders

to update to latest submodule code:

```shell
git submodule update --remote
```

to change branch:

```shell
git config -f .gitmodules submodule.MP-SPDZ.branch new-branch-name
git submodule update --remote
```
