![image](https://hub.steampipe.io/images/plugins/turbot/linear-social-graphic.png)

# Linear Plugin for Steampipe

Use SQL to query issues, teams, users and more from Linear.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/linear)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/linear/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-linear/issues)

### Linear GraphQL API

- In Linear plugin we have used a library called [genqlient](https://github.com/Khan/genqlient) to generate Go code to query the Linear GraphQL API.
- Please update the [genqlient.graphql](https://github.com/turbot/steampipe-plugin-linear/blob/main/gql/genqlient.graphql) file to add or modify any query.
- After updating the query run `go run github.com/Khan/genqlient` command to generate the Go code, which then you can use as per your requirement.

## Quick start

Download and install the latest Linear plugin:

```bash
steampipe plugin install linear
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/linear#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/linear#configuration).

### Configuring Linear Credentials

Configure your account details in `~/.steampipe/config/linear.spc`:

You may specify the Token to authenticate:

- `token`: The linear API token. It can be a personal access key or oAuth2 token.

```hcl
connection "linear" {
  plugin  = "linear"
  token   = "lin_api_0aHa1iYv9WMTLrEAoSNWlG1RHPy4N5DuM4uILY"
}
```

or you may specify the Token and Page size to authenticate:

- `token`: The linear API token.
- `page_size`: The page size per API request. The default is 50.

```hcl
connection "linear" {
  plugin    = "linear"
  token     = "lin_api_0aHa1iYv9WMTLrEAoSNWlG1RHPy4N5DuM4uILY"
  page_size = 100
}
```

or through environment variables

```sh
export LINEAR_TOKEN="lin_api_0aHa1iYv9WMTLrEAoSNWlG1RHPy4N5DuM4uILY"
```

Run steampipe:

```shell
steampipe query
```

List your Linear issues:

```sql
select
  title,
  created_at,
  branch_name,
  priority
from
  linear_issue;
```

```
+----------------------------------------------------------------+---------------------------+-------------------------------------------------------+----------+
| title                                                          | created_at                | branch_name                                           | priority |
+----------------------------------------------------------------+---------------------------+-------------------------------------------------------+----------+
| ProTip: Mouse over this issue & press [Space]                  | 2023-05-09T12:41:21+05:30 | sourav/tur-8-protip-mouse-over-this-issue-press-space | 4        |
| test linear                                                    | 2023-05-09T12:43:21+05:30 | sourav/tur-11-test-linear                             | 0        |
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-linear.git
cd steampipe-plugin-linear
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/linear.spc
```

Try it!

```
steampipe query
> .inspect linear
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-linear/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Linear Plugin](https://github.com/turbot/steampipe-plugin-linear/labels/help%20wanted)
