package dynamodb

import (
	"fmt"
	"learn-api-blitzbudget-com/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ParseQueryOutput(result *dynamodb.QueryOutput) (*[]models.DBItem, error) {
	items := make([]models.DBItem, 0)

	if result == nil || len(result.Items) == 0 {
		return &items, nil
	}

	for _, item := range result.Items {
		pkAttr := item["pk"]
		skAttr := item["sk"]
		authorAttr := item["Author"]
		categoryAttr := item["Category"]
		creationDateAttr := item["creation_date"]
		fileAttr := item["File"]
		nameAttr := item["Name"]
		tagsAttr := item["Tags"]

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
