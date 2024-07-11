## v0.3.0 [2024-7-11]

_Enhancements_

- The `organization_id` column has now been assigned as a connection key column across all the tables which facilitates more precise and efficient querying across multiple Linear accounts. ([#34](https://github.com/turbot/steampipe-plugin-linear/pull/34))

_Bug fixes_

- Fixed the plugin to correctly check for a valid Personal Access token. ([#33](https://github.com/turbot/steampipe-plugin-linear/pull/33))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5101-2024-05-09) which ensures that `QueryData` passed to `ConnectionKeyColumns` value callback is populated with `ConnectionManager`. ([#34](https://github.com/turbot/steampipe-plugin-linear/pull/34))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#23](https://github.com/turbot/steampipe-plugin-linear/pull/23))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#23](https://github.com/turbot/steampipe-plugin-linear/pull/23))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-linear/blob/main/docs/LICENSE). ([#23](https://github.com/turbot/steampipe-plugin-linear/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#22](https://github.com/turbot/steampipe-plugin-linear/pull/22))

## v0.1.2 [2023-12-06]

_Bug fixes_

- Fixed the invalid Go module path of the plugin. ([#20](https://github.com/turbot/steampipe-plugin-linear/pull/20))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#10](https://github.com/turbot/steampipe-plugin-linear/pull/10))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#7](https://github.com/turbot/steampipe-plugin-linear/pull/7))
- Recompiled plugin with Go version `1.21`. ([#7](https://github.com/turbot/steampipe-plugin-linear/pull/7))

## v0.0.1 [2023-07-06]

_What's new?_

- New tables added
  - [linear_attachment](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_attachment)
  - [linear_comment](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_comment)
  - [linear_integration](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_integration)
  - [linear_issue](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_issue)
  - [linear_issue_label](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_issue_label)
  - [linear_organization](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_organization)
  - [linear_project](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_project)
  - [linear_team](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_team)
  - [linear_team_membership](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_team_membership)
  - [linear_user](https://hub.steampipe.io/plugins/turbot/linear/tables/linear_user)
