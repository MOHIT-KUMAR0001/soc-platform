import yaml # type: ignore

def load_config(path: str) -> dict:
    with open(path, "r") as f:
        config = yaml.safe_load(f)
    return config

config = load_config("config.yaml")

Agent = config.get("agent", {})
Backend = config.get("backend", {})
Logs = config.get("logs", {})
Buffer = config.get("buffer", {})
Runtime = config.get("runtime", {})

def get_config() -> dict:
    return {
        "agent": Agent,
        "backend": Backend,
        "logs": Logs,
        "buffer": Buffer,
        "runtime": Runtime
    }

    