variable "name" {
    type = string
    description = "name of db addon"
}

variable "app" {
    type = string
    description = "heroku app to provision in"
}

variable "plan" {
    type = string
    description = "plan for addon"
    default = "heroku-postgresql:hobby-dev"
}
