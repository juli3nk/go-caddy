# Go Caddy

## Installation

```sh
# go get github.com/juli3nk/go-caddy/cmd/caddybuild
```

## Usage

### Plugin
To list plugins

```sh
# caddybuild plugin ls
NAME                  TYPE
http.upload           Directive
tls.dns.cloudflare    DNS Provider
http.awses            Directive
http.mailout          Directive
...
```

To get information about one plugin

```sh
# caddybuild plugin info http.git
Name: http.git
Type: Directive
Description: The git plugin makes it possible to deploy your site with a simple git push.

The git directive starts a service routine that runs during the lifetime of the server. When the service starts, it clones the repository. While the server is still up, it pulls the latest every so often. You can also set up a webhook to pull immediately after a push. In regular git fashion, a pull only includes changes, so it is very efficient.
Website: https://github.com/abiosoft/caddy-git
Docs: https://github.com/abiosoft/caddy-git/blob/master/README.md
```

### Build
To build a custom binary including plugins follow the steps below or check out the the directory `example`.

1. Create a `.caddybuild.yml` config file

```yaml
---
plugins:
  - http.git
telemetry: false
```

2. Build

```sh
# caddybuild build
```

3. Test newly built binary

```sh
# /tmp/caddy/caddy -version
Caddy v1.0.0 (h1:KI6RPGih2GFzWRPG8s9clKK28Ns4ZlVMKR/v7mxq6+c=)
```
