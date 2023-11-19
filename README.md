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


### droplet alarms
```js
//  GET ALL POINTS BY HOST
let pointsByHost = [];
let pointsByHostByModule = RQL.GetPointsByModuleAllHosts("lora");

// GET ALL POINTS NAMED VOLTAGE AND BELOW VOLTAGE LIMIT
let pointsWithLowVoltage = [];
let lowVolatgeLimit = 4.9;

for (let i = 0; i < pointsByHostByModule.length; i++) {
    let host = pointsByHostByModule[i];
    let hostUUID = host.HostUUID;
    let hostName = host.HostName;
    let points = host.Points;

    for (let i = 0; i < points.length; i++) {
        let point = points[i];
        let pointUUID = point.UUID;
        let pointName = point.Name;
        let sensorID = point.AddressUUID;
        let pv = parseFloat(point.PresentValue);

        if (pointName == "voltage") {
            if (pv < lowVolatgeLimit) {
                pointsWithLowVoltage.push({
                    hostUUID: hostUUID,
                    hostName: hostName,
                    pointUUID: pointUUID,
                    pointName: pointName,
                    presentValue: pv,
                    sensorID: sensorID,
                });
            }
        }
    }
}

let out = [];
let count = 0;

// RAISE THE ALARM!!!
for (let i = 0; i < pointsWithLowVoltage.length; i++) {
    count++;
    let point = pointsWithLowVoltage[i];
    let hostUUID = point.hostUUID;
    let hostName = point.hostName;
    let body = {
        hostUUID: hostUUID,
        entityUUID: point.uuid,
        entityType: "device",
        type: "threshold",
        status: "active",
        severity: "crucial",
        title: "lora sensor low volatge",
        body: `host: ${hostName} sensor-id: ${point.sensorID} value of voltage: ${point.presentValue}`,
        tags: [
            {
                tag: "volatge",
            },
            {
                tag: "droplet",
            },
        ],
    };

    let alert = RQL.AddAlert(hostUUID, body);
    out.push({ count: count, alert: alert });
}

RQL.Result = out;
```

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

let out = "no alerts";
let alertCount = 0;
let offlineCount = 1;

for (let host of hostList) {
    if (host.count >= offlineCount) {
        let body = {
            hostUUID: host.uuid,
            entityType: "device",
            type: "ping",
            status: "active",
            severity: "crucial",
            body: `host-name: ${host.name} fail count: ${host.count}`,
            tags: [
                {
                    tag: "rubix-compute",
                },
                {
                    tag: "ping",
                },
            ],
        };
        alertCount++;
        RQL.AddAlert(host.uuid, body);
        out = `alert count ${alertCount} host-count: ${hostList.length} fail-ping count: ${host.count}`;
    } else {
        out = `alert count ${alertCount} host-count: ${hostList.length}`;
    }
}

RQL.Result = out;

```