package dynamodb

import (
	"fmt"
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ParseQueryOutput(result *dynamodb.QueryOutput) (*[]models.DBItem, error) {
	items := make([]models.DBItem, 0)

	if result == nil || len(result.Items) == 0 {
		return &items, nil
	}

	for _, item := range result.Items {
		pkAttr := item[config.PK]
		skAttr := item[config.SK]
		authorAttr := item[config.Author]
		categoryAttr := item[config.Category]
		creationDateAttr := item[config.CreationDate]
		fileAttr := item[config.File]
		nameAttr := item[config.Name]
		tagsAttr := item[config.Tags]

		if pkAttr == nil || skAttr == nil || authorAttr == nil || categoryAttr == nil ||
			creationDateAttr == nil || fileAttr == nil || nameAttr == nil || tagsAttr == nil {
			return nil, fmt.Errorf("missing required attributes in item")
		}

		item := models.DBItem{
			PK:           *pkAttr.S,
			SK:           *skAttr.S,
			Author:       *authorAttr.S,
			Category:     *categoryAttr.S,
			CreationDate: *creationDateAttr.S,
			File:         *fileAttr.S,
			Name:         *nameAttr.S,
			Tags:         *tagsAttr.S,
		}

		items = append(items, item)
	}

	return &items, nil
}
