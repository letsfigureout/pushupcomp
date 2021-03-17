terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
      version = "4.1.0"
    }
  }
}

# Configure the Heroku provider
provider "heroku" {
  email   = var.heroku_email
  api_key = var.heroku_api_key
}

# Create a new application
module "app" {
  source = "./modules/app"
  name = "my-heroku-goapp"
  region = "us"
}

# Create a redis cache
#module "cache" {
#  source = "./modules/cache"
#  name = "goapp-cache"
#  app = module.app.output.name
#}

# Create a postgres DB
module "db" {
  source = "./modules/db"
  name = "goapp-db"
  app = module.app.output.name
}
