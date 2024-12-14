resource "adguard_rewrite" "proxy_record" {
  domain = "proxy.home.local"
  answer = "calisto.home.local"
}

resource "adguard_rewrite" "traefik_record" {
  domain = "traefik.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "traefik_int_record" {
  domain = "traefik-int.${var.domain}"
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

resource "adguard_rewrite" "torrent_record" {
  domain = "torrent.${var.domain}"
  answer = "proxy.home.local"
}

resource "adguard_rewrite" "portainer_record" {
  domain = "portainer.${var.domain}"
  answer = "proxy.home.local"
}
