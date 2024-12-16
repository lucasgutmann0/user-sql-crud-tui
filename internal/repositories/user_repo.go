package repositories

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/lucasgutmann0/user-sql-crud-tui/internal/models"
	"golang.org/x/crypto/bcrypt"
)

const layerName = "repo"

type UserRepo interface {
	GetAllUsers(context.Context) ([]models.User, error)
	InsertUser(context.Context, models.User) error
	GetOneUserByID(context.Context, int) (*models.User, error)
	DeleteUserByID(context.Context, int) error
	UpdateUserByID(context.Context, int, models.User) error
}

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo creates a new repository for user-related database operations.
//
// Parameters:
//   - db: A pointer to the database connection for database interactions
//
// Returns a UserRepo interface implementation that can be used
// for user-related database operations
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

// GetAllUsers retrieves a list with the users stored in the DB
// Returns a list of users and an error
func (r *userRepo) GetAllUsers(
	ctx context.Context,
) ([]models.User, error) {
	var users []models.User
	err := r.db.SelectContext(
		ctx, &users, "SELECT id, name, email, role, age FROM users")
	if err != nil {
		slog.Error(
			"Failed to query all users", "error", err,
			"method", "GetAllUsers", "layer", layerName)
		return nil, err
	}

	return users, nil
}

// GetOneUserByID retrieves user information stored in the DB.
//
// Parameters:
//   - ctx: Manage metadata and signals related to requests.
//   - userID: The unique identifier of the user to be found.
//
// Returns the user information and an error if something fails.
func (r *userRepo) GetOneUserByID(
	ctx context.Context, userID int,
) (*models.User, error) {
	var user models.User
	err := r.db.GetContext(
		ctx, &user, "SELECT id, name, role, age FROM users WHERE id = $1", userID)
	if err != nil {
		slog.Error(
			"Failed to query one user by id", "error", err, "method",
			"GetOneUserByID", "layer", layerName)
		return nil, err
	}

	return &user, nil
}

// DeleteUserByID deletes a user registry stored in the DB.
//
// Parameters:
//   - ctx: Manage metadata and signals related to requests.
//   - userID: The unique identifier of the user to be deleted.
//
// Returns an error if the deletion fails.
func (r *userRepo) DeleteUserByID(
	ctx context.Context, userID int,
) error {
	result, err := r.db.ExecContext(
		ctx, "DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		slog.Error(
			fmt.Sprintf(
				"Failed to delete user with ID %d from the DB", userID),
			"error", err, "method", "DeleteUserByID", "layer", layerName)
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		slog.Error(
			"Failed to check rows affected at user deletion", "error", err,
			"method", "DeleteUserByID", "layer", layerName)
	}

	if affectedRows < 1 {
		slog.Warn("No user was deleted from DB")
	} else {
		slog.Info(fmt.Sprintf("Users deleted: %d", affectedRows))
	}

	return nil
}

// InsertUser adds a new user into the DB.
//
// Parameters:
//   - ctx: Manage metadata and signals related to requests.
//   - userData: A struct containing the user information.
//
// Returns an error if the insertion fails
func (r *userRepo) InsertUser(
	ctx context.Context, userData models.User,
) error {
	cost := 10
	cryptedPswd, err := bcrypt.GenerateFromPassword(
		[]byte(userData.Password), cost)
	if err != nil {
		slog.Error(
			"Failed to hash user password before adding it to the DB",
			"error", err, "method", "InsertUser", "layer", layerName)
		return err
	}

	userData.Password = string(cryptedPswd)

	result, err := r.db.ExecContext(ctx, `
		INSERT INTO users (name, role, email, age, password) 
		VALUES (:name, :role, :email, :age, :password)`, userData)
	if err != nil {
		slog.Error(
			"Failed to insert user in the DB",
			"error", err, "method", "InsertUser", "layer", layerName)
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		slog.Error(
			"Failed to check rows affected at user insertion", "error", err,
			"method", "InsertUser", "layer", layerName)
		return err
	}

	if affectedRows < 1 {
		err := fmt.Errorf("user was not inserted in the DB")
		slog.Error(
			"Failed to insert user in the DB", "error", err,
			"method", "InsertUser", "layer", layerName)
		return err
	}

	return nil
}

// UpdateUserByID updates a user's information in the database.
//
// Parameters:
//   - ctx: Manage metadata and signals related to requests.
//   - userID: The unique identifier of the user to be updated
//   - userData: A struct containing the updated user information
//     to be applied
//
// Returns an error if the user is not found or the update fails
func (r *userRepo) UpdateUserByID(
	ctx context.Context, userID int, userData models.User,
) error {
	result, err := r.db.ExecContext(ctx, `
		UPDATE users SET name = :name, role = :role, 
		email = :email, age = :age, password = :password`, userData)
	if err != nil {
		slog.Error(
			fmt.Sprintf("Failed to update user with ID %d in the DB", userID),
			"error", err, "method", "UpdateUserByID", "layer", layerName)
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		slog.Error(
			"Failed to check rows affected at user update post-update",
			"error", err, "method", "UpdateUserByID", "layer", layerName)
		return err
	}

	if affectedRows < 1 {
		err := fmt.Errorf(
			"user with ID %d was not updated with success", userID)
		slog.Error(
			"Failed to update user in the DB", "error", err,
			"method", "UpdateUserByID", "layer", layerName)
		return err
	}

	return nil
}
