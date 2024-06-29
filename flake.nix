{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    {
      overlays.default = final: prev: {
        inherit (self.packages.${prev.system}) workon;
      };
    } // flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.gotem = pkgs.buildGoModule {
          pname = "gotem";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-TUebsXaMc0d3m5cCWuVS1/QZMGN5yjLGeqmRrkHxtKE=";
        };
        packages.default = self.packages.${system}.gotem;

        devShells.default = pkgs.mkShell {
          name = "gotem";
          buildInputs = with pkgs; [ cobra-cli go gopls ];
        };
      }
    );
}
