package model

type ArticleStoreInmemory struct {
	ArticleMap []Article
}

func NewArticleStoreInmemory() *ArticleStoreInmemory {
	return &ArticleStoreInmemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat website", Body: "Hallo ini bydi"},
		},
	}
}

func (store *ArticleStoreInmemory) All() []Article {
	return store.ArticleMap
}

func (store *ArticleStoreInmemory) Save(article *Article) error {
	// 1. calculate current length
	lastID := len(store.ArticleMap)

	// set article id
	article.ID = lastID + 1

	// push to article map slices
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}

// FIND ID
func (store *ArticleStoreInmemory) Find(id int) *Article {
	article := Article{}
	for _, item := range store.ArticleMap {
		if item.ID == id {
			article = item
		}
	}
	return &article
}

// UPDATE
func (store *ArticleStoreInmemory) Update(article *Article) error {
	for index, item := range store.ArticleMap {
		if item.ID == article.ID {
			store.ArticleMap[index] = *article
		}
	}

	return nil
}

// DELETE
func (store *ArticleStoreInmemory) Delete(article *Article) error {
	articles := []Article{}

	for _, item := range store.ArticleMap {
		if item.ID != article.ID {
			articles = append(articles, item)
		}
	}

	store.ArticleMap = articles
	return nil
}
