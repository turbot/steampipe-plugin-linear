# Table: linear_attachment

Issue attachments allow you to link external resources to issues and display them inside Linear similarly to GitHub Pull Requests. They are designed with API developers in mind and we also use them for upcoming integrations inside Linear.

## Examples

### Basic info

```sql
select
  id,
  title,
  subtitle,
  source_type,
  created_at,
  updated_at,
  url
from
  linear_attachment;
```

### List attachments where source type is unknown

```sql
select
  id,
  title,
  subtitle,
  source_type,
  created_at,
  updated_at,
  url
from
  linear_attachment
where
  source_type = 'unknown';
```

### List archived attachments

```sql
select
  id,
  title,
  subtitle,
  source_type,
  created_at,
  updated_at,
  url
from
  linear_attachment
where
  archived is not null;
```

### List attachments where source information is unavailable

```sql
select
  id,
  title,
  subtitle,
  source_type,
  created_at,
  updated_at,
  url
from
  linear_attachment
where
  source is null;
```

### List attachments created by admin

```sql
select
  id,
  title,
  subtitle,
  source_type,
  created_at,
  updated_at,
  url
from
  linear_attachment
where
  creator ->> 'admin' = 'true';
```

### List attachments for a particular issue

```sql
select
  id,
  title,
  subtitle,
  source_type,
  created_at,
  updated_at,
  url
from
  linear_attachment
where
  issue ->> 'title' = 'attachment check';
```
