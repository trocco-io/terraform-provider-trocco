package client

import (
	"fmt"
	"net/http"
	"net/url"
)

const teamBasePath = "/api/teams"

type Team struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Members     []TeamMember `json:"members"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type TeamWithoutMembers struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TeamMember struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

// List of Teams

type ListTeamsInput struct {
	limit  *int
	cursor *string
}

func (input *ListTeamsInput) SetLimit(limit int) {
	input.limit = &limit
}

func (input *ListTeamsInput) SetCursor(cursor string) {
	input.cursor = &cursor
}

type ListTeamsOutput struct {
	Items      []TeamWithoutMembers `json:"items"`
	NextCursor *string              `json:"next_cursor"`
}

const MaxListTeamsLimit = 100

func (client *TroccoClient) ListTeams(input *ListTeamsInput) (*ListTeamsOutput, error) {
	params := url.Values{}
	if input != nil && input.limit != nil {
		if *input.limit < 1 || *input.limit > MaxListTeamsLimit {
			return nil, fmt.Errorf("limit must be between 1 and %d", MaxListTeamsLimit)
		}
		params.Add("limit", fmt.Sprintf("%d", *input.limit))
	}
	if input != nil && input.cursor != nil {
		params.Add("cursor", *input.cursor)
	}
	path := fmt.Sprintf(teamBasePath+"?%s", params.Encode())
	output := new(ListTeamsOutput)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Get a Team

func (client *TroccoClient) GetTeam(id int64) (*Team, error) {
	path := fmt.Sprintf("%s/%d", teamBasePath, id)
	output := new(Team)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Create a Team

type CreateTeamInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Membmers    []struct {
		UserID int64  `json:"user_id"`
		Role   string `json:"role"`
	} `json:"members"`
}

func (client *TroccoClient) CreateTeam(input *CreateTeamInput) (*Team, error) {
	output := new(Team)
	err := client.do(http.MethodPost, teamBasePath, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Update a Team

type UpdateTeamInput struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Members     []struct {
		UserID int64  `json:"user_id"`
		Role   string `json:"role"`
	} `json:"members"`
}

func (client *TroccoClient) UpdateTeam(id int64, input *UpdateTeamInput) (*Team, error) {
	path := fmt.Sprintf("%s/%d", teamBasePath, id)
	output := new(Team)
	err := client.do(http.MethodPatch, path, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Delete a Team

func (client *TroccoClient) DeleteTeam(id int64) error {
	path := fmt.Sprintf("%s/%d", teamBasePath, id)
	return client.do(http.MethodDelete, path, nil, nil)
}
