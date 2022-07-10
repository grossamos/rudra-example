{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/660ac43ff9ab1f12e28bfb31d4719795777fe152.tar.gz") {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.which
    pkgs.htop
    pkgs.zlib
    pkgs.niv
    pkgs.go
    pkgs.wget
    pkgs.postman
  ];
  shellHook = ''
    if [ ! -f "./swag" ]; then
        mkdir .tmp
        cd .tmp
        wget https://github.com/swaggo/swag/releases/download/v1.8.3/swag_1.8.3_Linux_x86_64.tar.gz 
        tar -xvf swag_*_Linux*.tar.gz
        cp swag ../
        cd ..
        rm -rf .tmp
        chmod +x swag
    fi
    alias swag="$PWD/swag"
  '';
}


