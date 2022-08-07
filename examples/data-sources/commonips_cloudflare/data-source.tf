data "commonips_cloudflare" "example" {
}

resource "local_file" "test" {
  filename = "test"
  content  = jsonencode(data.commonips_cloudflare.example.cidr_blocks)
}