//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost GetUHostUpgradePrice

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetUHostUpgradePriceRequest is request schema for GetUHostUpgradePrice action
type GetUHostUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 	UHost实例ID。 参见 [DescribeUHostInstance](describe_uhost_instance.html)。
	UHostId *string `required:"true"`

	// 虚拟CPU核数。可选参数：1-32（可选范围与UHostType相关）。默认值为当前实例的CPU核数。
	CPU *int `required:"false"`

	// 内存大小。单位：MB。范围 ：[1024, 262144]，取值为1024的倍数（可选范围与UHostType相关）。默认值为当前实例的内存大小。
	Memory *int `required:"false"`

	// 【待废弃】数据盘大小，单位: GB，范围[0,1000]，步长: 10， 默认值是该主机当前数据盘大小。
	DiskSpace *int `required:"false"`

	// 【待废弃】系统大小，单位: GB，范围[20,100]，步长: 10。
	BootDiskSpace *int `required:"false"`

	// 方舟机型。No，Yes。默认是No。
	TimemachineFeature *string `required:"false"`

	// 网卡升降级（1，表示升级，2表示降级，0表示不变）
	NetCapValue *int `required:"false"`

	// 【待废弃】主机系列，目前支持N1,N2
	HostType *string `required:"false"`
}

// GetUHostUpgradePriceResponse is response schema for GetUHostUpgradePrice action
type GetUHostUpgradePriceResponse struct {
	response.CommonBase

	// 规格调整差价。精确到小数点后2位。
	Price float64
}

// NewGetUHostUpgradePriceRequest will create request of GetUHostUpgradePrice action.
func (c *UHostClient) NewGetUHostUpgradePriceRequest() *GetUHostUpgradePriceRequest {
	req := &GetUHostUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetUHostUpgradePrice - 获取UHost实例升级配置的价格。可选配置范围请参考[[api:uhost-api:uhost_type|云主机机型说明]]。
func (c *UHostClient) GetUHostUpgradePrice(req *GetUHostUpgradePriceRequest) (*GetUHostUpgradePriceResponse, error) {
	var err error
	var res GetUHostUpgradePriceResponse

	err = c.Client.InvokeAction("GetUHostUpgradePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
