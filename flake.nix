{
  description = "/x/perimental code";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils";

    rust-overlay = {
      url = "github:oxalica/rust-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "utils";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
  };

  outputs = { self, nixpkgs, utils, gomod2nix, rust-overlay }@attrs:
    utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ] (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            (final: prev: {
              go = prev.go_1_19;
              buildGoModule = prev.buildGo118Module;
            })
            gomod2nix.overlays.default
            rust-overlay.overlays.default
            #(final: prev: self.packages.${system})
          ];
        };

        everything = pkgs.buildGoApplication {
          pname = "xe-x-composite";
          version = "1.0.0";
          src = ./.;
          modules = ./gomod2nix.toml;

          buildInputs = with pkgs; [
            pkg-config
            libaom
            libavif
            sqlite-interactive
          ];
        };

        xedn = pkgs.buildGoApplication {
          pname = "xedn";
          version = "1.2.3";
          src = ./.;
          modules = ./gomod2nix.toml;
          subPackages = [ "./cmd/xedn" ];

          buildInputs = with pkgs; [
            pkg-config
            libaom
            libavif
          ];
        };

        copyFile = { pname, path ? pname }:
          pkgs.stdenv.mkDerivation {
            inherit pname;
            inherit (everything) version;
            src = everything;

            installPhase = ''
              mkdir -p $out/bin
              cp $src/bin/$pname $out/bin/$path
            '';
          };
      in {
        overlays.default = final: prev:
          let
            system = prev.system;
            selfpkgs = self.packages.${system};
          in { xeserv = selfpkgs; };

        packages = rec {
          default = everything;

          license = copyFile {
            pname = "license";
            path = "xlicense";
          };
          makeMastodonApp = copyFile {
            pname = "mkapp";
            path = "make-mastodon-app";
          };

          inherit xedn;

          aegis = copyFile { pname = "aegis"; };
          cadeybot = copyFile { pname = "cadeybot"; };
          hlang = copyFile { pname = "hlang"; };
          johaus = copyFile { pname = "johaus"; };
          mainsanow = copyFile { pname = "mainsanow"; };
          prefix = copyFile { pname = "prefix"; };
          quickserv = copyFile { pname = "quickserv"; };
          todayinmarch2020 = copyFile { pname = "todayinmarch2020"; };
          uploud = copyFile { pname = "uploud"; };
          whoisfront = copyFile { pname = "whoisfront"; };
          within-website = copyFile { pname = "within.website"; };

          xedn-docker = pkgs.dockerTools.buildLayeredImage {
            name = "registry.fly.io/xedn";
            tag = "latest";
            contents = [ pkgs.cacert ];
            config = {
              Cmd = [ "${xedn}/bin/xedn" ];
              WorkingDir = "${xedn}";
            };
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            go-tools
            gomod2nix.packages.${system}.default
            python
            strace
            hey
            boltbrowser

            pkg-config
            libaom
            libavif
            sqlite-interactive

            cargo
            cargo-watch
            rustfmt
            rust-analyzer
            wasmtime
            binaryen
            wabt
            bloaty
            (rust-bin.stable.latest.default.override {
              extensions = [ "rust-src" ];
              targets = [ "wasm32-wasi" ];
            })
          ];
        };
      }) // {
        nixosModules = {
          overlay = { ... }: {
            nixpkgs.overlays = [
              (final: prev:
                let
                  system = prev.system;
                  selfpkgs = self.packages.${system};
                in { xeserv = selfpkgs; })
            ];
          };

          default = { ... }: {
            imports = with self.nixosModules; [
              overlay
              aegis
              hlang
              todayinmarch2020
              within-website
            ];
          };

          aegis = import ./nix/aegis.nix self;
          hlang = import ./nix/hlang.nix self;
          todayinmarch2020 = import ./nix/todayinmarch2020.nix self;
          within-website = import ./nix/within-website.nix self;
        };

        checks.x86_64-linux = let
          pkgs = nixpkgs.legacyPackages.x86_64-linux;
          common = { pkgs, ... }: {
            users.groups.within = { };
            systemd.services.within-homedir-setup = {
              description = "Creates homedirs for /srv/within services";
              wantedBy = [ "multi-user.target" ];

              serviceConfig.Type = "oneshot";

              script = with pkgs; ''
                ${coreutils}/bin/mkdir -p /srv/within
                ${coreutils}/bin/chown root:within /srv/within
                ${coreutils}/bin/chmod 775 /srv/within
                ${coreutils}/bin/mkdir -p /srv/within/run
                ${coreutils}/bin/chown root:within /srv/within/run
                ${coreutils}/bin/chmod 770 /srv/within/run
              '';
            };
          };
        in {
          basic = pkgs.nixosTest ({
            name = "basic-tests";
            nodes.default = { config, pkgs, ... }: {
              imports = [ common self.nixosModules.default ];

              xeserv.services = {
                aegis.enable = true;
                hlang.enable = true;
                todayinmarch2020.enable = true;
                within-website.enable = true;
              };
            };

            testScript = ''
              start_all()

              default.wait_for_unit("hlang.service")
              print(
                  default.wait_until_succeeds(
                      "curl --unix-socket /srv/within/run/hlang.sock http://foo/"
                  )
              )

              default.wait_for_unit("todayinmarch2020.service")
              print(
                  default.wait_until_succeeds(
                      "curl --unix-socket /srv/within/run/todayinmarch2020.sock http://foo/"
                  )
              )

              default.wait_for_unit("within-website.service")
              print(
                  default.wait_until_succeeds(
                      "curl http://127.0.0.1:52838/"
                  )
              )

              default.wait_for_unit("aegis.service")
            '';
          });
        };
      };
}
