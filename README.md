<h1 align="center">Welcome to GreySec 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

<div style="text-align:center"><img alt="Alien" src="https://cdn.bulbagarden.net/upload/f/fd/605Elgyem.png"/></div>
> Greysec is a firewall based application that consume greynoise.io api.

## Running in docker

For use in docker install iptable and define --cap-add=NET_ADMIN --cap-add=NET_RAW
Example:
```sh
#  docker run --cap-add=NET_ADMIN --cap-add=NET_RAW -it ubuntu
```
## Compiling

```sh
# make
```

## Usage

```sh
# greysec --key [YOUR_KEY] --interface [YOUR_OPTIONAL_INTERFACE]
```

## Author

👤 **Pedro Santana**

* Github: [@pedrorsantana](https://github.com/pedrorsantana)
* LinkedIn: [@https:\/\/www.linkedin.com\/in\/pedro-ricardo-ramos-f-de-santana-a38444134\/](https://linkedin.com/in/https:\/\/www.linkedin.com\/in\/pedro-ricardo-ramos-f-de-santana-a38444134\/)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_