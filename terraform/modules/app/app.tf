terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
      version = "4.1.0"
    }
  }
}

resource "heroku_app" "app" {
    name = var.name
    region = var.region

    buildpacks = [
        "heroku/go"
    ]
}