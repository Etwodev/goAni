package gelbooru

import (
	"fmt"
	"os"
	"path"

	"github.com/Etwodev/goAni/pkg/api"
	"github.com/Etwodev/goAni/pkg/gelbooru/config"
	"github.com/Etwodev/goAni/pkg/img"
)

func Get(group string, q string, pages int) error {
	var data GelbooruSearch
	dir := fmt.Sprintf(GELBOORU_IMAGE_STORE, group)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("Get: failed to create directory: %w", err)
	}

	for i := 0; i < pages; i++ {
		url := fmt.Sprintf(GELBOORU_IMAGE_SEARCH, q, i)

		err := api.Get(url, &data)
		if err != nil {
			return fmt.Errorf("Get: failed to get search payload: %w", err)
		}

		if len(data.Posts) == 0 {
			return nil
		}

		for _, post := range data.Posts {
			if config.Exists(group, post.Identifier) {
				continue
			}

			ext := path.Ext(post.Preview)
			if ext != ".jpg" {
				continue
			}

			path := fmt.Sprintf("%s%d%s", dir, post.Identifier, ext)

			bin, err := api.GetRaw(post.Preview)
			if err != nil {
				return fmt.Errorf("Get: failed to get image data: %w", err)
			}

			err = img.Resize(bin, path, GELBOORU_IMAGE_SIZE)
			if err != nil {
				return fmt.Errorf("Get: failed to resize image: %w", err)
			}

			config.Add(group, post.Identifier)
		}
	}
	return nil
}
