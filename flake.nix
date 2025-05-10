{
  description = "Reproducible proto tool‑chain for flow‑proto";

  inputs = {
    nixpkgs.url     = "github:NixOS/nixpkgs/nixos-24.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go_1_22
            protobuf              # protoc
            protoc-gen-go
            protoc-gen-go-grpc
            git
            gnumake
          ];

          shellHook = ''
            # Set custom prompt
            export PS1="\[\033[1;32m\][nix:flow-proto]\[\033[0m\] \[\033[1;34m\]\w\[\033[0m\] \[\033[1;36m\]\$\[\033[0m\] "
            
            echo "Flow proto development environment ready!"
            echo "Go      : $(go version)"
            echo "protoc  : $(protoc --version)"
            echo "protoc‑gen‑go       -> $(protoc-gen-go --version 2>/dev/null || echo 'ok')"
            echo "protoc‑gen‑go‑grpc  -> $(protoc-gen-go-grpc --version 2>/dev/null || echo 'ok')"
          '';
        };
      });
}
