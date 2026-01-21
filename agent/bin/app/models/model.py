from datetime import datetime
from config.config import config
import uuid
import json
import msgpack

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


def stream_model(batch : list[dict]):
        stream_id = f"{agent['stream']}:{agent['id']}:{uuid.uuid4()}"
        stream = {
        "stream_id": stream_id,
        "worker_name": agent["name"],
        "stream_count": len(batch),
        "stream_size": len(json.dumps(batch)),
        "send_at": datetime.now().isoformat() + "Z",
        "events": batch,
    }
        print(len(stream['events']))
        return msgpack.packb(stream, use_bin_type=True)

