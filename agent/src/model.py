from datetime import datetime
from .config import config

agent = config.agent()


def create_event(source: str, raw: str) -> dict:
    return {
        "event_type": "log",
        "agent_id": agent["id"],
        "hostname": agent["name"],
        "source": source,
        "timestamp": datetime.now().isoformat() + "Z",
        "raw": raw,
    }
