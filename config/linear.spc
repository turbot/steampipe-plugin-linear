connection "linear" {
  plugin = "linear"

  # `token` - API token for your Linear account. It can be a personal access key or OAuth2 token. Required.
  # For more information on the API Token, please see https://developers.linear.app/docs/graphql/working-with-the-graphql-api.
  # Can also be set with the LINEAR_TOKEN environment variable.
  # token = "lin_api_0aHa1iYv9WMTLrEAoSNWlG1RHPy4N5DuM4uILY"

  # `page_size` - The requested page size per API request. Default is 50. Optional.
  # It is recommended to use lower page size when you are trying to fetch large data set to avoid complexity limit breach.
  # page_size = 50
}
