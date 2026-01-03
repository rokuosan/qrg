{
  description = "qrg - QR code cli generator";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      lib = nixpkgs.lib;

      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      forAllSystems = f: lib.genAttrs systems (system:
        let pkgs = import nixpkgs { inherit system; };
        in f { inherit pkgs system; }
      );

      commit = self.rev or "dirty";
      date =
        let d = self.lastModifiedDate or "19700101";
        in "${builtins.substring 0 4 d}-${builtins.substring 4 2 d}-${builtins.substring 6 2 d}T00:00:00Z";
      
      version =
        if self ? rev
        then "unstable-${builtins.substring 0 8 (self.rev)}"
        else "dev";
    in
    {
      packages = forAllSystems ({ pkgs, system }: rec {
        qrg = pkgs.buildGoModule {
          pname = "qrg";
          inherit version;

          src = self;

          vendorHash = "sha256-wzMLu5HV2Ypebjlc+M4G2n54idbPJE0UZN9KxxehCsE=";

          flags = [ "-trimpath" ];

          ldflags = [
            "-s" "-w"
            "-X" "github.com/rokuosan/qrg/cmd.version=${version}"
            "-X" "github.com/rokuosan/qrg/cmd.commit=${commit}"
            "-X" "main.date=${date}"
            "-X" "main.builtBy=nix"
          ];

          env = {
            CGO_ENABLED = "0";
          };

          meta = with lib; {
            description = "QR code CLI generator";
            homepage = "https://github.com/rokuosan/qrg";
            license = licenses.mit;
            mainProgram = "qrg";
            platforms = platforms.all;
          };
        };

        default = qrg;
      });

      apps = forAllSystems ({ system, ... }: {
        default = {
          type = "app";
          program = "${self.packages.${system}.qrg}/bin/qrg";
        };
      });

      devShells = forAllSystems ({ pkgs, ... }: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            gofumpt
          ];
        };
      });

      formatter = forAllSystems ({ pkgs, ... }: pkgs.nixfmt-rfc-style);
    };
}
