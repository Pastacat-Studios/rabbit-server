{ pkgs ? import <nixpkgs> {}}:

pkgs.mkShell {
  packages = [ 
    pkgs.opencv
    pkgs.gcc
    pkgs.go
    pkgs.delve
    pkgs.gopls
    ];
  hardeningDisable = [ "fortify" ];
}