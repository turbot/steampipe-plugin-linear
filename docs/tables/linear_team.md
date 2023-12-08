---
title: "Steampipe Table: linear_team - Query Linear Teams using SQL"
description: "Allows users to query Linear Teams, specifically the team details, providing insights into team structures and their respective workflows."
---

# Table: linear_team - Query Linear Teams using SQL

Linear is a software development management tool that allows teams to streamline their workflow. It provides a centralized platform for managing tasks, tracking progress, and coordinating team efforts. With Linear, teams can effectively collaborate, manage their tasks, and track their progress to ensure efficient project completion.

## Table Usage Guide

The `linear_team` table provides insights into team structures within Linear. As a project manager or team lead, explore team-specific details through this table, including team members, associated tasks, and progress status. Utilize it to uncover information about teams, such as their workflows, the distribution of tasks, and the overall progress of the team in the project.

## Examples

### Basic info
Explore the characteristics of your team settings on Linear, such as creation date, privacy settings, and triage capabilities. This can help you understand the current configuration and make necessary adjustments for better project management.

```sql+postgres
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team;
```

### List public teams
Explore the various public teams within your organization, allowing you to assess their characteristics and settings. This can help in understanding team structures and their respective configurations, which can be beneficial in managing resources and planning workflows.

```sql+postgres
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  not private;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  not private;
```

### List teams that are not using cycles
Discover the teams that have opted not to use cycles, providing insights into their work methodology and allowing for potential process optimization. This can be useful for assessing the effectiveness of various team strategies.

```sql+postgres
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  not cycles_enabled;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  not cycles_enabled;
```

### List teams that are not associated with any integration
Discover the teams that lack any associated integrations, enabling you to assess areas for potential improvement in team collaboration and productivity.

```sql+postgres
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  integrations_settings is null;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  integrations_settings is null;
```

### List archived teams
Explore which teams have been archived in your Linear organization. This can help you keep track of past teams and their settings, providing valuable context for organizational planning and resource allocation.

```sql+postgres
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  archived_at is not null;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  archived_at is not null;
```

### List teams where triage mode is enabled
Discover the segments that have the triage mode activated, allowing you to understand which teams are set up for crisis management. This is particularly beneficial in assessing the readiness of different teams in handling urgent issues.

```sql+postgres
select
  id,
  title,
  created_at,
  color,
  private,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  triage_enabled;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  color,
  private,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  triage_enabled = 1;
```