import time
from .logger import log
import requests
from .config import config
from .tailer import event_queue
import json
import uuid
from datetime import datetime

backend = config.Backend()
agent = config.agent()


def sender_worker(interval: str):
    while True:
        time.sleep(interval)
        batch = []
        while not event_queue.empty():
            try:
                batch.append(event_queue.get_nowait())
            except:
                event_queue.empty()
                break
            print(f"this is batch {batch}")
        if batch:
            send_batch(batch)


def send_batch(batch: str):
    log.info("sending log streams...")

    stream_id = f"{agent['stream']}:{agent['id']}:{uuid.uuid4()}"

    stream = {
        "stream_id": stream_id,
        "worker_name": agent["name"],
        "stream_count": len(batch),
        "stream_size": len(json.dumps(batch)),
        "send_at": datetime.now().isoformat() + "Z",
        "events": batch,
    }

    try:
        r = requests.post(backend["url"], json=stream, timeout=5)
        if r.status_code == 200:
            log.success(r.text)
        else:
            log.info(f"backend sending {r.status_code} status and {r.text}")
    except Exception as e:
        log.error("Error" + str(e))
