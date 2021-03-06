package serverRepo

import (
	"github.com/easysoft/zendata/src/model"
	constant "github.com/easysoft/zendata/src/utils/const"
	"github.com/jinzhu/gorm"
)

type RangesRepo struct {
	db *gorm.DB
}

func (r *RangesRepo) ListAll() (models []*model.ZdRanges) {
	r.db.Select("id,title,referName,fileName,folder,path,updatedAt").Find(&models)
	return
}

func (r *RangesRepo) List(keywords string, page int) (models []*model.ZdRanges, total int, err error) {
	query := r.db.Select("id,title,referName,fileName,folder,path").Order("id ASC")
	if keywords != "" {
		query = query.Where("title LIKE ?", "%"+keywords+"%")
	}
	if page > 0 {
		query = query.Offset((page - 1) * constant.PageSize).Limit(constant.PageSize)
	}

	err = query.Find(&models).Error

	err = r.db.Model(&model.ZdRanges{}).Count(&total).Error

	return
}

func (r *RangesRepo) Get(id uint) (ranges model.ZdRanges, err error) {
	err = r.db.Where("id=?", id).First(&ranges).Error
	return
}

func (r *RangesRepo) Create(ranges *model.ZdRanges) (err error) {
	err = r.db.Create(ranges).Error
	return
}
func (r *RangesRepo) Update(ranges *model.ZdRanges) (err error) {
	err = r.db.Save(ranges).Error
	return
}

func (r *RangesRepo) Remove(id uint) (err error) {
	ranges := model.ZdRanges{}
	ranges.ID = id

	err = r.db.Delete(ranges).Error
	err = r.db.Where("rangesID = ?", id).Delete(&model.ZdRangesItem{}).Error

	return
}

func (r *RangesRepo) GetItems(rangesId int) (items []*model.ZdRangesItem, err error) {
	err = r.db.Where("rangesId=?", rangesId).Find(&items).Error
	return
}
func (r *RangesRepo) GetItem(itemId uint) (item model.ZdRangesItem, err error) {
	err = r.db.Where("id=?", itemId).First(&item).Error
	return
}
func (r *RangesRepo) SaveItem(item *model.ZdRangesItem) (err error) {
	err = r.db.Save(item).Error
	return
}
func (r *RangesRepo) RemoveItem(id uint) (err error) {
	item := model.ZdRangesItem{}
	item.ID = id
	err = r.db.Delete(item).Error
	return
}
func (r *RangesRepo) GetMaxOrder(rangesId int) (ord int) {
	var preChild model.ZdField
	err := r.db.
		Where("rangesID=?", rangesId).
		Order("ord DESC").Limit(1).
		First(&preChild).Error

	if err != nil {
		ord = 1
	}
	ord = preChild.Ord + 1

	return
}

func (r *RangesRepo) UpdateYaml(po model.ZdRanges) (err error) {
	err = r.db.Model(&model.ZdRanges{}).Where("id=?", po.ID).Update("yaml", po.Yaml).Error
	return
}

func (r *RangesRepo) GenRangesRes(ranges model.ZdRanges, res *model.ResRanges) {
	res.Title = ranges.Title
	res.Desc = ranges.Desc
}

func (r *RangesRepo) UpdateItemRange(rang string, id uint) (err error) {
	err = r.db.Model(&model.ZdRangesItem{}).Where("id=?", id).Update("value", rang).Error

	return
}

func NewRangesRepo(db *gorm.DB) *RangesRepo {
	return &RangesRepo{db: db}
}
