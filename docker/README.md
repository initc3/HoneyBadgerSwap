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

* [`sbellem/mpspdz:randomshamirprep`](https://hub.docker.com/layers/139098383/sbellem/mpspdz/randomshamirprep/images/sha256-eab1c1ab1641ced1a2e89b4c0c6552948e4e1c9ade08d66b942b625169015f09?context=explore)
* [`sbellem/mpspdz:maliciousshamirparty`](https://hub.docker.com/layers/139098336/sbellem/mpspdz/maliciousshamirparty/images/sha256-ac1f16820a16149d6991c916dd2a686139c11d28c4db0e8c861f457796537f8d?context=explore)

Note that in order to run these pre-built binaries, `mpir` is necessary. See
the `Dockerfile` under the root of this repository for examples of how these
images can be used.
