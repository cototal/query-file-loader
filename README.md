# Query File Loader

This creates a Go file based on a SQL file, so that the SQL file can have readable formatting and syntax highlighting.

A SQL file like this, in the Go project's `db/queries.sql` file:

```sql
-- SelectOneEntry
SELECT * FROM Entry
WHERE id = ?

-- SelectAllEntries
SELECT * FROM Entry
```

Will generate a Go file in the Go project's same `db` directory:

```go
package db

func SelectOneEntry() string {
	return "SELECT * FROM Entry WHERE id = ?"
}

func SelectAllEntries() string {
	return "SELECT * FROM Entry"
}
```

Run by specifying the file name of the SQL file you want to parse.

```
query-file-loader db/queries.sql
```