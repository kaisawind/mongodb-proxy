<h1 align="center">
  <br>
  <img src="https://i.imgur.com/Xz0DUXf.png" alt="Olivia's character" width="400">
  <br>
</h1>

<h4 align="center">💁‍♀️ Your new best friend built with an artificial neural network</h4>
<h5 align="center">Inspired by <a href="https://github.com/leon-ai/leon">leon-ai/leon</a> :)</h5>

<p align="center">
  <a href="https://travis-ci.org/olivia-ai/olivia"><img src="https://travis-ci.org/olivia-ai/olivia.svg?branch=master"></a>
  <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_shield"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=shield"></a>
</p>

<p align="center">
  <a href="#getting-started">Getting started</a> •
  <a href="https://docs.olivia-ai.org">Documentation</a> •
  <a href="https://github.com/orgs/olivia-ai/projects">Projects</a> •
  <a href="https://www.youtube.com/watch?v=JmJZi9gmKvI">Video</a> •
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
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_large)

