# Superman Detector
Detect suspicious consecutive logins that would require unrealistic speeds to get to subsequent locations

## Build
The project uses `go mods`, so clone the repo to a folder of your choice then run `go build` in the project root.

Of course go must be previously installed

## Runing the Server
- Bare metal `./swchallenge`
- Dockerized:

```
docker build . -t swchal && docker run -d -p 8100:8100 swchal:latest
# if 8100 doesn't work on your system try with the `-P` instead of `-p 8100:8100` and see what the left side port is in `docker ps`

# To see the logs
docker ps # to get the container id and host port (left side of the '->')
# If host port is different than 8100, change in the curl commands below also

docker logs -f <the_container_id> # use the container id from above
```

## Try out the detector
POST at least of couple login attempts:  

```bash
curl -X POST -d '{"username": "bob", "unix_timestamp": 1514766000, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e43","ip_address": "24.242.71.20"}' http://localhost:8100/v1
curl -X POST -d '{"username": "bob", "unix_timestamp": 1514664800, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e41","ip_address": "206.90.252.6"}' http://localhost:8100/v1
curl -X POST -d '{"username": "bob", "unix_timestamp": 1514764800, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e42","ip_address": "206.81.252.6"}' http://localhost:8100/v1
```

## External packages used
- github.com/mattn/go-sqlite3 v1.10.0 - SQLite for storage of login attempts
- github.com/oschwald/maxminddb-golang v1.3.1 - Maxmind IP to Geolocation (City) database for lookup of geolocation from IP
- github.com/rohanthewiz/serr v0.4.2 - A structured error wrapping package
- github.com/sirupsen/logrus v1.4.2 - A structured logging package

## Credits
The idea for this project came from SecureWorks. Thanks for the opportunity!
