//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC UpdateVPCNetwork

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UpdateVPCNetworkRequest is request schema for UpdateVPCNetwork action
type UpdateVPCNetworkRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// VPC的ID
	VPCId *string `required:"true"`

	// 更新的全量网段
	Network []string `required:"true"`
}

// UpdateVPCNetworkResponse is response schema for UpdateVPCNetwork action
type UpdateVPCNetworkResponse struct {
	response.CommonBase

	// 错误信息
	Message string
}

// NewUpdateVPCNetworkRequest will create request of UpdateVPCNetwork action.
func (c *VPCClient) NewUpdateVPCNetworkRequest() *UpdateVPCNetworkRequest {
	req := &UpdateVPCNetworkRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateVPCNetwork - 修改VPC地址空间，只支持删除地址空间
func (c *VPCClient) UpdateVPCNetwork(req *UpdateVPCNetworkRequest) (*UpdateVPCNetworkResponse, error) {
	var err error
	var res UpdateVPCNetworkResponse

	err = c.Client.InvokeAction("UpdateVPCNetwork", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
