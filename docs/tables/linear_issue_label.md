---
title: "Steampipe Table: linear_issue_label - Query Linear Issue Labels using SQL"
description: "Allows users to query issue labels in Linear, specifically the issue label ID, name, and color, providing insights into label usage and organization within Linear issues."
---

# Table: linear_issue_label - Query Linear Issue Labels using SQL

Linear Issue Labels are a feature within Linear, a project management and issue tracking tool. They allow for categorization and organization of issues, enhancing the ability to manage and track tasks within a project. Labels can be customized with unique names and colors, providing visual cues and facilitating issue sorting and filtering.

## Table Usage Guide

The `linear_issue_label` table provides insights into issue labels within Linear's issue tracking system. As a project manager or developer, explore label-specific details through this table, including label names, colors, and associated issues. Utilize it to uncover information about label usage, such as frequency of use, the distribution of labels across issues, and the effectiveness of current label organization strategies.

## Examples

### Basic info
Explore the general details of issue labels in a project management tool. This is useful to understand the timeline and categorization of issues for better project tracking and management.

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label;
```

### List labels which are not associated with any team
Explore which issue labels are not linked to a team. This can be useful for identifying potential areas of miscommunication or disorganization within your project management system.

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  team is null;
```

### List all labels for each issue
Discover the segments that have been categorized under each issue. This is useful for understanding how issues are labelled and organized, which can aid in issue tracking and management.

```sql
select
  i.title as issue_title,
  l.title as label_tile
from
  linear_issue as i,
  linear_issue_label as l,
  jsonb_array_elements(issue_ids) as ids
where
  i.id = ids ->> 'id';
```

### Show archived labels
Uncover the details of archived labels within your project management tool. This query is useful in identifying labels that are no longer in active use, helping to streamline and organize your project management process.

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  archived_at is not null;
```

### List labels created by admin
Explore which issue labels have been created by an admin. This can be useful for understanding the administrative actions and categorisations within your project.

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  creator ->> 'admin' = 'true';
```

### List child labels of a particular label
Explore which child labels belong to a specific parent label, allowing you to understand the organization and categorization of issues within your project.

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  parent ->> 'name' = 'issueLabel';
```