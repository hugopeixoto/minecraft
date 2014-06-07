# Flayer

Flayer is a Minecraft skin cache server, indexed by the player's UUID.


## Installation

Make sure you have `go` installed before running these commands.

```
export GOPATH=/some/path
go get github.com/hugopeixoto/minecraft/flayer
/some/path/bin/flayer -cachedir /var/www/assets -listen your-server.com:8080
```


## API

### Full skins
Retrieving the full skin:
```
<img src='http://your-server.com:8080/skins/61208bec81194e228b4f510cd9aa6fe0/full.png'/>
```

### Heads with hat/helm/thing-that-goes-on-the-head
You can retrieve the skin's head, in several resolutions:
```
<img src='http://your-server.com:8080/skins/61208bec81194e228b4f510cd9aa6fe0/head-32.png'/>
<img src='http://your-server.com:8080/skins/61208bec81194e228b4f510cd9aa6fe0/head-48.png'/>
<img src='http://your-server.com:8080/skins/61208bec81194e228b4f510cd9aa6fe0/head-64.png'/>
```

And that's it! If the given profile doesn't exist, it defaults to steve
(creating a symlink for each uuid). If something else goes wrong, flayer will
respond with `404 Not found`.


## Webserver integration

You can point nginx/apache to the cachedir, to spare flayer from unnecessary hits.
Use `try_files` or `RewriteCond/Rule` to do this.


## Missing features

- Currently, the available resolutions are hardcoded. Make this dynamic/configurable.
- Add a cache sweeper that checks if there are updates for each uuid.
- Add authentication/authorization, so that we only generate skins for authorized callers.
- Place all configuration settings in a json file.


## Inspiration

This was definitely inspired by [minotar](https://github.com/minotar). The main
difference (apart from the lack of features of Flayer) is that Flayer caches
the heads.
