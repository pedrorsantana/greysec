<h1 align="center">Welcome to GreySec 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

<div style="text-align:center"><img alt="Alien" src="https://cdn.bulbagarden.net/upload/f/fd/605Elgyem.png"/></div>

> Greysec for packets that pass through the network interface and performs queries on a large internet base (greynoise.io) responsible for mapping robots that perform cyberattack and enumeration robots. When the communication with one of these robots is identified, it is automatically banned through iptables protecting your machine from being enumerated or invaded.

> With greysec you will be protected from worms that try to carry out the most varied attacks (Ransoware, Enumeration, Botminer, Botnet, etc.)

## Requirements

>To block access we use iptables
```sh
#  sudo apt-get install -y iptables
``` 


>Docker needs iptable installed and cap run properties defined (--cap-add=NET_ADMIN --cap-add=NET_RAW)
Example:
```sh
#  docker run --cap-add=NET_ADMIN --cap-add=NET_RAW -it ubuntu
#  root(Inside Docker): apt-get update;apt-get install -y iptables;    
``` 
## Start using

```sh
# curl -LO https://github.com/pedrorsantana/greysec/releases/download/v1.0/greysec; sudo chmod +x greysec; sudo mv greysec /bin/greysec;
# greysec --help
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