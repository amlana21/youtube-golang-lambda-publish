resource "aws_lambda_function" "golang_api_lambda" {
  function_name    = "golang_api_lambda"
  filename         = "main.zip"
  handler          = "main"
  source_code_hash = filebase64sha256("main.zip")
  role             = "${var.role_arn}"
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 30
  environment {
    variables = {
      test_value = ""      
    }
  }
}