# hsdp-function-http-request

Perform HTTP requests

## Usage
```hcl
terraform {
  required_providers {
    hsdp = {
      source  = "philips-software/hsdp"
      version = ">= 0.29.0"
    }
    cloudfoundry = {
      source  = "cloudfoundry-community/cloudfoundry"
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
  version = "0.8.0"
  
  cf_region   = "us-east"
  cf_org_name = var.cf_org
  cf_space    = var.cf_space
  cf_user     = var.cf_user
  iron_plan   = "medium-encrypted"
}

resource "hsdp_function" "request" {
  name = "http-request"
  docker_image = "philipslabs/hsdp-function-http-request:v0.7.0"
  
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

## Supported environment variables

| Name               | Description                                                                                 | Example                  |
|--------------------|---------------------------------------------------------------------------------------------|--------------------------|
| REQUEST_METHOD     | The request method to use                                                                   | `POST`                   |
| REQUEST_URL        | The URL to use. Can include query params                                                    | https://myapp.io/trigger |
| REQUEST_USERNAME   | Username to use for Basic Auth                                                              |                          |
| REQUEST_PASSWORD   | Password to use for Basic Auth                                                              |                          |
| REQUEST_BODY       | The request Body to use.                                                                    |                          |
| REQUEST_HEADER_XXX | The request header to use. Where XXX is the header name. Underscores are replaced by dashes |                          |

## License

License is MIT
