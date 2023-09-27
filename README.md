# module-core-rql

port from https://github.com/NubeIO/module-core-rql


```bash
go build -o module-core-rql
```
to build and run rubix-os you can use the bash script
```bash
bash build.bash <YOUR_ROS_PATH>
```
example
```bash
bash build.bash code/go
```


## examples

### ping all hosts and add an alert
```js
let groups = RQL.GetAllHostsStatus();

let hostList = [];
for (let group of groups) {
  for (let host of group.Hosts) {
    let newHost = {};
    newHost["name"] = host.Name;
    newHost["uuid"] = host.UUID;
    newHost["count"] = host.PingFailCount;
    hostList.push(newHost);
  }
}

let alertCount = 0;

for (let host of hostList) {
  if (host.count >= 5) {
    let body = {
      hostUUID: host.uuid,
      entityType: "device",
      type: "ping",
      status: "active",
      severity: "crucial",
      body: `host-name: ${host.name} fail count: ${host.count}`,
    };
    RQL.AddAlert(host.uuid, body);
    alertCount++;
  }
}

RQL.Return = alertCount;
```