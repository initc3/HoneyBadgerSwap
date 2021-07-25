# BadgerSwapv3
[IC3 Blockchain Summer Camp Project](https://www.initc3.org/events/2021-07-25-ic3-blockchain-summer-camp)

## Development Environment Setup
1. Fork the project
2. Clone your fork, e.g., if your GitHub username is `alice`:

```console
git clone --branch badgerswapv3 git@github.com:alice/HoneyBadgerSwap.git
```

3. Build the image with `docker-compose-dev.yml`:

```console
docker-compose -f docker-compose-dev.yml build
```

## Troubleshooting

### Expired TLS/SSL Certificates
See [./testkeys/README.md](./testkeys/README.md).
