---
title: "Steampipe Table: linear_integration - Query Linear Integrations using SQL"
description: "Allows users to query Linear Integrations, providing insights into the connected third-party tools in the Linear workspace."
---

# Table: linear_integration - Query Linear Integrations using SQL

Linear Integration is a feature within Linear that enables the connection and interaction of third-party tools with the Linear workspace. It facilitates the seamless integration of various tools, enhancing productivity and collaboration among team members. Linear Integration is instrumental in streamlining workflows and tasks, making project management more efficient.

## Table Usage Guide

The `linear_integration` table provides insights into the integrated third-party tools in the Linear workspace. As a project manager or a team lead, you can explore detailed information about these integrations through this table, including their types, configurations, and associated metadata. Utilize it to manage and optimize the use of integrated tools, thereby improving team collaboration and project efficiency.

## Examples

### Basic info
Explore the creation and modification timelines of various services. This query is useful in tracking the lifecycle of each service, aiding in maintenance and management tasks.

```sql+postgres
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration;
```

```sql+sqlite
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration;
```

### List integration which are not associated with any team
Explore which integrations are not linked to any team, to understand potential areas of resource allocation or reorganization. This can help streamline your workflow and ensure all integrations are being utilized effectively.

```sql+postgres
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  team is null;
```

```sql+sqlite
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  team is null;
```

### Show github integrations
This query is designed to identify all integrations with GitHub in Linear. This can be useful for tracking and managing these integrations, ensuring they're up to date and functioning properly.

```sql+postgres
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  service = 'github';
```

```sql+sqlite
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  service = 'github';
```

### Show archived integrations
Discover the segments that have archived integrations, providing a historical view of your service integrations. This can be useful for auditing purposes or for understanding changes in your integration landscape over time.

```sql+postgres
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  archived_at is not null;
```

```sql+sqlite
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  archived_at is not null;
```

### List integrations created by admin
Identify integrations that have been created by an admin. This can be useful for auditing purposes, helping to ensure there are no unauthorized integrations in your system.

```sql+postgres
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  creator ->> 'admin' = 'true';
```

```sql+sqlite
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration
where
  json_extract(creator, '$.admin') = 'true';
```