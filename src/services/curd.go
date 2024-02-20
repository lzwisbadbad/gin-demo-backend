package services

import (
	"github.com/gin-backend/src/db"

	"gorm.io/gorm"
)

const DEFAULT_BATCHES_SIZE = 1000

func (s *Server) GetGormObject() *gorm.DB {
	return s.gormDb
}

func (s *Server) InsertOneObjertToDB(object db.ModelStruct) error {
	if err := s.gormDb.Create(object).Error; err != nil {
		s.sulog.Infof("insert one object to db failed, err:[%s], object: [%+v]\n",
			err.Error(), object)
		return err
	}
	return nil
}

func (s *Server) InsertObjectsToDB(objects interface{}) error {
	if err := s.gormDb.CreateInBatches(objects, DEFAULT_BATCHES_SIZE).Error; err != nil {
		s.sulog.Infof("insert objects to db failed, err:[%s]\n",
			err.Error())
		return err
	}
	return nil
}

func (s *Server) QueryObjectById(modelStruct db.ModelStruct,
	id int32) error {

	if err := s.gormDb.Model(modelStruct).
		Where("id = ?", id).First(modelStruct).Error; err != nil {

		s.sulog.Infof("query object by id failed, err:[%s], object:[%+v]\n",
			err.Error(), modelStruct)
		return err
	}
	return nil
}

func (s *Server) QueryObjectByCondition(modelStruct db.ModelStruct,
	searchIndex, searchInput string) error {

	if err := s.gormDb.Model(modelStruct).
		Where(searchIndex+" = ?", searchInput).First(modelStruct).Error; err != nil {

		s.sulog.Infof("query object by [%s] failed, err:[%s], object:[%+v]\n",
			searchIndex, err.Error(), modelStruct)
		return err
	}
	return nil
}
