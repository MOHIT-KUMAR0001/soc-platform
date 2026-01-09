import time
from .model import create_event
import queue

event_queue = queue.Queue()

def tail_file(source: str, path: str):
    print("this is tailer file")
    with open(path, "r") as file:
        file.seek(0, 2)

        while True:
            line = file.readline()
            if not line:
                time.sleep(0.1)
                continue
            print(line)
            print(create_event(source,line))
            event_queue.put(create_event(source, line))
