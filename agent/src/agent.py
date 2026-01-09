import time
import threading
from .config import config
from .worker import sender_worker
from .tailer import tail_file

logs = config.logs()

def start():
    # 1. Start the sender worker thread
    # batch_interval=2 means it flushes the queue every 2 seconds
    threading.Thread(target=sender_worker, args=(2.0,), daemon=True).start()

    # 2. Start the tailing threads
    for log in logs:
        t = threading.Thread(
            target=tail_file, args=(log["source"], log["path"]), daemon=True
        )
        t.start()

    # Keep main thread alive
    while True:
        time.sleep(1)
