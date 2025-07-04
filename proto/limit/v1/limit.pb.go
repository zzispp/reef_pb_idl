// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: limit/v1/limit.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateLimitOrderRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ChainName      string                 `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`                          // 链名称
	UserAddress    string                 `protobuf:"bytes,2,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`                    // 用户地址
	TokenAddress   string                 `protobuf:"bytes,3,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"`                 // 代币地址
	Amount         float64                `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`                                               // 数量
	Direction      int32                  `protobuf:"varint,5,opt,name=direction,proto3" json:"direction,omitempty"`                                          // 方向 1: 买入 2: 卖出
	Price          *float64               `protobuf:"fixed64,6,opt,name=price,proto3,oneof" json:"price,omitempty"`                                           // 价格
	TokenInAmount  *float64               `protobuf:"fixed64,7,opt,name=token_in_amount,json=tokenInAmount,proto3,oneof" json:"token_in_amount,omitempty"`    // 输入量
	TokenOutAmount *float64               `protobuf:"fixed64,8,opt,name=token_out_amount,json=tokenOutAmount,proto3,oneof" json:"token_out_amount,omitempty"` // 输出量
	GasPrice       float64                `protobuf:"fixed64,9,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`                           //gas 价格
	Symbol         string                 `protobuf:"bytes,10,opt,name=symbol,proto3" json:"symbol,omitempty"`                                                // 交易符号 BNB/USDT
	PrivateKey     string                 `protobuf:"bytes,11,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`                      // 私钥
	Slippage       float64                `protobuf:"fixed64,12,opt,name=slippage,proto3" json:"slippage,omitempty"`                                          // 滑点
	BizId          string                 `protobuf:"bytes,13,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`                                     // 业务ID
	Percent        *float64               `protobuf:"fixed64,14,opt,name=percent,proto3,oneof" json:"percent,omitempty"`                                      // 百分比
	ExpireAt       uint64                 `protobuf:"varint,15,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`                           // 过期时间
	LimitType      int32                  `protobuf:"varint,16,opt,name=limit_type,json=limitType,proto3" json:"limit_type,omitempty"`                        // 限价类型 1: 按价格  2.按输入输出量 3.止盈止损
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateLimitOrderRequest) Reset() {
	*x = CreateLimitOrderRequest{}
	mi := &file_limit_v1_limit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLimitOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLimitOrderRequest) ProtoMessage() {}

