package query

const (
	AddBook = `
		INSERT INTO
		db_go_sql.books
		(
			title,
			author,
			description
		)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
	GetBookById = `
		SELECT * FROM db_go_sql.books 
		WHERE id = $1;
	`

	GetBooks = `
		SELECT * FROM db_go_sql.books
	`
	UpdateBook = `
		UPDATE db_go_sql.books 
		SET 
			title = $1, 
			author = $2, 
			description = $3 
			WHERE id = $4 
		RETURNING *;
	`

	DeleteBook = `
		DELETE FROM db_go_sql.books 
		WHERE id = $1;
	`
)
