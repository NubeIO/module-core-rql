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
let lowVolatge = 4.9;

// GET ALL THE POINTS
let allPoints = RQL.GetPointsAllHosts();
let pointsWithLowVoltage = [];
for (let i = 0; i < allPoints.length; i++) {
  // get all points
  let points = allPoints[i].Points;
  let host = allPoints[i]; // host
  RQL.Log(points.HostName);
  for (let i = 0; i < points.length; i++) {
    // get all points where name is voltage
    let point = points[i];
    if (point.Name == "voltage") {
      text = point;
      let newPoint = {
        name: point.Name,
        uuid: point.UUID,
        value: parseFloat(point.PresentValue),
        hostName: host.HostName,
        hostUUID: host.HostUUID,
      };
      pointsWithLowVoltage.push(newPoint);
    }
  }
}

// WORK OUT WHICH HAVE LOW VOLTAGE
for (let i = 0; i < pointsWithLowVoltage.length; i++) {
  let point = pointsWithLowVoltage[i];

  if (point.value < lowVolatge) {
  } else {
    // remove points from the array
    pointsWithLowVoltage.splice(i, 1);
  }
}

let out = [];
let count = 0;

// RAISE THE ALARM!!!
for (let i = 0; i < pointsWithLowVoltage.length; i++) {
  count++;
  let point = pointsWithLowVoltage[i];
  let hostUUID = point.hostUUID;
  let body = {
    hostUUID: hostUUID,
    entityUUID: point.uuid,
    entityType: "device",
    type: "threshold",
    status: "active",
    severity: "crucial",
    body: `low voltage: ${point.value}`,
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