# Test Keys for MP-SPDZ
**JUST for development & testing purposes!**

**NEVER PUBLISH PRIVATE KEYS!**

This directory contains public and private keys used by MP-SPDZ.

To generate new keys you can use the `genkeys.yml` file,
**from the root of the repo**:

```shell
docker-compose -f testkeys/genkeys.yml up
```

or from the directory where this `README.md` is:

```shell
docker-compose -f genkeys.yml up
```

By default it will create keys for 4 players. If you need to generate a different
number replace `"4"` in the following line in `genkeys.yml`:

```yml
    command: ["4"]
```

For 100 players:

```yml
    command: ["100"]
```

Note that the file ownership will `root` and you can change it to your user with:

```shell
sudo chown -R `id -un`:`id -gn` testkeys/
```
