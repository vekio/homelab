resource "adguard_rewrite" "proxy_record" {
  domain = "proxy.home.local"
  answer = "calisto.home.local"
}

resource "adguard_rewrite" "traefik_record" {
  domain = "traefik.vekio.dev"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "adguard_record" {
  domain = "adguard.vekio.dev"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "media_record" {
  domain = "media.vekio.dev"
  answer = "proxy.home.local"
}
