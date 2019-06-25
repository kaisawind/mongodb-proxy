<h1 align="center">
  <br>
  <img src="https://repository-images.githubusercontent.com/167809029/f477a980-86bb-11e9-9487-8c3627dee825" alt="Kaisawind's character" width="400">
  <br>
</h1>

<p align="center">
  <a href="https://github.com/kaisawind/mongodb-proxy/releases">
    <img src="https://img.shields.io/github/release/kaisawind/mongodb-proxy.svg" alt="GitHub release">
  </a>
  <a href="https://github.com/golang">
    <img src="https://img.shields.io/badge/golang-1.12.5-green.svg">
  </a>
  <a href="https://travis-ci.com/kaisawind/mongodb-proxy">
    <img src="https://travis-ci.com/kaisawind/mongodb-proxy.svg?token=zAYkhFNqwBwmfWpeEt2s&branch=master">
  </a>
  <a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fkaisawind%2Fmongodb-proxy?ref=badge_shield">
    <img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkaisawind%2Fmongodb-proxy.svg?type=shield">
  </a>
</p>

<p align="center">
  <a href="#getting-started">Getting started</a> •
  <a href="https://github.com/grafana/grafana/blob/master/docs/sources/plugins/developing/datasources.md">Documentation</a> •
  <a href="#contributors">Contributors</a> •
  <a href="#license">License</a>
</p>

## Getting started
### Installation
#### Install from code
Clone mongo-proxy from the master branch of Github repository

```bash
git clone git@github.com:kaisawind/mongodb-proxy.git
```

Then build the project

```bash
make build
```

And run the application

```bash
./bin/cmd-server
```

The REST Api is now listening on the port `8080`, to change it just set it inside the environment variable `PORT`, or add flag `--port`.

#### Install from docker
Clone mongo-proxy from the master branch of Github repository

```bash
git clone git@github.com:kaisawind/mongodb-proxy.git
```

Then build the project and build the image

```bash
make docker
```

And run the image

```bash
docker run --rm -p 8080:8080 kaisawind/mongodbproxy
```

#### Install from helm
Clone mongo-proxy from the master branch of Github repository

```bash
git clone git@github.com:kaisawind/mongodb-proxy.git
```

Then build the project and build the image

```bash
make docker
```

And deploy it by helm

```bash
helm install .
```

## How it works

It just convert mongodb datasource to simple json datasource, then use the plugin to insert data to grafana.

[grafana关联mongodb数据库](https://www.kaisawind.com/mongodb/2019/03/28/grafana.html)

[grafana-simple-json-datasource](https://grafana.com/plugins/grafana-simple-json-datasource)

## Contributors
- [kaisawind](https://github.com/kaisawind) - creator, maintainer

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkaisawind%2Fmongodb-proxy.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkaisawind%2Fmongodb-proxy?ref=badge_large)

