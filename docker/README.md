# docker images

## Ethereum `geth` image
A Debian-based image is hosted on [DockerHub]. The [`Dockerfile`][Dockerfile]
is maintained under [initc3/go-ethereum], which is a fork of
[`ethereum/go-ethereum`][ethereum]. The build is automated via GitHub CI.
See the [`docker.yml`][docker.yml] file. Build results can be viewed at
on the [actions] page.

## MP-SPDZ images
MP-SPDZ builds used by HoneyBadgerSwap:

* `malicious-shamir-party.x`
* `mal-shamir-offline.x`
* `random-shamir.x`

Pre-builts available from https://hub.docker.com/repository/docker/initc3/<*.x>.

* [`initc3/malicious-shamir-party.x`](https://hub.docker.com/repository/docker/initc3/malicious-shamir-party.x)
* [`initc3/mal-shamir-offline.x`](https://hub.docker.com/repository/docker/initc3/mal-shamir-offline.x)
* [`initc3/random-shamir.x`](https://hub.docker.com/repository/docker/initc3/random-shamir.x)


### Building the images
See https://github.com/initc3/MP-SPDZ/blob/dev/DOCKER.md.


[DockerHub]: https://hub.docker.com/repository/docker/initc3/geth
[Dockerfile]: https://github.com/initc3/go-ethereum/blob/dev/debian.Dockerfile
[initc3/go-ethereum]: https://github.com/initc3/go-ethereum
[ethereum/go-ethereum]: https://github.com/ethereum/go-ethereum
[docker.yml]: https://github.com/initc3/go-ethereum/blob/dev/.github/workflows/docker.yml
[actions]: https://github.com/initc3/go-ethereum/actions
