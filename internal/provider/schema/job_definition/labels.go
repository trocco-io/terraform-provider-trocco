package job_definition

import "github.com/hashicorp/terraform-plugin-framework/resource/schema"

func LabelsSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"id": schema.Int64Attribute{
					Computed:            true,
					MarkdownDescription: "The ID of the label",
				},
				"name": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "The name of the label",
				},
			},
		},
		MarkdownDescription: "Labels to be attached to the job definition",
	}
}
