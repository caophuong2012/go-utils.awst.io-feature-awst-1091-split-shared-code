package utils

import (
	"context"
	"encoding/json"
)

func MapStructure(ctx context.Context, input interface{}, output interface{}) error {
	jsonbody, err := json.Marshal(input)
	if err != nil {
		LogError(ctx, err).Send()
		return err
	}

	if err := json.Unmarshal(jsonbody, &output); err != nil {
		LogError(ctx, err).Send()
		return err
	}

	return nil
}
