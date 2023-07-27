---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/linear.svg"
brand_color: "#5E6AD2"
display_name: "Linear"
short_name: "linear"
description: "Steampipe plugin to query issues, teams, users and more from Linear."
og_description: "Query Linear with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/linear-social-graphic.png"
---

# Linear + Steampipe

[Linear](https://linear.app/) is an application that streamlines issues, sprints, and product roadmaps. It's the new standard for modern software development.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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
+----------------------------------------------------------------+---------------------------+-------------------------------------------------------+----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/linear/tables)**

## Quick start

### Install

Download and install the latest Linear plugin:

```sh
steampipe plugin install linear
```

### Credentials

| Item        | Description                                                                                                                                                               |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Linear requires an [API Token](https://developers.linear.app/docs/graphql/working-with-the-graphql-api) for all requests.                                                 |
| Permissions | API tokens have the same permission as the user who creates them, and if the user permissions change, the API key permissions also change.                                |
| Radius      | Each connection represents a single Linear Installation.                                                                                                                  |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/linear.spc`)<br />2. Credentials specified in environment variables, e.g., `LINEAR_TOKEN`. |

### Configuration

Installing the latest linear plugin will create a config file (`~/.steampipe/config/linear.spc`) with a single connection named `linear`:

```hcl
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
```

Alternatively, you can also use the standard Linear environment variables to obtain credentials **only if other argument (`token`) is not specified** in the connection:

```sh
export LINEAR_TOKEN=lin_api_0aHa1iYv9WMTLrEAoSNWlG1RHPy4N5DuM4uILY
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-linear
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
