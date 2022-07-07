{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/660ac43ff9ab1f12e28bfb31d4719795777fe152.tar.gz") {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.which
    pkgs.htop
    pkgs.zlib
    pkgs.niv
    pkgs.go
    pkgs.postman
  ];
  allowUnfree = true;
}


