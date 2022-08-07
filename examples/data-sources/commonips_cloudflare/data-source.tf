data "commonips_cloudflare" "example" {}

resource "local_file" "all_cidr_blocks" {
  filename = "test"
  content  = jsonencode(data.commonips_cloudflare.example.cidr_blocks)
}
