from core.logger import log
from config.config import config
from core.tailer import event_queue
from models.model import stream_model
from httpx import AsyncClient
import time
import asyncio

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
        if batch:
            asyncio.run(send_batch(batch))


async def send_batch(batch: str):
    log.info("sending log streams...")
    stream = stream_model(batch)
    headers = {
        "Content-Type": "application/msgpack",
        "X-API-KEY": "agent-secret"
    }

    try:
        print(backend['url'])
        async with AsyncClient(timeout=5) as client:
            r = await client.post(url=backend['url'],content=stream,headers=headers)
            if r.status_code == 200:
                log.success(r.text)
            else:
                log.info(f"backend sending {r.status_code} status and {r.text}")
    except Exception as e:
        log.error("Error" + str(e))


