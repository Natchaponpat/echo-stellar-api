package storage

import (
	"github.com/Natchaponpat/echo-stellar-api/echo/mongo/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserStorage struct {
	conn       *mgo.Session
	db         string
	collection string
}

func NewUserStorage(conn *mgo.Session, db string, collection string) *UserStorage {
	return &UserStorage{
		conn:       conn,
		db:         db,
		collection: collection,
	}
}

func (s *UserStorage) List() ([]model.User, error) {
	c := s.conn.Copy()
	defer c.Close()

	var res []model.User
	if err := c.DB(s.db).C(s.collection).Find(nil).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserStorage) Get(name string) (model.User, error) {
	c := s.conn.Copy()
	defer c.Close()

	var res model.User
	query := bson.M{
		"name": name,
	}
	if err := c.DB(s.db).C(s.collection).Find(query).One(&res); err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (s *UserStorage) Create(user model.User) error {
	c := s.conn.Copy()
	defer c.Close()

	col := c.DB(s.db).C(s.collection)
	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	if err := col.EnsureIndex(index); err != nil {
		return err
	}

	if err := col.Insert(user); err != nil {
		return err
	}
	return nil
}
