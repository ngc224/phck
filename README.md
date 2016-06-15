# PHCK
PHCK is process health checker web server.

## Install & Listen web server
```
$ go get github.com/ngc224/phck
$ phck phck.conf
```
or
```
$ sudo yum install https://github.com/ngc224/phck/releases/download/v0.2/phck-0.2-1.el7.centos.x86_64.rpm
$ sudo systemctl start phck.service
```


## Usage

#### check process & result success
HTTP status code 200

```
$ curl -i http://127.0.0.1:8939/
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 15 Jun 2016 08:16:41 GMT

{
  "datetime": "2016-06-15 16:38:27",
  "hostname": "localhost",
  "status": true,
  "process": [
    {
      "label": "nginx",
      "pidfile": "/var/run/nginx.pid",
      "running": true,
      "detail": {
        "name": "nginx",
        "pid": 10709,
        "ppid": 1,
        "command": "nginx: master process /usr/sbin/nginx -c /etc/nginx/nginx.conf",
        "stat": "S",
        "thread": 1,
        "use_memory": 0.10912581
      }
    },
    {
      "label": "td-agent",
      "pidfile": "/var/run/td-agent/td-agent.pid",
      "running": true,
      "detail": {
        "name": "ruby",
        "pid": 22499,
        "ppid": 1,
        "command": "/opt/td-agent/embedded/bin/ruby /usr/sbin/td-agent --log /var/log/td-agent/td-agent.log --use-v1-config --group td-agent --daemon /var/run/td-agent/td-agent.pid",
        "stat": "S",
        "thread": 2,
        "use_memory": 1.3995278
      }
    }
  ]
}
```

#### check process & result error
HTTP status code 500

```
$ curl -i http://127.0.0.1:8939/
HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Wed, 15 Jun 2016 08:15:22 GMT

{
  "datetime": "2016-06-15 16:38:27",
  "hostname": "localhost",
  "status": false,
  "process": [
    {
      "label": "nginx",
      "pidfile": "/var/run/nginx.pid",
      "running": true,
      "detail": {
        "name": "nginx",
        "pid": 10709,
        "ppid": 1,
        "command": "nginx: master process /usr/sbin/nginx -c /etc/nginx/nginx.conf",
        "stat": "S",
        "thread": 1,
        "use_memory": 0.10912581
      }
    },
    {
      "label": "td-agent",
      "pidfile": "/var/run/td-agent/td-agent.pid",
      "running": false,
      "message": "PID file not open",
      "detail": {}
    }
  ]
}
```

#### CLI mode
```
$ phck --cli phck.conf
```

## Help
```
Usage:
  phck [options] CONFIGFILE

Application Options:
  -c, --cli      CLI mode
      --pidfile= Set PIDFILE
  -h, --help     Show this help message
  -v, --version  Show this build version
```