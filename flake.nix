# nix-direnv cache busting line: sha256-vJQADsOFkEuye9tYoMQjcNAc9P4kvDyqJOTOXpkgw3A=

{
  description = "/x/perimental code";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils";

    naersk = {
      url = "github:nix-community/naersk";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    rust-overlay = {
      url = "github:oxalica/rust-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "utils";
    };

    xess = {
      url = "github:Xe/Xess";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };

    ## don't follow nixpkgs because that causes a very long font build process.
    iaso-fonts.url = "github:Xe/iosevka";
  };

  outputs = { self, nixpkgs, utils, rust-overlay, naersk, xess, iaso-fonts }@inputs:
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
              go = prev.go_1_21;
            })
            rust-overlay.overlays.default
            #(final: prev: self.packages.${system})
          ];
        };

        vendorSha256 = pkgs.lib.fileContents ./.go.mod.sri;

        version = "${self.sourceInfo.lastModifiedDate}";

        rust = pkgs.rust-bin.stable.latest.default.override {
          extensions = [ "rust-src" "rust-analysis" "clippy" "rustfmt" ];
          targets = [ "wasm32-wasi" "wasm32-unknown-unknown" "riscv64gc-unknown-linux-gnu" ];
        };

        naersk' = pkgs.callPackage naersk {
          cargo = rust;
          rustc = rust;
        };

        everything = pkgs.buildGo121Module {
          pname = "xe-x-composite";
          inherit version vendorSha256;
          src = ./.;

          nativeBuildInputs = with pkgs; [ pkg-config ];
          buildInputs = with pkgs; [
            pkg-config
            libaom
            libavif
            sqlite-interactive
          ];
        };

        mimi = pkgs.buildGo121Module {
          pname = "mimi";
          inherit version vendorSha256;
          src = ./.;
          subPackages = [ "cmd/mimi" ];
        };

        xedn = pkgs.buildGo121Module {
          pname = "xedn";
          inherit version vendorSha256;
          src = ./.;
          subPackages = [ "cmd/xedn" ];

          nativeBuildInputs = with pkgs; [ pkg-config ];
          buildInputs = with pkgs; [ pkg-config libaom libavif ];
        };

        xedn-static = pkgs.stdenvNoCC.mkDerivation {
          dontUnpack = true;
          pname = "xedn-static";
          version = xedn.version;

          buildPhase = ''
            mkdir -p $out/static/css/iosevka
            mkdir -p $out/static/pkg/iosevka
            mkdir -p $out/static/pkg/xess
            mkdir -p $out/static/pkg/xeact

            tar xf ${
              pkgs.fetchurl {
                name = "xeact-0.69.71";
                url =
                  "https://registry.npmjs.org/@xeserv/xeact/-/xeact-0.69.71.tgz";
                sha256 = "19rfg5fbiz69vh4hzg6694szm9agz3hfk4sfsfj6ws0cq4ss805l";
              }
            }

            mkdir -p $out/static/pkg/xeact/0.69.71
            cp -vrf ./package/* $out/static/pkg/xeact/0.69.71
            rm -rf package

            tar xf ${
              pkgs.fetchurl {
                name = "xeact-0.70.0";
                url =
                  "https://registry.npmjs.org/@xeserv/xeact/-/xeact-0.70.0.tgz";
                sha256 = "1mxwrs04vj1mixs418cy7gd5dy3nfaqcc00vmrjwsgzc36wnmwxz";
              }
            }

            mkdir -p $out/static/pkg/xeact/0.70.0
            cp -vrf ./package/* $out/static/pkg/xeact/0.70.0
            rm -rf package

            tar xf ${
              pkgs.fetchurl {
                name = "xeact-0.71.0";
                url =
                  "https://registry.npmjs.org/@xeserv/xeact/-/xeact-0.71.0.tgz";
                sha256 = "069qky042lmlvslbgsyl9j8jmkhl2hr6hy5slf1b03g3fzvkp6x4";
              }
            }

            mkdir -p $out/static/pkg/xeact/0.71.0
            cp -vrf ./package/* $out/static/pkg/xeact/0.71.0
            rm -rf package

            ln -s ${
              xess.packages.${system}.aoi
            }/static/css/xess.css $out/static/css/xess.css
            for file in ${iaso-fonts.packages.${system}.default}/*; do
              ln -s $file $out/static/css/iosevka
              ln -s $file $out/static/pkg/iosevka
            done

            for file in ${xess.packages.${system}.combined}/*; do
              ln -s $file $out/static/pkg/xess
            done
          '';
        };

        robocadey2 = pkgs.buildGo121Module {
          pname = "robocadey2";
          inherit version vendorSha256;
          src = ./.;
          nativeBuildInputs = with pkgs; [ pkg-config ];
          subPackages = [ "cmd/robocadey2" ];
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

        gowasi = pkgs.writeShellScriptBin "gowasi" ''
          export GOOS=wasip1
          export GOARCH=wasm
          exec ${pkgs.go}/bin/go $*
        '';
      in {
        overlays.default = final: prev:
          let
            system = prev.system;
            selfpkgs = self.packages.${system};
          in { xeserv = selfpkgs; };

        packages = rec {
          default = everything;

          mastosan-wasm = naersk'.buildPackage {
            src = ./web/mastosan;
            targets = [ "wasm32-wasi" ];
          };

          license = copyFile {
            pname = "license";
            path = "xlicense";
          };
          makeMastodonApp = copyFile {
            pname = "mkapp";
            path = "make-mastodon-app";
          };

          inherit xedn xedn-static robocadey2 mimi;

          aegis = copyFile { pname = "aegis"; };
          cadeybot = copyFile { pname = "cadeybot"; };
          hlang = copyFile { pname = "hlang"; };
          johaus = copyFile { pname = "johaus"; };
          mainsanow = copyFile { pname = "mainsanow"; };
          prefix = copyFile { pname = "prefix"; };
          quickserv = copyFile { pname = "quickserv"; };
          sanguisuga = copyFile { pname = "sanguisuga"; };
          todayinmarch2020 = copyFile { pname = "todayinmarch2020"; };
          uploud = copyFile { pname = "uploud"; };
          vest-pit-near = copyFile { pname = "vest-pit-near"; };
          whoisfront = copyFile { pname = "whoisfront"; };
          within-website = copyFile { pname = "within.website"; };
          yeet = copyFile { pname = "yeet"; };
        };

        legacyPackages = {
          docker = let
            robocadey2 = self.packages.${system}.robocadey2;
            xedn = self.packages.${system}.xedn;
            mimi = self.packages.${system}.mimi;
          in {
            mimi = pkgs.dockerTools.buildLayeredImage {
              name = "registry.fly.io/mimi";
              tag = "latest";
              contents = [ pkgs.cacert ];
              config = {
                Cmd = [ "${mimi}/bin/mimi" ];
                WorkingDir = "${mimi}";
              };
            };
            robocadey2 = pkgs.dockerTools.buildLayeredImage {
              name = "registry.fly.io/xe-robocadey2";
              tag = "latest";
              contents = [ pkgs.cacert ];
              config = {
                Cmd = [ "${robocadey2}/bin/robocadey2" ];
                WorkingDir = "${robocadey2}";
              };
            };
            xedn = pkgs.dockerTools.buildLayeredImage {
              name = "registry.fly.io/xedn";
              tag = "latest";
              contents = [ pkgs.cacert ];
              config = {
                Cmd = [ "${xedn}/bin/xedn" ];
                WorkingDir = "${xedn}";
                Env = [ "XEDN_STATIC=${self.packages.${system}.xedn-static}" ];
              };
            };
          };
          portable = {
            xedn = let
              service = pkgs.substituteAll {
                name = "xedn.service";
                src = ./run/xedn.service.in;
                xedn = self.packages.${system}.xedn;
              };
            in pkgs.portableService {
              inherit (self.packages.${system}.xedn) version;
              pname = "xedn";
              description = "Xe's CDN service";
              homepage = "https://xeiaso.net";
              units = [ service ];
              symlinks = [{
                object = "${pkgs.cacert}/etc/ssl";
                symlink = "/etc/ssl";
              }];
            };
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            go-tools
            python3
            strace
            hey
            boltbrowser
            opentofu
            litestream
            nodejs
            deno
            esbuild
            tailwindcss

            python3Packages.huggingface-hub
            python3Packages.datasets
            
            perl
            
            pkg-config
            libaom
            libavif
            sqlite-interactive

            jq
            jo

            cargo-watch
            wasmtime
            binaryen
            wabt
            bloaty
            rust
            flyctl

            tinyemu
            zig

            gowasi
          ];

          GOAMD64 = "v3";
          XEDN_STATIC = self.packages.${system}.xedn-static;
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
              sanguisuga
              vest-pit-near
              within-website
            ];
          };

          aegis = import ./nix/aegis.nix self;
          hlang = import ./nix/hlang.nix self;
          todayinmarch2020 = import ./nix/todayinmarch2020.nix self;
          sanguisuga = import ./nix/sanguisuga.nix self;
          vest-pit-near = import ./nix/vest-pit-near.nix self;
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
