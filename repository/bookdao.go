package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	db "github.com/nitinigi2/practice/book-api/database"
	"github.com/nitinigi2/practice/book-api/model"
)

func SaveBook(book model.Book) error {
	// get db connection
	conn := db.GetDbConn()

	// close the db connection
	defer db.CloseDBConnection(conn)

	sqlStatement := `INSERT INTO BOOK (isbn, title, category, description, author) VALUES (?, ?, ?, ?, ?)`

	// execute the sql statement
	// Scan function will save the insert id in the id

	_, err := conn.Exec(sqlStatement, book.Isbn, book.Title, book.Category, book.Description, book.Author)

	if err != nil {
		fmt.Println("Unable to execute the query", err)
		return err
	}

	fmt.Println("Book is saved in DB successfully.")
	return nil
}

func DeleteBook(id int) error {
	// get db connection
	conn := db.GetDbConn()

	// close the db connection
	defer db.CloseDBConnection(conn)

	sqlStatement := `delete from BOOK where id = ?`

	// execute the sql statement
	_, err := conn.Exec(sqlStatement, id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(book model.Book, bookId int) error {
	// get db connection
	conn := db.GetDbConn()

	// close the db connection
	defer db.CloseDBConnection(conn)

	sqlStatement := `UPDATE BOOK SET isbn=?, title=?, category=?, description=?, author=? WHERE id=?`

	// execute the sql statement
	_, err := conn.Exec(sqlStatement, book.Isbn, book.Title, book.Category, book.Description, book.Author, bookId)

	return err
}

func GetBook(id int) (model.Book, error) {
	// get db connection
	conn := db.GetDbConn()

	// close the db connection
	defer db.CloseDBConnection(conn)

	sqlStatement := `select * from BOOK where id = ?`

	var book model.Book
	// execute the sql statement
	row := conn.QueryRow(sqlStatement, id)

	err := row.Scan(&book.ID, &book.Isbn, &book.Title, &book.Category, &book.Description, &book.Author)

	switch err {

	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return book, errors.New("no book found")

	case nil:
		return book, nil

	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return book, err
}

func GetAllBooks() ([]model.Book, error) {
	// get db connection
	conn := db.GetDbConn()

	// close the db connection
	defer db.CloseDBConnection(conn)

	sqlStatement := `select * from BOOK`

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	books := []model.Book{}

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Category, &book.Description, &book.Author)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
