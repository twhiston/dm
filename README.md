# DM
Tom Whiston <tom.whiston@gmail.com>

## Bootstrapper for Docker for Mac

## Requirements

`brew install socat`
`brew install libssh2`

## Install

1. Download the binary [(github)](https://github.com/twhiston/dm/releases/) or build from source (requires golang to be installed)
2. `chmod +x dm`
3. Add the binary to your path
4. Change the settings on your docker for mac machine so that only `/tmp` is shared

## About

- Creates a local dir for data
    - all nfs, mariadb and other settings and data are kept in one place
- Sets up NFS shares on your docker for mac machine
    - Note that you must remove all shares from the docker for mac interface other than /tmp
- Starts a mariadb container that can be linked in your other applications
    - uses a path in our local data dir, so that data persists between container builds
- Starts a selenium-chrome container for behat testing
- Sets up an xdebug loopback and enables socat for phpstorm docker integration
- Starts an nginx proxy that allows you to have hostnames for your docker sites

For example a simple drupal 8 development environment docker-compose.yml could be
```
app:
  # This is based on our generic drupal s2i image, which is used for Openshift builds
  image: openshift/php-56-centos7
  volumes:
    - ./:/opt/app-root/src
  external_links
    - mariadb_local
  environment:
        # Needed for xdebug
      - PHP_IDE_CONFIG="serverName=dev"
      - VIRTUAL_HOST=test.dev
```

## Networks

pxd will expose a network called 'dm_bridge' that can be used to connect to the mariadb instance

for example

```
version: "2"

services:
  app:
  # Local development image,
  # This is based on our generic drupal s2i image, which is used for Openshift builds
    image: openshift/php-56-centos7
    container_name: my-container
    volumes:
    # link the whole project into the image at the appropriate point.
    # This allows nginx in the container to serve the correct files.
    # ./ must be under the current users home folder
      - ./:/opt/app-root/src
    ports:
      - 8000:8000
    networks:
      - dm_bridge
    environment:
        # Needed for xdebug
      - PHP_IDE_CONFIG="serverName=dev"
        # Needed for command line utils, e.g. clear
      - TERM=xterm
      - VIRTUAL_HOST=test.dev

networks:
  dm_bridge:
    external: true
 ```


## Assets

To add assets to the project get the go-bindata library
`go get -u github.com/jteeuwen/go-bindata/...`
then run
`go-bindata -o cmd/resources.go assets/...`
this will generate the asset output for the binary. You can then `go install` or `go build` to get your app + assets



