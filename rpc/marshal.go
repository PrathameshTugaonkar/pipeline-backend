package rpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/instill-ai/pipeline-backend/pkg/model"

	pb "github.com/instill-ai/protogen-go/pipeline/v1alpha"
)

func marshalRecipeSource(d *model.Source) *pb.Source {
	return &pb.Source{
		Type: d.Type,
	}
}

func marshalRecipeDestination(d *model.Destination) *pb.Destination {
	return &pb.Destination{
		Type: d.Type,
	}
}

func marshalRecipeModel(v []*model.Model) []*pb.Model {
	var ret []*pb.Model
	for _, vv := range v {
		ret = append(ret, &pb.Model{
			Name:    vv.Name,
			Version: vv.Version,
		})
	}
	return ret
}

func marshalRecipe(recipe *model.Recipe) *pb.Recipe {
	return &pb.Recipe{
		Source:      marshalRecipeSource(recipe.Source),
		Destination: marshalRecipeDestination(recipe.Destination),
		Models:      marshalRecipeModel(recipe.Model),
	}
}

func marshalPipeline(pipeline *model.Pipeline) *pb.Pipeline {
	ret := &pb.Pipeline{
		Id:          pipeline.Id,
		Name:        pipeline.Name,
		Description: pipeline.Description,
		Active:      pipeline.Active,
		CreatedAt:   timestamppb.New(pipeline.CreatedAt),
		UpdatedAt:   timestamppb.New(pipeline.UpdatedAt),
		FullName:    pipeline.FullName,
	}

	if pipeline.Recipe != nil {
		ret.Recipe = marshalRecipe(pipeline.Recipe)
	}

	return ret
}
