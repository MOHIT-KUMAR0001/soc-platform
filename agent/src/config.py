import yaml  # type:ignore
from .logger import Logging

log = Logging()


class Define_Config:

    def __init__(self, source: str):
        with open(source, "r") as cfg:
            self._config = yaml.safe_load(cfg)
            if not self._config:
                log.error("file not found !")

    def get_config(self, name: str):
        return self._config.get(name, {})

    def agent(self):
        return self.get_config("agent") or {}

    def Backend(self):
        return self.get_config("backend") or {}

    def logs(self):
        return self.get_config("logs") or {}

    def buffer(self):
        return self.get_config("buffer") or {}

    def runtime(self):
        return self.get_config("runtime") or {}


config = Define_Config("config.yaml")
