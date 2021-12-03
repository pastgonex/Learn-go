package store

//
//import (
//	mystore "Learn_go/tonybai/hellogo/bookstore/store"
//	factory "Learn_go/tonybai/hellogo/bookstore/store/factory"
//	"sync"
//)
//
//func init() {
//	factory.Register("mem", &MemStore{
//		books: make(map[string]*mystore.Book),
//	})
//}
//
//type MemStore struct {
//	sync.RWMutex
//	books map[string]*mystore.Book
//}
//
//// Create creates a new Book in the store.
//func (ms *MemStore) Create(book *mystore.Book) error {
//	ms.Lock()
//	defer ms.Unlock()
//	if _, ok := ms.books[book.Id]; ok {
//		return mystore.ErrExist
//	}
//
//	nBook := *book
//	ms.books[book.Id] = &nBook
//
//	return nil
//}
//
//// Update updates the existed Book in the store.
//func (ms *MemStore) Update(book *mystore.Book) error {
//	ms.Lock()
//	defer ms.Unlock()
//
//	oldBook, ok := ms.books[book.Id]
//	if !ok {
//		return mystore.ErrNotFound
//	}
//
//	nBook := *oldBook
//	if book.Name != "" {
//		nBook.Name = book.Name
//	}
//
//	if book.Authors != nil {
//		nBook.Authors = book.Authors
//	}
//
//	if book.Press != "" {
//		nBook.Press = book.Press
//	}
//
//	ms.books[book.Id] = &nBook
//	return nil
//}
//
//// Get retrieves a book from the store, by id. If no such id exists. an
//// error is returned.
//func (ms *MemStore) Get(id string) (mystore.Book, error) {
//	ms.RLock()
//	defer ms.RUnlock()
//
//	t, ok := ms.books[id]
//	if ok {
//		return *t, nil
//	}
//	return mystore.Book{}, mystore.ErrNotFound
//}
//
//// Delete deletes the book with the given id. If no such id exists, an error
//// is returned.
//func (ms *MemStore) Delete
