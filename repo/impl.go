package repo

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/ISS-Dating/service-analyzer/model"
)

var (
	statsFields = []string{
		"id", "banned_before", "users_met", "messages_sent", "average_message_length", "links_in_messages", "user_id",
	}
)

type Repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) ReadUserStatsById(id uint64) (model.Stats, error) {
	var stats model.Stats

	tx, err := r.db.Begin()
	if err != nil {
		return model.Stats{}, err
	}

	row := tx.QueryRow("SELECT "+strings.Join(statsFields, ", ")+" FROM \"stats\" WHERE user_id=$1", id)
	if err := row.Scan(getModifyStatsFields(&stats)...); err != nil {
		return model.Stats{}, err
	}

	tx.Commit()
	return stats, nil
}

func (r *Repository) UpdateUserStats(stats model.Stats) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var list []interface{}
	list = append(list, stats.ID)
	list = append(list, getReadStatsFields(stats)[1:]...)
	_, err = tx.Exec("UPDATE \"stats\" SET "+generateEqualsPlaceholder(statsFields[1:], 2)+" WHERE id=$1",
		list...)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func generateEqualsPlaceholder(arr []string, start int) string {
	var placeholders []string
	for i, field := range arr {
		placeholders = append(placeholders, field+"=$"+strconv.Itoa(i+start))
	}

	return strings.Join(placeholders, ", ")
}

func getReadStatsFields(stats model.Stats) []interface{} {
	var fields []interface{}
	fields = append(fields,
		stats.ID,
		stats.BannedBefore,
		stats.UsersMet,
		stats.MessagesSent,
		stats.AverageMessageLen,
		stats.LinksInMessages,
		stats.UserID,
	)

	return fields
}

func getModifyStatsFields(stats *model.Stats) []interface{} {
	var fields []interface{}
	fields = append(fields,
		&stats.ID,
		&stats.BannedBefore,
		&stats.UsersMet,
		&stats.MessagesSent,
		&stats.AverageMessageLen,
		&stats.LinksInMessages,
		&stats.UserID,
	)

	return fields
}
