package client

import (
	"fmt"
	"net/http"
	"net/url"
)

const BasePath = "/api/users"

type User struct {
	ID                           int64  `json:"id"`
	Email                        string `json:"email"`
	Role                         string `json:"role"`
	CanUseAuditLog               bool   `json:"can_use_audit_log"`
	IsRestrictedConnectionModify bool   `json:"is_restricted_connection_modify"`
	LastSignInAt                 string `json:"last_sign_in_at"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

// List of Users

type ListUsersInput struct {
	limit  *int
	cursor *string
}

func (input *ListUsersInput) SetLimit(limit int) {
	input.limit = &limit
}

func (input *ListUsersInput) SetCursor(cursor string) {
	input.cursor = &cursor
}

type ListUsersOutput struct {
	Items      []User  `json:"items"`
	NextCursor *string `json:"next_cursor"`
}

const MaxListUsersLimit = 100

func (client *TroccoClient) ListUsers(input *ListUsersInput) (*ListUsersOutput, error) {
	params := url.Values{}
	if input != nil && input.limit != nil {
		if *input.limit < 1 || *input.limit > MaxListUsersLimit {
			return nil, fmt.Errorf("limit must be between 1 and %d", MaxListUsersLimit)
		}
		params.Add("limit", fmt.Sprintf("%d", *input.limit))
	}
	if input != nil && input.cursor != nil {
		params.Add("cursor", *input.cursor)
	}
	path := fmt.Sprintf(BasePath+"?%s", params.Encode())
	output := new(ListUsersOutput)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Get a User

type GetUserOutput struct {
	ID                           int64  `json:"id"`
	Email                        string `json:"email"`
	Role                         string `json:"role"`
	CanUseAuditLog               bool   `json:"can_use_audit_log"`
	IsRestrictedConnectionModify bool   `json:"is_restricted_connection_modify"`
	LastSignInAt                 string `json:"last_sign_in_at"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

func (client *TroccoClient) GetUser(id int64) (*GetUserOutput, error) {
	path := fmt.Sprintf(BasePath+"/%d", id)
	output := new(GetUserOutput)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Create a User

type CreateUserInput struct {
	Email                        string `json:"email"`
	Password                     string `json:"password"`
	Role                         string `json:"role"`
	CanUseAuditLog               bool   `json:"can_use_audit_log"`
	IsRestrictedConnectionModify bool   `json:"is_restricted_connection_modify"`
}

type CreateUserOutput struct {
	ID                           int64  `json:"id"`
	Email                        string `json:"email"`
	Role                         string `json:"role"`
	CanUseAuditLog               bool   `json:"can_use_audit_log"`
	IsRestrictedConnectionModify bool   `json:"is_restricted_connection_modify"`
	LastSignInAt                 string `json:"last_sign_in_at"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

func (client *TroccoClient) CreateUser(input *CreateUserInput) (*CreateUserOutput, error) {
	output := new(CreateUserOutput)
	err := client.do(http.MethodPost, BasePath, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Update a User

type UpdateUserInput struct {
	Role                         string `json:"role"`
	CanUseAuditLog               bool   `json:"can_use_audit_log"`
	IsRestrictedConnectionModify bool   `json:"is_restricted_connection_modify"`
}

type UpdateUserOutput struct {
	ID                           int64  `json:"id"`
	Email                        string `json:"email"`
	Role                         string `json:"role"`
	CanUseAuditLog               bool   `json:"can_use_audit_log"`
	IsRestrictedConnectionModify bool   `json:"is_restricted_connection_modify"`
	LastSignInAt                 string `json:"last_sign_in_at"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

func (client *TroccoClient) UpdateUser(id int64, input *UpdateUserInput) (*UpdateUserOutput, error) {
	path := fmt.Sprintf(BasePath+"/%d", id)
	output := new(UpdateUserOutput)
	err := client.do(http.MethodPatch, path, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Delete a User

func (client *TroccoClient) DeleteUser(id int64) error {
	path := fmt.Sprintf(BasePath+"/%d", id)
	return client.do(http.MethodDelete, path, nil, nil)
}
