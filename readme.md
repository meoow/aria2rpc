# Add Downloading Tasks through Aria2 RPC


A CLI tool for adding downloading tasks to Aria2 through RPC.  
Useful for **Firefox** extensions like **Flashgot** who can customize download utilities.

Implemented in both **Go** and **Python**, choose which you like.
Note that aria2 comes with its own script(at doc/xmlrpc/aria2rpc under the source file tree) which implemented the full XML-RPC APIs in python, you might check on that if you need more than just adding download links.

![Example](example.png)

Argument Example:

```sh
aria2rpc -cookie 'id=xxx; name=yyy;' \
  -dir /path/to/dest -out filename \
  -rpc http://127.0.0.1:6800/jsonrpc \
  http://example.org/file.zip
```
