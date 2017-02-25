# DM
Tom Whiston <tom.whiston@gmail.com>

## Bootstrapper for Docker for Mac

## Requirements

`brew install socat`
`brew install libssh2`

## Install

1. Download the binary [(github)](https://github.com/twhiston/pxd/releases/) or build from source (requires golang to be installed)
2. `chmod +x pxd`
2. Add the binary to your path
3. Change the settings on your docker for mac machine so that only `/tmp` is shared
3. you `pxd` yourself

## About

- Creates a local dir for data
    - all nfs, mariadb and other settings and data are kept in one place
- Sets up NFS shares on your docker for mac machine
    - Note that you must remove all shares from the docker for mac interface other than /tmp
- Starts a mariadb container that can be linked in your other applications
    - uses a path in our local data dir, so that data persists between container builds
- Starts a selenium-chrome container for behat testing
- Sets up an xdebug loopback and enables socat for phpstorm docker integration

For example a simple drupal 8 development environment docker-compose.yml could be
```
app:
  # This is based on our generic drupal s2i image, which is used for Openshift builds
  image: pwcsexperiencecenter/drupal-7.0-local
  volumes:
    # Because dockers mounts our /Users folder this is reading directly from our local machine
    - ./:/opt/app-root/src
  external_links
    # System wide mariadb container
    - mariadb_local
  ports:
      - 8000:8000
  environment:
        # Needed for xdebug
      - PHP_IDE_CONFIG="serverName=localhost"
```

## Networks

pxd will expose a network called 'dockerpwc_bridge' that can be used to connect to the mariadb instance

for example

```
version: "2"

services:
  app:
  # Local development image,
  # This is based on our generic drupal s2i image, which is used for Openshift builds
    image: pwcsexperiencecenter/drupal-base:local
    container_name: my-container
    volumes:
    # link the whole project into the image at the appropriate point.
    # This allows nginx in the container to serve the correct files
    # Because docker mounts our /Users folder this is reading directly from our local machine
      - ./:/opt/app-root/src
    ports:
      - 8000:8000
    networks:
      - dockerpwc_bridge
    environment:
        # Needed for xdebug
      - PHP_IDE_CONFIG="serverName=dev"
        # Needed for command line utils, e.g. clear
      - TERM=xterm

networks:
  dockerpwc_bridge:
    external: true
 ```

## Upgrade

If upgrading from an earlier version of pxd then you must take the following steps

1. delete ~/.docker-pwc (this will delete all your databases so back up if necessary)
2. delete your mariadb containers, use `docker ps -a` to see these
3. delete any site containers that used the mariadb link
4. edit your /etc/exports and remove all directories added by pxd
5. `sudo pxd -start`

If you have any problems at this stage try deleting your mariadb image and letting pxd download it again on start.

## Usage

`dm -start` starts the environment

`dm -stop`  stops the environment

#### Flags

`-clean` remove ALL existing data from the data folder

`-data-dir` set a data folder for your docker-pwc setup. Defaults to ~/.docker-pwc

## Assets

To add assets to the project get the go-bindata library
`go get -u github.com/jteeuwen/go-bindata/...`
then run
`go-bindata -o cmd/resources.go assets/...`
this will generate the asset output for the binary. You can then `go install` or `go build` to get your app + assets



