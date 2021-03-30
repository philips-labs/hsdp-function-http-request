# hsdp-function-http-request

Perform HTTP requests

# Usage
```hcl
module "siderite_backend" {
  source  = "philips-labs/siderite-backend/cloudfoundry"
  version = "0.2.0"
  cf_region   = "eu-west"
  cf_org_name = "client-your-org"
  cf_space    = "prod"
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
  
  schedule {
    run_every = "1h"
  }
  
  backend {
    type = "siderite"
    credentials = module.siderite_backend.credentials
  }
}
```

# License

License is MIT
