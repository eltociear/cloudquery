# Datadog Plugin

The CloudQuery Datadog plugin extracts your Datadog information and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

CloudQuery requires only `DD_CLIENT_API_KEY` and `DD_CLIENT_APP_KEY`. Follow [this guide](https://docs.datadoghq.com/account_management/api-app-keys/) for how to create an API key and app key for CloudQuery.

CloudQuery requires only *read* permissions (we will never make any changes to your Datadog account),
so, following the principle of the least privilege, it's recommended to grant it read-only permissions.

## Query Examples

### Find all not verified users

```sql
SELECT 
    "attributes" ->> 'name' AS username
FROM
    datadog_users
WHERE
    ("attributes" ->> 'verified')::boolean is distinct FROM TRUE
```