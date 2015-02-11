# Add Downloading Tasks through Aria2 RPC

A CLI tool for adding downloading tasks to Aria2 through RPC.  
Useful for **Firefox** extensions like **Flashgot** who can customize download utilities.

![Example](example.png)

Argument Example:

```sh
aria2rpc -cookie 'id=xxx; name=yyy;' \
  -dir /path/to/dest -out filename \
  -rpc http://127.0.0.1:6800/jsonrpc \
  http://example.org/file.zip
```
