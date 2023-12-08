---
title: "Steampipe Table: linear_issue - Query Linear Issues using SQL"
description: "Allows users to query Issues in Linear, specifically the details of each issue, providing insights into project management and task tracking."
---

# Table: linear_issue - Query Linear Issues using SQL

Linear is a project management and issue tracking tool. It helps teams plan, track, and coordinate tasks across their projects. It provides a streamlined way to manage software projects, tasks, and bug tracking, ensuring teams stay on track and meet their goals.

## Table Usage Guide

The `linear_issue` table provides insights into issues within Linear's project management tool. As a project manager or team lead, explore issue-specific details through this table, including statuses, assignees, and associated metadata. Utilize it to uncover information about issues, such as their current progress, the team members assigned to them, and their priority levels.

## Examples

### Basic info
Explore the issues in your project, including their creation date, branch name, priority, and estimate time, to manage your workflow and prioritize tasks more effectively. This query helps you gain insights into the status of different issues, allowing you to make informed decisions and streamline your project management process.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue;
```

### List urgent issues
Discover the segments that contain high-priority tasks in your workflow. This query assists in identifying urgent issues that require immediate attention, helping to prioritize and manage tasks effectively.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  priority = 1;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  priority = 1;
```

### List issues that have not been started
Identify instances where certain issues have yet to begin, allowing for better prioritization and task management.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  started_at is null;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  started_at is null;
```

### List issues that have crossed the due date
Discover the segments that have missed their deadlines by identifying issues that have surpassed their due dates. This can be useful in project management to assess the elements within your workflow that need attention or re-evaluation.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  due_date < now();
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  due_date < datetime('now');
```

### List trashed issues
Explore which issues have been moved to trash on Linear. This can assist in tracking the progress of tasks and identifying any potential bottlenecks or problems that have led to tasks being discarded.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  trashed;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  trashed = 1;
```

### List issues created by admin
Identify instances where administrative users have created issues. This allows for a quick overview of issues which may have higher priority or require more immediate attention.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  creator ->> 'admin' = 'true';
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  json_extract(creator, '$.admin') = 'true';
```

### List issues of a particular team
Explore which issues are currently being handled by a specific team. This can help in understanding the team's workload, priority tasks, and the timeline of their projects.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  team ->> 'name' = 'linear';
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  json_extract(team, '$.name') = 'linear';
```

### List unassigned issues
Explore which issues are currently unassigned, allowing you to understand bottlenecks and allocate resources more effectively. This is particularly useful for project management and ensuring tasks are not overlooked.

```sql+postgres
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  assignee is null;
```

```sql+sqlite
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  assignee is null;
```