package articles

import (
	"errors"
	"strings"
)

//Article stores article
type Article struct {
	ArticleID   string `json:"article_id" db:"article_id"`
	Title       string `json:"title" db:"title"`
	Content     string `json:"content" db:"content"`
	IsPublished bool   `json:"flag" db:"flag"`
}

//ArticleForm stores article from html forms
type ArticleForm struct {
	Title       []string
	Content     []string
	IsPublished []string
}

func (artF *ArticleForm) convertToArticle() (Article, error) {
	art := Article{}
	art.ArticleID = "0"

	if len(artF.Title) == 0 || strings.Join(artF.Title, "") == "" {
		return Article{}, errors.New("Title cannot be empty")
	}
	art.Title = strings.Join(artF.Title, "")

	if len(artF.Content) == 0 || strings.Join(artF.Content, "") == "" {
		return Article{}, errors.New("Content cannot be empty")
	}
	art.Content = strings.Join(artF.Content, "")

	if strings.Join(artF.IsPublished, "") == "" {
		art.IsPublished = false
	} else {
		art.IsPublished = true
	}

	return art, nil
}
