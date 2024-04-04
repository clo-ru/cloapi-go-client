package recipe

import "github.com/clo-ru/cloapi-go-client/clo"

type LicenseRecipe struct {
	Addon    string `json:"addon"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

type Recipe struct {
	ID                string          `json:"id"`
	Name              string          `json:"name"`
	MinDisk           int             `json:"min_disk"`
	MinRam            int             `json:"min_ram"`
	MinVcpus          int             `json:"min_vcpus"`
	SuitableImages    []string        `json:"suitable_images"`
	AvailableLicenses []LicenseRecipe `json:"available_licenses"`
}

type RecipeListResponse = clo.ListResponse[Recipe]
