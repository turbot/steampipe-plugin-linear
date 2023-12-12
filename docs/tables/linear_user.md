---
title: "Steampipe Table: linear_user - Query Linear User using SQL"
description: "Allows users to query Linear User, specifically the details of users who are part of the Linear platform, providing insights into user profiles and their associated activities."
---

# Table: linear_user - Query Linear User using SQL

Linear is a streamlined software project management service, designed to help teams prioritize tasks, manage projects, and track their progress. It provides a centralized platform to manage all aspects of a software project, from issue tracking to sprint planning. Linear helps teams stay organized, move work forward, and continuously improve their processes.

## Table Usage Guide

The `linear_user` table provides insights into users within the Linear platform. As a project manager or team lead, explore user-specific details through this table, including usernames, email addresses, and associated team information. Utilize it to uncover information about user activity, such as task assignment, project involvement, and user status within the platform.

**Important Notes**
- There are three types of users in Linear:
  - Admins: Admins have full access to a workspace and can manage team members.
  - Members: Members have read-write access to a workspace and can participate in team discussions.
  - Guests: Guests have read-only access to a workspace and cannot participate in team discussions.

## Examples

### Basic info
Explore which users are active or inactive, along with their admin status and personalized status labels. This can be particularly useful for understanding the overall user activity and administrative roles within your Linear organization.

```sql+postgres
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user;
```

```sql+sqlite
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user;
```

### List admin users
Uncover the details of which users have administrative privileges in your system. This information is crucial to understand who has elevated access and can make significant changes to your configurations.

```sql+postgres
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  admin;
```

```sql+sqlite
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  admin = 1;
```

### List inactive users
Discover the segments that consist of inactive users within your system. This allows for efficient user management, enabling you to identify and possibly re-engage or remove these dormant profiles.

```sql+postgres
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  not active;
```

```sql+sqlite
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  active = 0;
```

### Show details of the currently authenticated user
Explore the details of your user profile on Linear, including your status and administrative privileges. This can be useful for understanding your permissions and activity within the platform.

```sql+postgres
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  is_me;
```

```sql+sqlite
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  is_me = 1;
```

### List guest users
Explore which users are guests in your system. This is useful for managing access rights and ensuring appropriate levels of user permissions.

```sql+postgres
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  guest;
```

```sql+sqlite
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  guest;
```

### List archived users
Discover the segments that consist of archived users in your system. This can be beneficial to identify inactive users, assess their previous contributions, and manage system resources more effectively.

```sql+postgres
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  archived_at is not null;
```

```sql+sqlite
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  archived_at is not null;
```