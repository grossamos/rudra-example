{
  description = "An example for the rudra tool";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils }: 
    let
      # geberare a version string for each build
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;

      # Magix to make it work on different build systems
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in {
          # Go build flake
          rudra-example = pkgs.buildGoModule {
            pname = "rudra-example";
            inherit version;
            src = ./.;
            vendorSha256 = "bWLGInOuAK0a+2Hj57sdBZfGgjrchzD1WTJZZg9vjtM=";
          };
        }
      );
      devShell = forAllSystems(system: 
        nixpkgsFor.${system}.mkShell {
          buildInputs = [ 
            nixpkgsFor.${system}.go 
            nixpkgsFor.${system}.zlib 
            nixpkgsFor.${system}.wget 
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
      );
      defaultPackage = forAllSystems (system: self.packages.${system}.rudra-example);
    };
}
