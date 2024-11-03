resource "adguard_rewrite" "oberon_record" {
  domain = "oberon.home.local"
  answer = var.oberon_ip
}

resource "adguard_rewrite" "titan_record" {
  domain = "titan.home.local"
  answer = var.titan_ip
}

resource "adguard_rewrite" "calisto_record" {
  domain = "calisto.home.local"
  answer = var.calisto_ip
}
