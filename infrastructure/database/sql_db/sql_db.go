package sql_db

import (
	"database/sql"
	"fmt"
	model "reactionservice/internal/model/domain"
	customerror "reactionservice/internal/model/error"

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
	query := `
		INSERT INTO reactionservice.likePosts (
        	postId, 
        	username
    	) VALUES ($1, $2)
	`
	err := sd.insertData(query, like.PostId, like.Username)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create likePost, username: %s -> postId: %s", like.Username, like.PostId)
		return err
	}

	log.Info().Msgf("LikePost created successfully, username: %s -> postId: %s", like.Username, like.PostId)
	return nil
}

func (sd *SqlDatabase) CreateSuperlikePost(superlike *model.SuperlikePost) error {
	query := `
		INSERT INTO reactionservice.superlikePosts (
        	postId, 
        	username
    	) VALUES ($1, $2)
	`

	err := sd.insertData(query, superlike.PostId, superlike.Username)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to create superlikePost, username: %s -> postId: %s", superlike.Username, superlike.PostId)
		return err
	}

	log.Info().Msgf("SuperlikePost created successfully, username: %s -> postId: %s", superlike.Username, superlike.PostId)
	return nil
}

func (sd *SqlDatabase) insertData(query string, args ...any) error {
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

	_, err = tx.Exec(
		query,
		args...)

	if err != nil {
		return err
	}

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
			log.Info().Msgf("No likePost found for postId: %s and username: %s", postId, username)
			return nil, nil
		}

		log.Error().Stack().Err(err).Msgf("Failed to get likePost for postId: %s and username: %s", postId, username)
		return nil, err
	}

	log.Info().Msgf("LikePost retrieved successfully for postId: %s and username: %s", postId, username)
	return like, nil
}

func (sd *SqlDatabase) GetSuperlikePost(postId, username string) (*model.SuperlikePost, error) {
	query := `
		SELECT postId, username
		FROM reactionservice.superlikePosts
		WHERE postId = $1 AND username = $2
	`

	row := sd.Client.QueryRow(query, postId, username)

	superlike := &model.SuperlikePost{}
	err := row.Scan(&superlike.PostId, &superlike.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msgf("No superlikePost found for postId: %s and username: %s", postId, username)
			return nil, nil
		}

		log.Error().Stack().Err(err).Msgf("Failed to get superlikePost for postId: %s and username: %s", postId, username)
		return nil, err
	}

	log.Info().Msgf("SuperlikePost retrieved successfully for postId: %s and username: %s", postId, username)
	return superlike, nil
}

func (sd *SqlDatabase) DeleteLikePost(data *model.LikePost) error {
	query := `
		DELETE FROM reactionservice.likePosts
		WHERE postId = $1 AND username = $2
	`
	err := sd.deleteData(query, data.PostId, data.Username)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to delete likePost, username: %s -> postId: %s", data.Username, data.PostId)
		return err
	}

	log.Info().Msgf("LikePost deleted successfully, username: %s -> postId: %s", data.Username, data.PostId)

	return nil
}

func (sd *SqlDatabase) DeleteSuperlikePost(data *model.SuperlikePost) error {
	query := `
		DELETE FROM reactionservice.superlikePosts
		WHERE postId = $1 AND username = $2
	`
	err := sd.deleteData(query, data.PostId, data.Username)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Failed to delete superlikePosts, username: %s -> postId: %s", data.Username, data.PostId)
		return err
	}

	log.Info().Msgf("SuperlikePosts deleted successfully, username: %s -> postId: %s", data.Username, data.PostId)

	return nil
}

func (sd *SqlDatabase) deleteData(query string, args ...any) error {
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

	result, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Stack().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		err = customerror.NewNotFoundError()
		log.Error().Msg("Data not found")
		return err
	}

	return nil
}
