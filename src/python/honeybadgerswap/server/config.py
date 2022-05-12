from pathlib import Path
import os

import toml
from pydantic import BaseSettings

CONFIG_PATH_ENV = "HBSWAP_SERVER_CONFIG"
DEFAULT_CONFIG_PATH = "/opt/hbswap/conf/server.toml"


def toml_config_settings_source(settings: BaseSettings):
    return toml.load(Path(os.getenv(CONFIG_PATH_ENV, DEFAULT_CONFIG_PATH)))


class Settings(BaseSettings):
    N: int
    T: int
    LeaderHostname: str
    Servers: list
    EthNode: dict
    NODE_ID: int

    class Config:
        @classmethod
        def customise_sources(cls, init_settings, env_settings, file_secret_settings):
            return (
                init_settings,
                toml_config_settings_source,
                env_settings,
                file_secret_settings,
            )


settings = Settings()
