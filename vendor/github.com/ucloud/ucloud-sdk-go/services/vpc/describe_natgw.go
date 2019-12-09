// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeNATGWRequest is request schema for DescribeNATGW action
type DescribeNATGWRequest struct {
	request.CommonBase

	// [公共参数] 项目Id。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 数据分页值。默认为20
	Limit *int `required:"false"`

	// NAT网关Id。默认为该项目下所有NAT网关
	NATGWIds []string `required:"false"`

	// 数据偏移量。默认为0
	Offset *int `required:"false"`
}

// DescribeNATGWResponse is response schema for DescribeNATGW action
type DescribeNATGWResponse struct {
	response.CommonBase

	// 查到的NATGW信息列表
	DataSet []NatGatewayDataSet

	// 满足条件的实例的总数
	TotalCount int
}

// NewDescribeNATGWRequest will create request of DescribeNATGW action.
func (c *VPCClient) NewDescribeNATGWRequest() *DescribeNATGWRequest {
	req := &DescribeNATGWRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeNATGW - 获取NAT网关信息
func (c *VPCClient) DescribeNATGW(req *DescribeNATGWRequest) (*DescribeNATGWResponse, error) {
	var err error
	var res DescribeNATGWResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeNATGW", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
