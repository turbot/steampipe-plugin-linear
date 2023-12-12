---
title: "Steampipe Table: linear_comment - Query Linear Comments using SQL"
description: "Allows users to query Linear Comments, specifically the details about comments on issues, projects, and cycles in Linear, providing insights into communication and collaboration patterns."
---

# Table: linear_comment - Query Linear Comments using SQL

Linear Comment is a feature within Linear that allows users to communicate and collaborate on issues, projects, and cycles. It provides a centralized way to manage and view comments made by team members, enhancing transparency and facilitating discussions. Linear Comment helps you stay informed about the progress and updates of your Linear resources and understand the context of decisions and changes.

## Table Usage Guide

The `linear_comment` table provides insights into comments within Linear. As a project manager or team lead, explore comment-specific details through this table, including the author, text, created and updated timestamps, and associated issue. Utilize it to uncover information about comments, such as those related to specific issues, the frequency of updates, and the involvement of team members in discussions.

## Examples

### Basic info
Explore the timeline of user interactions with your platform by identifying when specific comments were created, edited, or updated. This allows for a better understanding of user engagement patterns over time.

```sql+postgres
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment;
```

### Show user details of each comment
Explore the details of users who have made comments, including their active status and administrative privileges. This can be useful for understanding user engagement and identifying key contributors.

```sql+postgres
select
  id,
  title,
  comment_user ->> 'id' as creator_id,
  comment_user ->> 'name' as creator_name,
  comment_user ->> 'active' as active,
  comment_user ->> 'email' as email,
  comment_user ->> 'admin' as admin,
  comment_user ->> 'createdAt' as created_at
from
  linear_comment;
```

```sql+sqlite
select
  id,
  title,
  json_extract(comment_user, '$.id') as creator_id,
  json_extract(comment_user, '$.name') as creator_name,
  json_extract(comment_user, '$.active') as active,
  json_extract(comment_user, '$.email') as email,
  json_extract(comment_user, '$.admin') as admin,
  json_extract(comment_user, '$.createdAt') as created_at
from
  linear_comment;
```

### List comments for a particular issue
Explore the comments related to a specific issue to gain insights into its history and ongoing discussions. This can be useful for understanding the context and progression of the issue, as well as for tracking any changes or edits made over time.

```sql+postgres
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  issue ->> 'title' = 'attachment check';
```

```sql+sqlite
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  json_extract(issue, '$.title') = 'attachment check';
```

### List comments written by admin
Explore which comments have been authored by an admin user. This can help in understanding the context and engagement of administrative users in discussions.

```sql+postgres
select
  id,
  title,
  created_at,
  edited_at,
  comment_user,
  url
from
  linear_comment
where
  comment_user ->> 'admin' = 'true';
```

```sql+sqlite
select
  id,
  title,
  created_at,
  edited_at,
  comment_user,
  url
from
  linear_comment
where
  json_extract(comment_user, '$.admin') = 'true';
```

### List comments older than 90 days
Discover the segments that contain comments older than 90 days to better understand user feedback trends and manage content accordingly. This is useful for identifying outdated or irrelevant discussions and maintaining a current and engaging user experience.

```sql+postgres
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  created_at < now() - interval '90' day;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  created_at < datetime('now', '-90 day');
```