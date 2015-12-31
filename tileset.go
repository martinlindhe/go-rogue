package rogue

import (
	"io/ioutil"

	"github.com/gobuild/log"
	"gopkg.in/yaml.v2"
)

type tilesetSpec struct {
	Props        tilesetProps `json:"props"`
	DeepWater    []int        `json:"deepwater"`
	Beach        []int        `json:"beach"`
	Grass        []int        `json:"grass"`
	ShallowWater []int        `json:"shallowwater"`
	Lava         []int        `json:"lava"`
	Dirt         []int        `json:"dirt"`
	Wall         []int        `json:"wall"`
}

type tilesetProps struct {
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	TileMap string `json:"tilemap"`
}

func parseGroundTilesetDefinition(defFileName string) (tilesetSpec, error) {

	var specs tilesetSpec

	data, err := ioutil.ReadFile(defFileName)
	if err != nil {
		return specs, err
	}

	err = yaml.Unmarshal(data, &specs)
	if err != nil {
		return specs, err
	}

	log.Infof("Read %s", defFileName)
	return specs, nil
}
