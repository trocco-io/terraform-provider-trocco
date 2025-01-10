package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.List = &UniqueObjectAttributeValue{}

type UniqueObjectAttributeValue struct {
	AttributeName string
}

func (m UniqueObjectAttributeValue) Description(ctx context.Context) string {
	return "Ensures the value of the specified attribute of the object in the list is unique."
}

func (m UniqueObjectAttributeValue) MarkdownDescription(ctx context.Context) string {
	return "Ensures the value of the specified attribute of the object in the list is unique."
}

func (m UniqueObjectAttributeValue) PlanModifyList(
	ctx context.Context,
	req planmodifier.ListRequest,
	resp *planmodifier.ListResponse,
) {
	var objects types.List
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path, &objects)...)
	if resp.Diagnostics.HasError() {
		return
	}

	existingAttributeValues := []attr.Value{}
	for i, object := range objects.Elements() {
		object, ok := object.(types.Object)
		if !ok {
			continue
		}

		// Get the attribute value of the object.
		attributeValue := object.Attributes()[m.AttributeName]

		// Ignore unknown or null attribute values.
		if attributeValue.IsUnknown() || attributeValue.IsNull() {
			continue
		}

		// Check the attribute value of the object is already in the list.
		for _, existingAttributeValue := range existingAttributeValues {
			if existingAttributeValue.Equal(attributeValue) {
				resp.Diagnostics.AddAttributeError(
					req.Path.AtListIndex(i),
					"Duplicated Value of Attribute of Object in List",
					fmt.Sprintf(
						"Attribute value %s of %s is duplicated",
						attributeValue,
						m.AttributeName,
					),
				)

				break
			}
		}

		existingAttributeValues = append(existingAttributeValues, attributeValue)
	}
}
