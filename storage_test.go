package gourmand

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func newTestFirebaseStorage() *firebaseStorage {
	s := NewFirebaseStorage()

	s.uri = map[string]string{
		"menu": "test-menus",
	}

	return s
}

func TestStorageFindMenu(t *testing.T) {
	s := newTestFirebaseStorage()

	var menu *Menu

	menu, err := s.FindMenu("test")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(menu.Items))
	assert.Equal(t, "Test", menu.Items[0].Name)

	menu, err = s.FindMenu("inexistent")
	assert.Nil(t, err)
	assert.Nil(t, menu)
}

func TestStorageSaveMenu(t *testing.T) {
	s := newTestFirebaseStorage()
	now := time.Now()
	menu := &Menu{
		Date:      "foo",
		Items:     []*MenuItem{},
		CrawledAt: now,
	}

	err := s.SaveMenu(menu)
	assert.Nil(t, err)

	newMenu, err := s.FindMenu("foo")
	assert.Nil(t, err)
	assert.Equal(t, menu.Date, newMenu.Date)
}