func (x *CreateLimitOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLimitOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateLimitOrderRequest) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLimitOrderRequest) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetTokenInAmount() float64 {
	if x != nil && x.TokenInAmount != nil {
		return *x.TokenInAmount
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetTokenOutAmount() float64 {
	if x != nil && x.TokenOutAmount != nil {
		return *x.TokenOutAmount
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetGasPrice() float64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetSlippage() float64 {
	if x != nil {
		return x.Slippage
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *CreateLimitOrderRequest) GetPercent() float64 {
	if x != nil && x.Percent != nil {
		return *x.Percent
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetExpireAt() uint64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *CreateLimitOrderRequest) GetLimitType() int32 {
	if x != nil {
		return x.LimitType
	}
	return 0
}

type CreateLimitOrderReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // 订单ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLimitOrderReply) Reset() {
	*x = CreateLimitOrderReply{}
	mi := &file_limit_v1_limit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLimitOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLimitOrderReply) ProtoMessage() {}

func (x *CreateLimitOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLimitOrderReply.ProtoReflect.Descriptor instead.
func (*CreateLimitOrderReply) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLimitOrderReply) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CancelLimitOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // 订单ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelLimitOrderRequest) Reset() {
	*x = CancelLimitOrderRequest{}
	mi := &file_limit_v1_limit_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelLimitOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelLimitOrderRequest) ProtoMessage() {}

func (x *CancelLimitOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelLimitOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelLimitOrderRequest) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{2}
}

func (x *CancelLimitOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CancelLimitOrderReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelLimitOrderReply) Reset() {
	*x = CancelLimitOrderReply{}
	mi := &file_limit_v1_limit_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelLimitOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelLimitOrderReply) ProtoMessage() {}

func (x *CancelLimitOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelLimitOrderReply.ProtoReflect.Descriptor instead.
func (*CancelLimitOrderReply) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{3}
}

type GetLimitOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BizId         string                 `protobuf:"bytes,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"` // 业务ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLimitOrderRequest) Reset() {
	*x = GetLimitOrderRequest{}
	mi := &file_limit_v1_limit_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLimitOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLimitOrderRequest) ProtoMessage() {}

func (x *GetLimitOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLimitOrderRequest.ProtoReflect.Descriptor instead.
func (*GetLimitOrderRequest) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{4}
}

func (x *GetLimitOrderRequest) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

type LimitOrder struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                          // 订单ID
	CreatedAt      string                 `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                           // 创建时间
	UpdatedAt      string                 `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                           // 修改时间
	DeletedAt      string                 `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`                           // 删除时间
	ExpireAt       string                 `protobuf:"bytes,5,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`                              // 过期时间
	ChainName      string                 `protobuf:"bytes,6,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`                           // 链名称
	UserAddress    string                 `protobuf:"bytes,7,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`                     // 用户地址
	TokenAddress   string                 `protobuf:"bytes,8,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"`                  // 代币地址
	Amount         float64                `protobuf:"fixed64,9,opt,name=amount,proto3" json:"amount,omitempty"`                                                // 数量
	Direction      int32                  `protobuf:"varint,10,opt,name=direction,proto3" json:"direction,omitempty"`                                          // 方向 0:买入 1:卖出
	LimitType      int32                  `protobuf:"varint,11,opt,name=limit_type,json=limitType,proto3" json:"limit_type,omitempty"`                         // 限价类型 0: 按价格  1.按输入输出量 2.止盈止损
	Price          *float64               `protobuf:"fixed64,12,opt,name=price,proto3,oneof" json:"price,omitempty"`                                           // 价格
	TokenInAmount  *float64               `protobuf:"fixed64,13,opt,name=token_in_amount,json=tokenInAmount,proto3,oneof" json:"token_in_amount,omitempty"`    // 输入量
	TokenOutAmount *float64               `protobuf:"fixed64,14,opt,name=token_out_amount,json=tokenOutAmount,proto3,oneof" json:"token_out_amount,omitempty"` // 输出量
	GasPrice       float64                `protobuf:"fixed64,15,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`                           // gas价格 gwei/lamport
	Symbol         string                 `protobuf:"bytes,16,opt,name=symbol,proto3" json:"symbol,omitempty"`                                                 // 交易符号 BNB/USDT
	PrivateKey     string                 `protobuf:"bytes,17,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`                       // 私钥
	Slippage       float64                `protobuf:"fixed64,18,opt,name=slippage,proto3" json:"slippage,omitempty"`                                           // 滑点
	BizId          string                 `protobuf:"bytes,19,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`                                      // 业务ID
	Status         int32                  `protobuf:"varint,20,opt,name=status,proto3" json:"status,omitempty"`                                                // 状态
	Result         string                 `protobuf:"bytes,21,opt,name=result,proto3" json:"result,omitempty"`                                                 // 结果
	Percent        *float64               `protobuf:"fixed64,22,opt,name=percent,proto3,oneof" json:"percent,omitempty"`                                       // 百分比
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LimitOrder) Reset() {
	*x = LimitOrder{}
	mi := &file_limit_v1_limit_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LimitOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitOrder) ProtoMessage() {}

func (x *LimitOrder) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitOrder.ProtoReflect.Descriptor instead.
func (*LimitOrder) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{5}
}

