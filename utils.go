package main

import (
	"fmt"
	"math"
)

func intMax(i1 int, i2 int) int {
	return int(math.Max(float64(i1), float64(i2)))
}

func maxInt(ints []int) int {
	if len(ints) == 0 {
		panic("oops maxInt with []")
	}
	max := ints[0]
	for _, i := range ints {
		max = intMax(max, i)
	}
	return max
}

func getLevelOrder(repos []Repository) ([][]Repository, error) {
	idToRepo := make(map[string]Repository)
	for _, repo := range repos {
		idToRepo[repo.ID] = repo
	}

	idToLevel := make(map[string]int)
	maxLevel := 0
	var assignOrder func(Repository) (int, error)
	assignOrder = func(repo Repository) (int, error) {
		if _, levelAssigned := idToLevel[repo.ID]; !levelAssigned {
			if len(repo.Dependencies) == 0 {
				idToLevel[repo.ID] = 0
			} else {
				var dependencyLevels []int
				for _, dependencyID := range repo.Dependencies {
					dependencyRepo, dependencyExists := idToRepo[dependencyID]
					if !dependencyExists {
						return -1, fmt.Errorf("dependency %v does not exist", dependencyID)
					}
					dependencyLevel, err := assignOrder(dependencyRepo)
					if err != nil {
						return -1, err
					}
					dependencyLevels = append(dependencyLevels, dependencyLevel)
				}
				idToLevel[repo.ID] = maxInt(dependencyLevels) + 1
			}
		}
		maxLevel = intMax(maxLevel, idToLevel[repo.ID])
		return idToLevel[repo.ID], nil
	}
	for _, repo := range repos {
		_, err := assignOrder(repo)
		if err != nil {
			return nil, err
		}
	}

	levelOrder := make([][]Repository, maxLevel+1)
	for id, level := range idToLevel {
		if repo, exists := idToRepo[id]; exists {
			levelOrder[level] = append(levelOrder[level], repo)
		} else {
			return nil, fmt.Errorf("%v does not exist", id)
		}
	}

	return levelOrder, nil
}
