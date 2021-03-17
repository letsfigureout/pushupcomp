terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
      version = "4.1.0"
    }
  }
}

resource "heroku_addon" "cache" {
    name = var.name
    app = var.app
    plan = var.plan
}