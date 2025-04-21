package sql_db

import (
	"database/sql"
	"fmt"
	model "reactionservice/internal/model/domain"

	"github.com/rs/zerolog/log"
)

type SqlDatabase struct {
	Client *sql.DB
}

func NewDatabase(connStr string) (*SqlDatabase, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Couldn't open a connection with the database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Database is not reachable")
		return nil, err
	}

	return &SqlDatabase{
		Client: db,
	}, nil
}

func (sd *SqlDatabase) Clean() {
	tx, err := sd.Client.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Clean each table
	for _, table := range tables {
		query := fmt.Sprintf("DELETE FROM reactionservice.%s", table)
		_, err = tx.Exec(query)
		if err != nil {
			log.Error().Stack().Err(err).Msgf("Failed to clean table %s", table)
		}
	}

	log.Info().Msg("Database cleaned successfully")
	return
}

func (sd *SqlDatabase) CreateLikePost(like *model.LikePost) error {
	tx, err := sd.Client.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
		INSERT INTO reactionservice.likePosts (
        	postId, 
        	username
    	) VALUES ($1, $2)
	`
	_, err = tx.Exec(
		query,
		like.PostId,
		like.Username)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create likepost, username: %s -> postId: %s", like.Username, like.PostId)
		return err
	}

	log.Info().Msgf("Like created successfully, username: %s -> postId: %s", like.Username, like.PostId)
	return nil
}

func (sd *SqlDatabase) GetLikePost(postId, username string) (*model.LikePost, error) {
	query := `
		SELECT postId, username
		FROM reactionservice.likePosts
		WHERE postId = $1 AND username = $2
	`

	row := sd.Client.QueryRow(query, postId, username)

	like := &model.LikePost{}
	err := row.Scan(&like.PostId, &like.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msgf("No like found for postId: %s and username: %s", postId, username)
			return nil, nil
		}

		log.Error().Stack().Err(err).Msgf("Failed to get like for postId: %s and username: %s", postId, username)
		return nil, err
	}

	log.Info().Msgf("Like retrieved successfully for postId: %s and username: %s", postId, username)
	return like, nil
}

func (sd *SqlDatabase) DeleteLikePost(data *model.LikePost) error {
	tx, err := sd.Client.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
		DELETE FROM reactionservice.likePosts
		WHERE postId = $1 AND username = $2
	`
	result, err := tx.Exec(query, data.PostId, data.Username)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to delete likepost, username: %s -> postId: %s", data.Username, data.PostId)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Msgf("No like found to delete for username: %s -> postId: %s", data.Username, data.PostId)
	} else {
		log.Info().Msgf("Like deleted successfully, username: %s -> postId: %s", data.Username, data.PostId)
	}

	return nil
}
