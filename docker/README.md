# docker images

## [goeth.Dockerfile](https://hub.docker.com/repository/docker/sbellem/go-ethereum)
`go-ethereum` image, commit `cfbb969da` by default. Pre-built at
https://hub.docker.com/repository/docker/sbellem/go-ethereum.

## [openssl.Dockerfile](https://hub.docker.com/repository/docker/sbellem/openssl)
`openssl` on `bullseye-slim` base image. Pre-built at
https://hub.docker.com/repository/docker/sbellem/openssl.

## [mpspdzbuilds.Dockerfile](https://hub.docker.com/repository/docker/sbellem/mpspdz)
MP-SPDZ builds used by HoneyBadgerSwap:`random-shamir.x` and
`malicious-shamir-party.x`. Pre-builts available from
https://hub.docker.com/repository/docker/sbellem/mpspdz. Different builds
are available as different tags:

* [`sbellem/mpspdz:randomshamirprep`](https://hub.docker.com/layers/sbellem/mpspdz/randomshamirshares-c77cc7ab6cc/images/sha256-9be48505316ffb810130b35c7f63ecf4988fb026c05db307e79f933161a29c75?context=explore)
* [`sbellem/mpspdz:maliciousshamirparty`](https://hub.docker.com/layers/sbellem/mpspdz/maliciousshamirparty-c77cc7ab6cc/images/sha256-c1bba9cb1c64036571b0b93ab72cbb7871aa0eccd0bea9158b40cb1079957043?context=explore)

Note that in order to run these pre-built binaries, `mpir` is necessary. See
the `Dockerfile` under the root of this repository for examples of how these
images can be used.

### Building the images
Pay attention to where you run the command relative to the root of this repository.

#### random shamir shares for input masks
Under the root of the repository:

```console
docker build \
    --target random-shamir-shares \
    --tag random-shamir-shares \
    --file docker/mpspdzbuilds.Dockerfile .
```

#### malicious-shamir-party.x

```console
docker build \
    --target malicious-shamir-party \
    --tag malicious-shamir-party \
    --file docker/mpspdzbuilds.Dockerfile .
```