func (x *LimitOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LimitOrder) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LimitOrder) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *LimitOrder) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *LimitOrder) GetExpireAt() string {
	if x != nil {
		return x.ExpireAt
	}
	return ""
}

func (x *LimitOrder) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *LimitOrder) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *LimitOrder) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *LimitOrder) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LimitOrder) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *LimitOrder) GetLimitType() int32 {
	if x != nil {
		return x.LimitType
	}
	return 0
}

func (x *LimitOrder) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *LimitOrder) GetTokenInAmount() float64 {
	if x != nil && x.TokenInAmount != nil {
		return *x.TokenInAmount
	}
	return 0
}

func (x *LimitOrder) GetTokenOutAmount() float64 {
	if x != nil && x.TokenOutAmount != nil {
		return *x.TokenOutAmount
	}
	return 0
}

func (x *LimitOrder) GetGasPrice() float64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *LimitOrder) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *LimitOrder) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *LimitOrder) GetSlippage() float64 {
	if x != nil {
		return x.Slippage
	}
	return 0
}

func (x *LimitOrder) GetBizId() string {
	if x != nil {
		return x.BizId
	}
	return ""
}

func (x *LimitOrder) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LimitOrder) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *LimitOrder) GetPercent() float64 {
	if x != nil && x.Percent != nil {
		return *x.Percent
	}
	return 0
}

type GetLimitOrderReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LimitOrders   *LimitOrder            `protobuf:"bytes,1,opt,name=limit_orders,json=limitOrders,proto3" json:"limit_orders,omitempty"` // 订单列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLimitOrderReply) Reset() {
	*x = GetLimitOrderReply{}
	mi := &file_limit_v1_limit_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLimitOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLimitOrderReply) ProtoMessage() {}

func (x *GetLimitOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLimitOrderReply.ProtoReflect.Descriptor instead.
func (*GetLimitOrderReply) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{6}
}

func (x *GetLimitOrderReply) GetLimitOrders() *LimitOrder {
	if x != nil {
		return x.LimitOrders
	}
	return nil
}

type GetLimitOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BizIds        []string               `protobuf:"bytes,1,rep,name=biz_ids,json=bizIds,proto3" json:"biz_ids,omitempty"` // 业务ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLimitOrdersRequest) Reset() {
	*x = GetLimitOrdersRequest{}
	mi := &file_limit_v1_limit_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLimitOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLimitOrdersRequest) ProtoMessage() {}

func (x *GetLimitOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLimitOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetLimitOrdersRequest) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{7}
}

func (x *GetLimitOrdersRequest) GetBizIds() []string {
	if x != nil {
		return x.BizIds
	}
	return nil
}

type GetLimitOrdersReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LimitOrders   []*LimitOrder          `protobuf:"bytes,1,rep,name=limit_orders,json=limitOrders,proto3" json:"limit_orders,omitempty"` // 订单列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLimitOrdersReply) Reset() {
	*x = GetLimitOrdersReply{}
	mi := &file_limit_v1_limit_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLimitOrdersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLimitOrdersReply) ProtoMessage() {}

func (x *GetLimitOrdersReply) ProtoReflect() protoreflect.Message {
	mi := &file_limit_v1_limit_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLimitOrdersReply.ProtoReflect.Descriptor instead.
func (*GetLimitOrdersReply) Descriptor() ([]byte, []int) {
	return file_limit_v1_limit_proto_rawDescGZIP(), []int{8}
}

func (x *GetLimitOrdersReply) GetLimitOrders() []*LimitOrder {
	if x != nil {
		return x.LimitOrders
	}
	return nil
}

var File_limit_v1_limit_proto protoreflect.FileDescriptor

var file_limit_v1_limit_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x22, 0xd0, 0x04, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f,
	0x75, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x02, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6c,
	0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x6c,
	0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x48, 0x03,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x17, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x22, 0xe0, 0x05, 0x0a, 0x0a, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x02, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73,
	0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69,
	0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d,
	0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x69, 0x7a, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x7a, 0x49, 0x64,
	0x73, 0x22, 0x52, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0xf8, 0x02, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x5e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x5e, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x27, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x15, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_limit_v1_limit_proto_rawDescOnce sync.Once
	file_limit_v1_limit_proto_rawDescData []byte
)

