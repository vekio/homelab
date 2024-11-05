resource "adguard_rewrite" "proxy_record" {
  domain = "proxy.home.local"
  answer = "calisto.home.local"
}

resource "adguard_rewrite" "traefik_record" {
  domain = "traefik.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "adguard_record" {
  domain = "adguard.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "media_record" {
  domain = "media.${var.domain}"
  answer = "proxy.home.local"
}
