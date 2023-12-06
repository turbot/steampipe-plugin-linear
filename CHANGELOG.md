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
