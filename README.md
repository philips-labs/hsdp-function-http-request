# hsdp-function-http-request

Perform HTTP requests

# Usage
```hcl
terraform {
  required_providers {
    hsdp = {
      source  = "philips-software/hsdp"
      version = ">= 0.14.2"
    }
    cloudfoundry = {
      source  = "philips-labs/cloudfoundry"
      version = ">= 0.1400.0"
    }
  }
}

provider "hsdp" {
}

provider "cloudfoundry" {
  api_url  = data.hsdp_config.cf.url
  user     = var.cf_user
  password = var.cf_password
}

data "hsdp_config" "cf" {
  region  = "us-east"
  service = "cf"
}

module "siderite_backend" {
  source  = "philips-labs/siderite-backend/cloudfoundry"
  version = "0.4.0"
  cf_region   = "us-east"
  cf_org_name = var.cf_org
  cf_space    = var.cf_space
  cf_user     = var.cf_user
  iron_plan   = "medium-encrypted"
}

resource "hsdp_function" "request" {
  name = "http-request"
  docker_image = "philipslabs/hsdp-function-http-request:v0.1.0"
  
  environment = {
    REQUEST_METHOD   = "POST"
    REQUEST_URL      = "https://myapp.io/trigger"
    REQUEST_USERNAME = "r0n"
    REQUEST_PASSWORD = "SwaNs0n"
  }
 
  backend {
    credentials = module.siderite_backend.credentials
  }
}
```

# License

License is MIT
