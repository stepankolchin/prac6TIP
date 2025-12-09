package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:200;uniqueIndex;not null"`
	Notes     []Note // 1:N
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Note struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200;not null"`
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"not null"`
	User      User
	Tags      []Tag `gorm:"many2many:note_tags;"` // M:N
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:50;uniqueIndex;not null"`
	Notes     []Note `gorm:"many2many:note_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NoteTag - явное определение join таблицы (опционально, для большего контроля)
// Если нужны дополнительные поля в join таблице, можно раскомментировать:
// type NoteTag struct {
// 	NoteID    uint      `gorm:"primaryKey"`
// 	TagID     uint      `gorm:"primaryKey"`
// 	CreatedAt time.Time // например, когда была создана связь
// }
