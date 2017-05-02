# DM
Tom Whiston <tom.whiston@gmail.com>

## Bootstrapper for Docker for Mac
Dm aims to make developing on docker in OSX a painless experience and to give you lots of tools to get started quickly

## Requirements

`brew install socat`
`brew install libssh2`

## Install

1. Download the binary [(github)](https://github.com/twhiston/dm/releases/) or build from source (requires golang to be installed)
2. `chmod +x dm`
3. Add the binary to your path
4. Change the settings on your docker for mac machine so that only `/tmp` is shared
5. run `dm check` to check your system is set up correctly
6. run `dm init` to initialize your system and create all necessary directories and assets
7. run `dm start` to bring up the dm container stack

## About

Dm tries to make it easier to do docker development on your mac. This means:
- Giving you a useful stack of docker images to get started with
    - mariadb
    - selenium
    - nginx-proxy
    - blackfire
- Giving you a network for connection to them
    - dm_bridge
- Giving you nfs file sharing for better performance
    - Sets up base nfs shares as needed
    - Makes ~/ a docker volume so you can easily bind project files to containers
- Giving you an easy way to use hostnames with your containers
    - add env var VIRTUAL_HOST to your docker compose
    - run `dm hosts add myhostname.dev`
- Lots more
    - run `dm` to see all the base commands
    - run `dm {command_name} -h` to see all available subcommands and help

## Examples

A simple drupal 8 development environment docker-compose.yml could be
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

dm will expose a network called 'dm_bridge' that can be used to connect to container instances

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

Note that you DO NOT need to expose your ports in your docker compose file, as the nginx will detect the port you expose in the container
If the container exposes multiple ports you will need to use the environment variable VIRTUAL_PORT.
For more information about configuring your containers with the proxy please see: https://github.com/jwilder/nginx-proxy

## Assets

To add assets to the project get the go-bindata library
`go get -u github.com/jteeuwen/go-bindata/...`
then run
`go-bindata -o cmd/resources.go assets/...`

This will generate the asset output for the binary. You can then `go install` or `go build` to get your app + assets.

#### Note:
Unfortunately this currently makes a file in the namespace `main` so you will need to manually change this to cmd. This may be changed in future


## Common Issues

### Shared Folders

Currently all the maria data, nfs share data and stack files are stored in /Users/Shared/.dm because of the way that nfs shares work.
Because we dont want to force the site hosting folders to be in a specific we share the whole current user directory.
This means that we need to store the database data (which requires a different set of permissions) outside of the current user dir.
This means that it is not really suitable for use on a shared environment. But by adding some complexity to the init config this could be solved.
If you want to work out an elegant fix for this drop me a line.

### Mariadb

If the mariadb data folder has some issues around permissions and crashes when accessed you should `sudo chmod -R 777 /Users/Shared/.dm/maria/data`
It is also possible that it gets into a crash loop if it crashes a few times. In this case you should delete `/Users/Shared/.dm/maria/data/tc.log` and restart dm

### PHPStorm and container connections

If you are running PHPUnit or similar from phpstorm you need the container to be on the dm_bridge network to access any shared resources.
To do this make sure that the (confusingly named) "network mode" field in the docker container settings is set to `dm_bridge`