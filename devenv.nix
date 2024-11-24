{ pkgs, lib, config, inputs, ... }:

{

  # https://devenv.sh/packages/
  packages = with pkgs;[ git air ];

  # https://devenv.sh/languages/
  languages.go.enable = true;
  services.postgres = {
    enable = true;
    package = pkgs.postgresql_16;
    listen_addresses = "127.0.0.1";
    port = 5435;
    initialDatabases = [
      {
      name = "acervo";
      }
    ];
  };
}
