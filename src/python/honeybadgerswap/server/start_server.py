import asyncio

from .Server import http_server


async def main():
    tasks = []
    tasks.append(asyncio.ensure_future(http_server()))

    for task in tasks:
        await task


# TODO Use argparse to accept config file to use.
if __name__ == "__main__":
    asyncio.run(main())
