# Table: linear_integration

Represents the list of integrations associated with the organization.

## Examples

### Basic info

```sql
select
  id,
  created_at,
  service,
  updated_at
from
  linear_integration;
```

### List integration which are not associated with any team

```sql
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

```sql
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

```sql
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

```sql
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
