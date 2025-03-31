package notification_destination

import "terraform-provider-trocco/internal/client/parameter"

type EmailConfigInput struct {
	Email *parameter.NullableString `json:"email,omitempty"`
}
