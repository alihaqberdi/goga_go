package repo

import "gorm.io/gorm"

type Probs struct {
	DB *gorm.DB
}

//
//func (r *Probs) Create(m *models.Prob) (ObjID, error) {
//	defer esUpdate(m)
//	updating(m)
//
//	return r.crud.Create(m)
//}
//
//func (r *Probs) Update(m *models.Prob) error {
//	defer esUpdate(m)
//	updating(m)
//
//	return r.crud.Update(m, m.GetObjID())
//}
//
//func (r *Probs) Find(probId, course string) (*models.Prob, error) {
//	m := new(models.Prob)
//
//	err := r.crud.getColl().
//		FindOne(ctx, bson.M{"probId": probId, "course": course}).
//		Decode(m)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return m, nil
//}
//
//func (r *Probs) FindList(data *dtos.ListOptions) ([]models.Prob, error) {
//	coll := r.crud.getColl()
//
//	findOptions := options.Find().
//		SetLimit(int64(data.Limit)).
//		SetSkip(int64(data.Offset))
//
//	filter := bson.M{}
//	for key, val := range data.Filter {
//		filter[key] = val
//	}
//
//	cursor, err := coll.Find(ctx, filter, findOptions)
//	if err != nil {
//		return nil, err
//	}
//
//	list, err := InitList[models.Prob](cursor)
//	if err != nil {
//		return nil, err
//	}
//
//	return list, nil
//}
//
//func updating(m *models.Prob) {
//	m.UpdatedAt = time.Now().UTC()
//}
//
//func esUpdate(m *models.Prob) {
//
//	if config.ES_AUTO_UPDATE {
//		es_repo.GoSafeUpsert(m)
//	}
//
//}
