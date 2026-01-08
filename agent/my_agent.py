import time
import threading
import config
from datetime import datetime
import requests

def tail_file(source: str, path: str, emit_event):
    with open(path, "r") as file:
        file.seek(0, 2)  # move to end of file

        while True:
            line = file.readline()
            if not line:
                time.sleep(0.1)
                continue

            event = {
                "event_type": "log",
                "agent_id": config.Agent['id'],
                "hostname": config.Agent['name'],
                "source": source,
                "timestamp": datetime.utcnow().isoformat() + "Z",
                "raw": line.strip()
            }

            emit_event(event)

def emit_event(event: dict):
    print(event)
    r = requests.post("http://localhost:8000/api/test",
                      json=event , timeout=5)
    print(r.text)

def start():
    threads = []

    for log in config.Logs:
        t = threading.Thread(
            target=tail_file,
            args=(log["source"], log["path"], emit_event),
            daemon=True
        )
        t.start()
        threads.append(t)

    # keep main thread alive
    while True:
        time.sleep(1)

