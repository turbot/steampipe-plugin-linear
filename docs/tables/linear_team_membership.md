---
title: "Steampipe Table: linear_team_membership - Query Linear Team Memberships using SQL"
description: "Allows users to query Team Memberships in Linear, specifically providing details about the relationship between teams and members within the Linear service."
---

# Table: linear_team_membership - Query Linear Team Memberships using SQL

Linear Team Memberships are a critical part of the Linear service, defining the relationship between teams and their members. They provide a mechanism for grouping members into teams for better organization and management. Team Memberships in Linear are essential for managing workload distribution, task assignment, and overall project management within the service.

## Table Usage Guide

The `linear_team_membership` table provides insights into the team memberships within Linear. As a project manager or team leader, you can explore specific details about the relationship between teams and members through this table. Use it to understand team composition, member roles, and to manage workload distribution and task assignments more effectively.

## Examples

### Basic info
Explore the creation and modification details of team memberships in Linear. This information can help to understand team dynamics and track changes over time.

```sql
select
  id,
  created_at,
  owner,
  sort_order,
  updated_at
from
  linear_team_membership;
```

### List teams with owner details
Discover the segments that include team ownership details, allowing you to understand who holds control over different teams and when the last updates were made. This can be particularly useful in larger organizations where team ownership may shift frequently.

```sql
select
  id,
  jsonb_pretty(team) as team,
  jsonb_pretty(membership_user) as user,
  updated_at
from
  linear_team_membership
where
  owner;
```

### List members of a particular team
Explore which members belong to a specific team, gaining insights into their user ID, name, admin status, email, and activity status. This is particularly useful for managing team dynamics and understanding the roles of different team members.

```sql
select
  membership_user ->> 'id' as user_id,
  membership_user ->> 'name' as name,
  membership_user ->> 'admin' as admin,
  membership_user ->> 'email' as email,
  membership_user ->> 'active' as active
from
  linear_team_membership
where
  team ->> 'name' = 'linear_team';
```

### List archived membership
Explore which team memberships in Linear have been archived to manage team structure and access effectively. This could be useful in maintaining the organization's data hygiene and understanding team dynamics over time.

```sql
select
  id,
  created_at,
  owner,
  sort_order,
  updated_at
from
  linear_team_membership
where
  archived_at is not null;
```