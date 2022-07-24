# Rudra example

## Overview
This repository contains a sample pipeline to showcase how rudra can be used in practice.
The configuration file for this pipeline can be found under ``./.github/workflows/test.yml``.

Openapi files are automatically generated using a tool called `swag` (see [here](https://github.com/swaggo/swag) for more information).
These files can be found under the `./docs` directory or under `http://localhost:8080/openapi/index.html` when starting the service.

## local setup
### On GNU/Linux systems or MacOS
Install nix onto your system
```bash
curl -L https://nixos.org/nix/install | sh
```

Enable flakes by adding the following line to `~/.config/nix/nix.conf` or `/etc/nix/nix.conf` (preferably the first).
```
experimental-features = nix-command flakes
```

### On NixOS 
Enable flakes by adding the following options to your `nix.conf`
```nix
{ pkgs, ... }: {
  nix = {
    package = pkgs.nixFlakes;
    extraOptions = ''
      experimental-features = nix-command flakes
    '';
   };
}
```

## Running Rudra example
To start rudra-example localy simply run:
```bash
nix run 
```

To start rudra-example in a docker container run:
```bash
docker build -t rudra-example .

docker run -d -rm --name app rudra-example
```

## Maintenance
### Updating openapi files
This example uses `swag` to automatically generate its openapi files.
To update these run:
```bash
nix develop
./swag i 
```
