package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"zametki/pkg/models"
)

// SnippetModel - Определяем тип который обертывает пул подключения sql.DB
type ZametModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *ZametModel) Insert(title, content string) {
	stmt := `INSERT INTO zametki (title, content, created)
    VALUES(?, ?, ?)`
	time := "2023-07-07 17:19:34"

	_, err := m.DB.Exec(stmt, title, content, time)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Добавил в БД")
}

// Get - Метод для отображения всех заметок на главной странице
func (m *ZametModel) Get() ([]models.Zametki, error) {

	// SQL запрос для получения данных
	rows, err := m.DB.Query("SELECT id, title, content, created FROM zametki")
	var zametki = []models.Zametki{}

	// Цикл по строкам,
	// используя Scan для назначения данных столбца полям структуры.
	for rows.Next() {
		var s models.Zametki
		if err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created); err != nil {
			return zametki, err
		}
		zametki = append(zametki, s)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return zametki, nil
}

// GetOne - Метод для возвращения данных заметки по её идентификатору ID.
func (m *ZametModel) GetOne(id int) (*models.Zametki, error) {

	stmt := `SELECT id, title, content, created FROM zametki
    WHERE ID = ?`
	// Используем метод QueryRow() для выполнения SQL запроса,

	row := m.DB.QueryRow(stmt, id)

	s := &models.Zametki{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Del - Метод для удаления данных заметки по её идентификатору ID.
func (m *ZametModel) Del(id int) {
	_, err := m.DB.Exec("delete from zametki WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}
}
