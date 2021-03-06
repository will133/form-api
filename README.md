> # 🗃 Forma [![Tweet][icon_twitter]][twitter_publish] <img align="right" width="126" src=".github/character.png">
> [![Analytics][analytics_pixel]][page_promo]
> Data Collector as a Service &mdash; your personal server for HTML forms.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![Build Status][icon_build]][page_build]
[![Code Quality][icon_quality]][page_quality]
[![Code Coverage][icon_coverage]][page_quality]
[![Research][icon_research]](../../tree/research)
[![License][icon_license]](LICENSE)

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**May 31, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [x] v2: [Accounts and CLI CRUD][project_v2]
  - [**August 31, 2018**][project_v2_dl]
  - Command line interface for create, read, update and delete operations above gRPC.
- [ ] v3: [Templating and RESTful API][project_v3]
  - [**September 30, 2018**][project_v3_dl]
  - Template system and Edge Side Includes/Server Side Includes support.
  - Integrate gRPC gateway.
  - Improve gRPC layer.
- [ ] v4: [DSL for validation, CSI, and GUI][project_v4]
  - [**October 31, 2018**][project_v4_dl]
  - Domain-specific language to define validation rules.
  - Client-side integration.
  - Graphical user interface and admin panel to perform create, read, update and delete operations.
- [ ] Forma, SaaS
  - **December 31, 2018**
  - Ready to apply on Cloud.
  - Move to [OctoLab](https://github.com/octolab/) organization and rename project to `forma`.

## Motivation

- We need better integration with static sites built with [Hugo](https://gohugo.io/).
- We want cheaper products than [Formata](https://www.formata.io/) or [FormKeep](https://formkeep.com/).
- We have to full control over our users' data and protect it from third parties.

## Quick start

Requirements:

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

       Name                     Command               State                                  Ports
------------------------------------------------------------------------------------------------------------------------
form-api_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
form-api_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
form-api_service_1   form-api run --with-profil ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091
```

<details>
<summary><strong>GET curl /api/v1/UUID</strong></summary>

```bash
$ curl http://localhost:8080/api/v1/10000000-2000-4000-8000-160000000004
# <form id="10000000-2000-4000-8000-160000000004" lang="en" title="Email subscription"
#       action="http://localhost/api/v1/10000000-2000-4000-8000-160000000004" method="post"
#       enctype="application/x-www-form-urlencoded">
#       <input id="10000000-2000-4000-8000-160000000004_email" name="email" type="email" title="Email"
#              maxlength="64" required="true"></input>
# </form>
```
</details>

<details>
<summary><strong>POST /api/v1/UUID</strong></summary>

```bash
$ curl -v -H "Content-Type: application/x-www-form-urlencoded" \
       --data-urlencode "email=test@my.email" \
       http://localhost:8080/api/v1/10000000-2000-4000-8000-160000000004
# > POST /api/v1/10000000-2000-4000-8000-160000000004 HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/7.54.0
# > Accept: */*
# > Content-Type: application/x-www-form-urlencoded
# > Content-Length: 21
# >
# < HTTP/1.1 302 Found
# < Location: https://kamil.samigullin.info/#eyJpZCI6IjEwMDAwMDAwLTIwMDAtNDAwMC04MDAwLTE2MDAwMDAwMDAwNCIsInJlc3VsdCI6InN1Y2Nlc3MifQ==
# < Date: Sat, 05 May 2018 09:34:47 GMT
# < Content-Length: 0
# <
```
</details>

## Specification

### API

You can find API specification [here](env/client/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest/)
HTTP client to work with the API - you can import data for it from the [file](env/client/insomnia.json).
Or you can choose [Postman](https://www.getpostman.com/) - its import data is [here](env/client/postman.json) and
[here](env/client/postman.env.json).

### CLI

You can use CLI not only to start the HTTP server but also to execute
[CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) operations.

<details>
<summary><strong>CLI interface</strong></summary>

```bash
$ form-api --help
Forma

Usage:
  form-api [command]

Available Commands:
  completion  Print Bash or Zsh completion
  ctl         Communicate with Forma server via gRPC
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for form-api

Use "form-api [command] --help" for more information about a command.
```
</details>

#### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ form-api completion -f bash > /path/to/bash_completion.d/form-api.sh
$ form-api completion -f zsh  > /path/to/zsh-completions/_form-api.zsh
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/form-api
```

### Binary

```bash
$ export VER=2.0.0      # all available versions are on https://github.com/kamilsk/form-api/releases/
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ wget -q -O form-api.tar.gz \
       https://github.com/kamilsk/form-api/releases/download/"${VER}/form-api_${VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf form-api.tar.gz -C "${GOPATH}"/bin/ && rm form-api.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/form-api:2.x
```

### From source code

```bash
$ egg github.com/kamilsk/form-api@^2.0.0 -- make test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/form-api@^2.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

## Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

<sup id="egg">1</sup> The project is still in prototyping. [↩](#anchor-egg)

---

[![Gitter][icon_gitter]](https://gitter.im/kamilsk/form-api)
[![@kamilsk][icon_tw_author]](https://twitter.com/ikamilsk)
[![@octolab][icon_tw_sponsor]](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)

[analytics_pixel]: https://ga-beacon.appspot.com/UA-109817251-15/form-api/readme?pixel

[icon_build]:      https://travis-ci.org/kamilsk/form-api.svg?branch=master
[icon_coverage]:   https://scrutinizer-ci.com/g/kamilsk/form-api/badges/coverage.png?b=master
[icon_gitter]:     https://badges.gitter.im/Join%20Chat.svg
[icon_license]:    https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]:    https://img.shields.io/badge/patreon-donate-orange.svg
[icon_quality]:    https://scrutinizer-ci.com/g/kamilsk/form-api/badges/quality-score.png?b=master
[icon_research]:   https://img.shields.io/badge/research-in%20progress-yellow.svg
[icon_tw_author]:  https://img.shields.io/badge/author-%40kamilsk-blue.svg
[icon_tw_sponsor]: https://img.shields.io/badge/sponsor-%40octolab-blue.svg
[icon_twitter]:    https://img.shields.io/twitter/url/http/shields.io.svg?style=social

[page_build]:      https://travis-ci.org/kamilsk/form-api
[page_promo]:      https://kamilsk.github.io/form-api/
[page_quality]:    https://scrutinizer-ci.com/g/kamilsk/form-api/?branch=master

[project_v1]:      https://github.com/kamilsk/form-api/projects/1
[project_v1_dl]:   https://github.com/kamilsk/form-api/milestone/1
[project_v2]:      https://github.com/kamilsk/form-api/projects/2
[project_v2_dl]:   https://github.com/kamilsk/form-api/milestone/2
[project_v3]:      https://github.com/kamilsk/form-api/projects/3
[project_v3_dl]:   https://github.com/kamilsk/form-api/milestone/3
[project_v4]:      https://github.com/kamilsk/form-api/projects/4
[project_v4_dl]:   https://github.com/kamilsk/form-api/milestone/4

[twitter_publish]: https://twitter.com/intent/tweet?text=Data%20Collector%20as%20a%20Service&url=https://kamilsk.github.io/form-api/&via=ikamilsk&hashtags=go,service,data-collector,form-handler
