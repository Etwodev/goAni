package gelbooru

const (
	GELBOORU_IMAGE_SEARCH string = `https://www.gelbooru.com/index.php?page=dapi&s=post&q=index&json=1&tags=%s&pid=%d`
	GELBOORU_IMAGE_STORE string = `./store/%s/`
	GELBOORU_IMAGE_SIZE int = 256
)