func file_limit_v1_limit_proto_rawDescGZIP() []byte {
	file_limit_v1_limit_proto_rawDescOnce.Do(func() {
		file_limit_v1_limit_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_limit_v1_limit_proto_rawDesc), len(file_limit_v1_limit_proto_rawDesc)))
	})
	return file_limit_v1_limit_proto_rawDescData
}

var file_limit_v1_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_limit_v1_limit_proto_goTypes = []any{
	(*CreateLimitOrderRequest)(nil), // 0: api.limit.v1.CreateLimitOrderRequest
	(*CreateLimitOrderReply)(nil),   // 1: api.limit.v1.CreateLimitOrderReply
	(*CancelLimitOrderRequest)(nil), // 2: api.limit.v1.CancelLimitOrderRequest
	(*CancelLimitOrderReply)(nil),   // 3: api.limit.v1.CancelLimitOrderReply
	(*GetLimitOrderRequest)(nil),    // 4: api.limit.v1.GetLimitOrderRequest
	(*LimitOrder)(nil),              // 5: api.limit.v1.LimitOrder
	(*GetLimitOrderReply)(nil),      // 6: api.limit.v1.GetLimitOrderReply
	(*GetLimitOrdersRequest)(nil),   // 7: api.limit.v1.GetLimitOrdersRequest
	(*GetLimitOrdersReply)(nil),     // 8: api.limit.v1.GetLimitOrdersReply
}
var file_limit_v1_limit_proto_depIdxs = []int32{
	5, // 0: api.limit.v1.GetLimitOrderReply.limit_orders:type_name -> api.limit.v1.LimitOrder
	5, // 1: api.limit.v1.GetLimitOrdersReply.limit_orders:type_name -> api.limit.v1.LimitOrder
	0, // 2: api.limit.v1.Limit.CreateLimitOrder:input_type -> api.limit.v1.CreateLimitOrderRequest
	2, // 3: api.limit.v1.Limit.CancelLimitOrder:input_type -> api.limit.v1.CancelLimitOrderRequest
	4, // 4: api.limit.v1.Limit.GetLimitOrder:input_type -> api.limit.v1.GetLimitOrderRequest
	7, // 5: api.limit.v1.Limit.GetLimitOrders:input_type -> api.limit.v1.GetLimitOrdersRequest
	1, // 6: api.limit.v1.Limit.CreateLimitOrder:output_type -> api.limit.v1.CreateLimitOrderReply
	3, // 7: api.limit.v1.Limit.CancelLimitOrder:output_type -> api.limit.v1.CancelLimitOrderReply
	6, // 8: api.limit.v1.Limit.GetLimitOrder:output_type -> api.limit.v1.GetLimitOrderReply
	8, // 9: api.limit.v1.Limit.GetLimitOrders:output_type -> api.limit.v1.GetLimitOrdersReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_limit_v1_limit_proto_init() }
func file_limit_v1_limit_proto_init() {
	if File_limit_v1_limit_proto != nil {
		return
	}
	file_limit_v1_limit_proto_msgTypes[0].OneofWrappers = []any{}
	file_limit_v1_limit_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_limit_v1_limit_proto_rawDesc), len(file_limit_v1_limit_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_limit_v1_limit_proto_goTypes,
		DependencyIndexes: file_limit_v1_limit_proto_depIdxs,
		MessageInfos:      file_limit_v1_limit_proto_msgTypes,
	}.Build()
	File_limit_v1_limit_proto = out.File
	file_limit_v1_limit_proto_goTypes = nil
	file_limit_v1_limit_proto_depIdxs = nil
}
