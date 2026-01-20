import time
import threading
from config.config import config
from core.worker import sender_worker
from core.tailer import tail_file

logs = config.logs()

def start():
    '''
    # 1. Start the sender worker thread
    # batch_interval=2 means it flushes the queue every 2 seconds
    '''
    threads = []

    threading.Thread(target=sender_worker, args=(2.0,), daemon=True).start()

    # 2. append all thread into threads
    for log in logs:
        t = threading.Thread(target=tail_file, args=(log["source"], log["path"]), daemon=True)
        threads.append(t)

    for thread in threads:
        thread.start()

    for thread in threads:
        thread.join()
