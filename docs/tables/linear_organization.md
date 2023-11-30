---
title: "Steampipe Table: linear_organization - Query Linear Organizations using SQL"
description: "Allows users to query Linear Organizations, offering detailed information about each organization's name, domain, and other related attributes."
---

# Table: linear_organization - Query Linear Organizations using SQL

Linear is a project management tool that streamlines software projects, tasks, and bug tracking. It's designed for high-performance teams to drive their work forward. With Linear, teams can plan, organize, and track all their software development processes in one place.

## Table Usage Guide

The `linear_organization` table provides insights into organizations within Linear. As a project manager or team lead, you can use this table to get detailed information about your organization, including its name, domain, and associated attributes. This can be particularly useful for managing and organizing your software development processes more efficiently.

## Examples

### Basic info
Explore which organizations have enabled the roadmap feature and assess the user count for each to gain insights into their release channels. This can help in understanding the usage trends and making informed decisions.

```sql
select
  id,
  title,
  created_at,
  url_key,
  user_count,
  roadmap_enabled,
  release_channel,
  updated_at
from
  linear_organization;
```

### List teams in the organization
Explore the different teams within your organization to understand their characteristics and recent updates. This can assist in managing team-specific resources or assessing the overall structure of your organization.

```sql
select
  t.id,
  t.title,
  t.color,
  t.key,
  t.private,
  t.updated_at
from
  linear_team as t,
  linear_organization as o
where
  o.id = t.organization ->> 'id';
```

### List users in the organization
Explore which users are active within your organization and gain insights into their administrative status and last update time. This can be particularly useful for managing user roles and tracking activity.

```sql
select
  u.id,
  u.title,
  u.active,
  u.admin,
  u.email,
  u.updated_at
from
  linear_user as u,
  linear_organization as o
where
  o.id = u.organization ->> 'id';
```

### List integrations in the organization
Explore which integrations are present within the organization. This can be used to identify the tools and services being utilized, aiding in resource management and strategic planning.

```sql
select
  i.id,
  i.created_at,
  i.service,
  i.updated_at
from
  linear_integration as i,
  linear_organization as o
where
  o.id = i.organization ->> 'id';
```

### Show subscription details of the organization
Determine the specifics of your organization's subscription, including the number of seats and type of subscription, and gain insights into future billing dates and creation times. This can help in managing your organization's resources and planning for future expenses.

```sql
select
  subscription ->> 'id' as creator_id,
  subscription ->> 'nextBillingAt' as next_billing_at,
  subscription ->> 'seats' as seats,
  subscription ->> 'type' as type,
  subscription ->> 'createdAt' as created_at
from
  linear_organization;
```

### List organizations that have roadmap enabled
Discover the segments that have the roadmap feature enabled, which allows you to manage your organization's future plans and developments more effectively.

```sql
select
  id,
  title,
  created_at,
  url_key,
  user_count,
  roadmap_enabled,
  release_channel,
  updated_at
from
  linear_organization
where
  roadmap_enabled;
```

### List the organizations with SAML authentication enabled
Discover the organizations that have enabled SAML authentication to understand their security measures and manage user access more efficiently.

```sql
select
  id,
  title,
  created_at,
  url_key,
  user_count,
  saml_enabled,
  release_channel,
  updated_at
from
  linear_organization
where
  saml_enabled;
```

### List the organizations with SCIM provisioning enabled
Discover the segments that have SCIM provisioning enabled in their organizations. This can be useful in assessing the level of security and user management within these organizations.

```sql
select
  id,
  title,
  created_at,
  url_key,
  user_count,
  scim_enabled,
  release_channel,
  updated_at
from
  linear_organization
where
  scim_enabled;
```