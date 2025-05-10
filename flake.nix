{
  description = "Reproducible proto tool‑chain for flow‑proto";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";       # pick a channel
    go-grpc.url = "github:bufbuild/protoc-gen-go-grpc/v1.4.0";
    go-proto.url = "github:protocolbuffers/protobuf-go/v1.36.0";
  };

  outputs = { self, nixpkgs, go-grpc, go-proto }:
  let
    pkgs = import nixpkgs { system = "x86_64-linux"; };
    # Build the two Go plugins once and cache them in the store.
    protoc-gen-go = pkgs.buildGoModule {
      name = "protoc-gen-go";
      src  = go-proto;
      vendorSha256 = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="; # fill in after first build
      subPackages = [ "cmd/protoc-gen-go" ];
    };

    protoc-gen-go-grpc = pkgs.buildGoModule {
      name = "protoc-gen-go-grpc";
      src  = go-grpc;
      vendorSha256 = "sha256-BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB=";
      subPackages = [ "." ];
    };
  in
  {
    devShells.default = pkgs.mkShell {
      buildInputs = [
        pkgs.go_1_22                # Go tool‑chain
        pkgs.protobuf               # `protoc`
        protoc-gen-go
        protoc-gen-go-grpc
      ];
      # Put the plugins on PATH so `protoc` finds them.
      shellHook = ''
        export PATH=${protoc-gen-go}/bin:${protoc-gen-go-grpc}/bin:$PATH
        echo "go: $(go version)"
        echo "protoc: $(protoc --version)"
      '';
    };
  };
}
