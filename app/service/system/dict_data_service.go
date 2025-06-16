package systemservice

import (
	"encoding/json"

	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/redis"
)

type SysDictDataService struct{}

// 字典数据分页查询
func (*SysDictDataService) QueryDictDataList(params vo.DictPagingParam) (vo.PagingBackVo[vo.DictDataListVo], error) {
	var dictDataList []vo.DictDataListVo
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysDictData{}).
		Select("sys_dict_data.*", "t.type as dict_type", "t.node_type").
		Order("sort,id").
		Joins("LEFT JOIN sys_dict_type as t ON sys_dict_data.dict_id = t.id")
	if params.Id != nil {
		query.Where("dict_id = ?", params.Id)
	}
	if params.Label != "" {
		query.Where("sys_dict_data.label LIKE ?", "%"+params.Label+"%")
	}
	if params.Status != "" {
		query.Where("sys_dict_data.status = ?", params.Status)
	}
	result := query.
		Count(&count).
		Scopes(service.PagingScope(params.Page, params.Size)).
		Find(&dictDataList)
	return vo.PagingBackVo[vo.DictDataListVo]{Data: dictDataList, Count: int(count)}, result.Error
}

// 通过字段类型查询字典数据
func (*SysDictDataService) QueryDictDataListByType(dictType string) ([]vo.DictDataListVo, error) {
	var dictDataList []vo.DictDataListVo
	if cache, _ := DB.Redis.HGet(redis.SysDictKey, dictType).Result(); cache != "" {
		if err := json.Unmarshal([]byte(cache), &cache); err == nil {
			return dictDataList, nil
		}
	}
	result := DB.Gorm.Model(&sysmodel.SysDictData{}).
		Select("sys_dict_data.*", "t.type as dict_type", "t.node_type").
		Order("sort,id").
		Joins("LEFT JOIN sys_dict_type as t ON sys_dict_data.dict_id = t.id").
		Where("sys_dict_data.status = 1").
		Where("t.type = ?", dictType).
		Find(&dictDataList)
	if len(dictDataList) > 0 {
		dictStrs, _ := json.Marshal(&dictDataList)
		DB.Redis.HSet(redis.SysDictKey, dictType, dictStrs)
	}
	return dictDataList, result.Error
}

// 创建字典数据
func (*SysDictDataService) CreateDictData(dictData vo.CreateDictData) error {
	// 删除缓存
	if dictType, err := (&SysDictTypeService{}).QueryDictTypeById(dictData.DictId); err == nil {
		DB.Redis.HDel(redis.SysDictKey, dictType.Type)
	}
	return DB.Gorm.Model(&sysmodel.SysDictData{}).Create(&sysmodel.SysDictData{
		DictId:    dictData.DictId,
		Sort:      dictData.Sort,
		Label:     dictData.Label,
		Value:     dictData.Value,
		Type:      dictData.Type,
		CssClass:  dictData.CssClass,
		ListClass: dictData.ListClass,
		IsDefault: dictData.IsDefault,
		Status:    dictData.Status,
		Remark:    dictData.Remark,
	}).Error
}

// 更新字典数据
func (*SysDictDataService) UpdateDictData(dictData vo.UpdateDictData) error {
	// 删除缓存
	if dictType, err := (&SysDictTypeService{}).QueryDictTypeById(dictData.DictId); err == nil {
		DB.Redis.HDel(redis.SysDictKey, dictType.Type)
	}
	return DB.Gorm.Model(&sysmodel.SysDictData{}).
		Scopes(service.UpdateOmitScope()).
		Where("id = ?", dictData.Id).
		Updates(&sysmodel.SysDictData{
			DictId:    dictData.DictId,
			Sort:      dictData.Sort,
			Label:     dictData.Label,
			Value:     dictData.Value,
			Type:      dictData.Type,
			CssClass:  dictData.CssClass,
			ListClass: dictData.ListClass,
			IsDefault: dictData.IsDefault,
			Status:    dictData.Status,
			Remark:    dictData.Remark,
		}).Error
}

// 删除字典数据
func (*SysDictDataService) DeleteDictData(ids []int) error {
	// 删除缓存
	for _, id := range ids {
		if dictType, err := (&SysDictTypeService{}).QueryDictTypeById(id); err == nil {
			DB.Redis.HDel(redis.SysDictKey, dictType.Type)
		}
	}
	return DB.Gorm.Model(&sysmodel.SysDictData{}).Delete(&sysmodel.SysDictData{}, ids).Error
}

// 校验字典名称
func (*SysDictDataService) IsHasSameName(label string, dictId int, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysDictData{}).
		Where("label = ?", label).
		Where("dict_id = ?", dictId)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}

// 校验字典值
func (*SysDictDataService) IsHasSameValue(value string, dictId int, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysDictData{}).
		Where("value = ?", value).
		Where("dict_id = ?", dictId)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}
