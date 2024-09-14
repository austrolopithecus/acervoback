{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    # Adicione outros pacotes necessários aqui (como `air`, etc.)
  ];
}

