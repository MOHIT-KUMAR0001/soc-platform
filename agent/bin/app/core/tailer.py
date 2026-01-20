import time
from models.model import create_event
import queue

event_queue = queue.Queue()

def tail_file(source: str, path: str):
    '''
    Function open the file then goes to last line and watch the file changes
    '''
    try:
        with open(path, "r") as file:
            file.seek(0, 2)
            while True:
                line = file.readline()
                if not line:
                    time.sleep(0.1)
                    continue
                # print(create_event(source,line))
                event_queue.put(create_event(source, line))
    except:
        print('something went wrong')
        exit()
