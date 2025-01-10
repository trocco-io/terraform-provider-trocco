package list

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.List = UniqueObjectAttributeValue{}

type UniqueObjectAttributeValue struct {
	AttributeName string
}

func (v UniqueObjectAttributeValue) Description(ctx context.Context) string {
	return "Ensures the value of the specified attribute of the object in the list is unique."
}

func (v UniqueObjectAttributeValue) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v UniqueObjectAttributeValue) ValidateList(
	ctx context.Context,
	req validator.ListRequest,
	resp *validator.ListResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var objects []types.Object
	if diags := req.ConfigValue.ElementsAs(ctx, &objects, false); diags.HasError() {
		resp.Diagnostics.Append(diags...)

		return
	}

	existingAttributeValues := []attr.Value{}
	for i, object := range objects {
		// Get the value of the attribute.
		attributeValue, ok := object.Attributes()[v.AttributeName]
		if !ok {
			// Ignore the object if the attribute does not exist.
			continue
		}

		// Ignore the value of the attribute if it is unknown or null.
		if attributeValue.IsUnknown() || attributeValue.IsNull() {
			continue
		}

		// Check if the attribute value is already in the list.
		for _, existingAttributeValue := range existingAttributeValues {
			if existingAttributeValue.Equal(attributeValue) {
				resp.Diagnostics.AddAttributeError(
					req.Path.AtListIndex(i),
					"Duplicated Value of Attribute of Object in List",
					fmt.Sprintf(
						"Attribute value %s of %s is duplicated",
						attributeValue,
						v.AttributeName,
					),
				)

				break
			}
		}

		existingAttributeValues = append(existingAttributeValues, attributeValue)
	}
}
