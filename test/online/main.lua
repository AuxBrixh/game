local enet = require("enet")
local host = enet.host_create()
host:connect("localhost:8888")
local event = host:service(100)
while event do
  if event.type == "connect" then
    print(event.peer)
    event.peer:send("DATA")
  end
  event = host:service()
end