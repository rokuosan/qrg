{
  description = "qrg - QR code cli generator (qrg <text>)";

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
        let
          pkgs = import nixpkgs { inherit system; };
        in
          f { inherit system pkgs; }
      );

      version = "0.3.1";
    in
    {
      packages = forAllSystems ({ pkgs, ... }: rec {
        qrg = pkgs.buildGoModule {
          pname = "qrg";
          inherit version;
          src = self;

          # nix build すると正しいハッシュがエラーに出るので貼り替える
          vendorHash = lib.fakeHash;

          ldflags = [ "-s" "-w" ];

          meta = with lib; {
            description = "qrcode cli generator; run \"qrg <message>\" on your terminal";
            homepage = "https://github.com/rokuosan/qrg";
            license = licenses.mit;
            mainProgram = "qrg";
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

      checks = forAllSystems ({ system, ... }: {
        build = self.packages.${system}.qrg;
      });

      formatter = forAllSystems ({ pkgs, ... }: pkgs.nixfmt-rfc-style);

      # overlayとしても使えるようにする
      overlays.default = final: prev: {
        qrg = self.packages.${prev.system}.qrg;
      };
    };
}

