variable "name" {
    type = string
    description = "name of cache addon"
}

variable "app" {
    type = string
    description = "heroku app to provision in"
}

variable "plan" {
    type = string
    description = "plan for addon"
    default = "heroku-redis:hobby-dev"
}
