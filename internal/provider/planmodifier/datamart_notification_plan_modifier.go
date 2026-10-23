package planmodifier

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func addNotificationAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Notification Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
