# hsdp-function-http-request

Perform HTTP requests

# Usage
```hcl
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
