resource "adguard_rewrite" "proxy_record" {
  domain = "proxy.home.local"
  answer = "calisto.home.local"
}

resource "adguard_rewrite" "traefik1_record" {
  domain = "traefik1.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "traefik2_record" {
  domain = "traefik2.${var.domain}"
  answer = "oberon.home.local"
}

resource "adguard_rewrite" "whoami_record" {
  domain = "whoami.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "status_record" {
  domain = "status.${var.domain}"
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

resource "adguard_rewrite" "auth_record" {
  domain = "auth.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "error_record" {
  domain = "error.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "torrent_record" {
  domain = "torrent.${var.domain}"
  answer = "proxy.home.local"
}