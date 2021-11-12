package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://das.cn-south-1.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://das.cn-south-2.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://das.cn-north-1.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://das.cn-north-4.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://das.cn-east-3.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://das.cn-east-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-south-1": CN_SOUTH_1,
	"cn-south-2": CN_SOUTH_2,
	"cn-north-1": CN_NORTH_1,
	"cn-north-4": CN_NORTH_4,
	"cn-east-3":  CN_EAST_3,
	"cn-east-2":  CN_EAST_2,
}

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}
	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}