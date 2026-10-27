package pipeline_definition

import (
	"context"
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
)

// Not in alphabetical order, so that sorting by a set would be detected.
var unsortedVariables = []string{"$order_number_list$", "$date$", "$number$"}

func TestCustomVariableLoopConfigVariablesKeepOrder(t *testing.T) {
	ctx := context.Background()

	t.Run("bigquery", func(t *testing.T) {
		m := NewBigqueryCustomVariableLoopConfig(ctx, &pipelineDefinitionEntities.BigqueryCustomVariableLoopConfig{Variables: unsortedVariables})
		assertVariables(t, ctx, m.Variables, m.ToInput(ctx).Variables)
	})

	t.Run("redshift", func(t *testing.T) {
		m := NewRedshiftCustomVariableLoopConfig(ctx, &pipelineDefinitionEntities.RedshiftCustomVariableLoopConfig{Variables: unsortedVariables})
		assertVariables(t, ctx, m.Variables, m.ToInput(ctx).Variables)
	})

	t.Run("snowflake", func(t *testing.T) {
		m := NewSnowflakeCustomVariableLoopConfig(ctx, &pipelineDefinitionEntities.SnowflakeCustomVariableLoopConfig{Variables: unsortedVariables})
		assertVariables(t, ctx, m.Variables, m.ToInput(ctx).Variables)
	})
}

func assertVariables(t *testing.T, ctx context.Context, model types.List, input []string) {
	t.Helper()

	var fromModel []string
	if diags := model.ElementsAs(ctx, &fromModel, false); diags.HasError() {
		t.Fatalf("failed to read variables from model: %v", diags)
	}
	if !slices.Equal(fromModel, unsortedVariables) {
		t.Errorf("model variables = %v, want %v", fromModel, unsortedVariables)
	}
	if !slices.Equal(input, unsortedVariables) {
		t.Errorf("input variables = %v, want %v", input, unsortedVariables)
	}
}
