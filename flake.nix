{
  description = "Get your git";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }:
    {
      overlays.default = final: prev: {
        inherit (self.packages.${prev.system}) workon;
      };
    }
    // flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.gotem = pkgs.buildGoModule {
          pname = "gotem";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-zxq2/zQXXGFyS/oeZ18E2359f/v3GV9cZCE0HboOfsM=";
        };
        packages.default = self.packages.${system}.gotem;

        devShells.default = pkgs.mkShell {
          name = "gotem";
          buildInputs = with pkgs; [
            cobra-cli
            go
            gopls
          ];
        };
      }
    );
}
