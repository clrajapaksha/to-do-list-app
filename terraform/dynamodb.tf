resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "Task"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }

  ttl {
    attribute_name = "TimeToExist"
    enabled        = true
  }


  tags = {
    Name        = "dynamodb-table-1"
    Environment = "production"
  }
}