---
title: "Steampipe Table: linear_attachment - Query Linear Attachments using SQL"
description: "Allows users to query Attachments in Linear, specifically the attachment details tied to issues, providing insights into issue-related files and resources."
---

# Table: linear_attachment - Query Linear Attachments using SQL

Linear is a streamlined software designed to manage software projects, where Attachments serve as a crucial resource. Attachments in Linear refer to any files or resources attached to issues within the project management tool. This includes a variety of file types, such as images, documents, or any other file that provides additional context or information about an issue.

## Table Usage Guide

The `linear_attachment` table provides insights into attachment details within Linear's issue management system. As a project manager or software developer, explore attachment-specific details through this table, including file types, associated issues, and original upload dates. Utilize it to uncover information about issue-related resources, such as the types of files most commonly attached to issues, the issues with the most attachments, and the chronology of attachment uploads.

## Examples

### Basic info
This query allows you to gain insights into the basic information of attachments in a Linear project. It's useful when you need to quickly assess details like titles, subtitles, source types, creation and update dates, and URLs, which can help in managing and organizing project resources.

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
Identify instances where attachments have an unknown source type. This can be useful in assessing potential data integrity issues or tracking unclassified data within your system.

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
Discover the segments that contain archived attachments to understand their relevance and the time when they were last updated. This is useful in managing storage and ensuring efficient data use.

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
  archived_at is not null;
```

### List attachments where source information is unavailable
Discover the attachments that lack source information, enabling you to identify and rectify gaps in your data. This is particularly useful in maintaining data integrity and ensuring comprehensive record-keeping.

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
Explore which attachments have been created by an admin to gain insights into the source and timing of these documents. This can be useful for auditing purposes or to understand administrative contribution to content.

```sql
select
  id,
  title,
  source_type,
  created_at,
  creator
from
  linear_attachment
where
  creator ->> 'admin' = 'true';
```

### List attachments for a particular issue
Explore which attachments are linked to a specific issue. This is useful for quickly accessing all relevant files and data associated with a particular problem or topic.

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