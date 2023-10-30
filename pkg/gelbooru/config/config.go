package config

import (
	"fmt"
	"os"
	"time"

	"github.com/Etwodev/goAni/pkg/img"
	"go.mongodb.org/mongo-driver/bson"
)

var c *Config

func load() error {
	_, err := os.Stat(CONFIG_FILE_PATH)
	if os.IsNotExist(err) {
		if err := write(&Config{CreatedAt: time.Now().String(), UpdatedAt: time.Now().String()}); err != nil {
			return fmt.Errorf("load: failed creating config file: %w", err)
		}
	}

	file, err := os.ReadFile(CONFIG_FILE_PATH)
	if err != nil {
		return fmt.Errorf("load: failed reading config file: %w", err)
	}

	err = bson.Unmarshal(file, &c)
	if err != nil {
		return fmt.Errorf("load: failed to marshal config file: %w", err)
	}

	if c.Files == nil {
		c.Files = make(map[string][]int)
	}

	return nil
}

func write(config *Config) error {
	file, err := bson.Marshal(config)
	if err != nil {
		return fmt.Errorf("write: failed to marshal config file: %w", err)
	}

	err = os.WriteFile(CONFIG_FILE_PATH, file, 0644)
	if err != nil {
		return fmt.Errorf("write: failed writing config file: %w", err)
	}
	return nil
}

// Add a new id to child list to the config
func Add(id string, child ...int) error {
	if c == nil {
		err := load()
		if err != nil {
			return fmt.Errorf("Add: failed loading config file: %w", err)
		}
	}

	_, ok := c.Files[id]
	if ok {
		c.Files[id] = append(c.Files[id], child...)
	} else {
		c.Files[id] = child
	}

	c.UpdatedAt = time.Now().String()

	err := write(c)
	if err != nil {
		return fmt.Errorf("Add: failed writing config: %w", err)
	}

	return nil
}

// Checks if a child exists under an id
func Exists(id string, child int) bool {
	if c == nil {
		return false
	}

	val, ok := c.Files[id]
	if ok {
		for _, n := range val {
			if n == child {
				return true
			}
		}
	}
	return false
}

func GetDataset() ([]float64, []float64, int, int, error) {
	if c == nil {
		err := load()
		if err != nil {
			return nil, nil, 0, 0, fmt.Errorf("Add: failed loading dataset: %w", err)
		}
	}

	var inputs []float64
	var labels []float64
	size := len(c.Files)

	i := 0
	j := 0
	for k, v := range c.Files {
		for _, id := range v {
			out, err := img.ImageToArray(fmt.Sprintf("./store/%s/%d.jpg", k, id))
			if err != nil {
				panic(err)
			}

			inputs = append(inputs, out...)
			keys := make([]float64, size)
			keys[i] = 1
			labels = append(labels, keys...)
			j++
		}
		i++
	}

	return inputs, labels, size, j, nil
}

// red, green, blue
