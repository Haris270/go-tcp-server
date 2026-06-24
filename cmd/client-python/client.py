import asyncio


async def send_message(writer: asyncio.StreamWriter):

    while True:
        # offload getting user input to a background thread so event loop is not blocked (e.g. when we use input())
        message = await asyncio.to_thread(input, ">> ")
        writer.write(f"{message}\n".encode())
        await writer.drain()


async def receive_message(reader: asyncio.StreamReader):
    while True:
        
        recv_message = await reader.readline()
        
        if not recv_message:
            print("User Disconnected")
            break 
        
        print(recv_message.decode().strip())
    

async def main():
    reader, writer = await asyncio.open_connection("localhost", 8080)
    async with asyncio.TaskGroup() as tg:
        task1 = tg.create_task(receive_message(reader))
        task2 = tg.create_task(send_message(writer))

    


if __name__ == "__main__":
    asyncio.run(main())


# Why use asyncio & asyncio.create_connection instead of sockets and socket.socket.connect