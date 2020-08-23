package strategy

import (
	"fmt"
	"runtime"
)

//------------------------------------------------------------
//我的代码没有`else`系列
//策略模式
//@auhtor TIGERB<https://github.com/TIGERB>
//------------------------------------------------------------

const (
	// ConstWechatPay 微信支付
	ConstWechatPay = "wechat_pay"
	// ConstAliPayWap 支付宝支付 网页版
	ConstAliPayWap = "AliPayWapwap"
	// ConstBankPay 银行卡支付
	ConstBankPay = "quickbank"
)

// Context 上下文
type Context struct {
	// 用户选择的支付方式
	PayType string `json:"pay_type"`
}

// PaymentInterface 支付方式接口
type PaymentInterface interface {
	Pay(ctx *Context) error    // 支付
	Refund(ctx *Context) error // 退款
}

// WechatPay 微信支付
type WechatPay struct {
}

// Pay 当前支付方式的支付逻辑
func (p *WechatPay) Pay(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用微信支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (p *WechatPay) Refund(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用微信退款...")
	return
}

// AliPayWap 支付宝网页版
type AliPayWap struct {
}

// Pay 当前支付方式的支付逻辑
func (p *AliPayWap) Pay(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用支付宝网页版支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (p *AliPayWap) Refund(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用支付宝网页版退款...")
	return
}

// BankPay 银行卡支付
type BankPay struct {
}

// Pay 当前支付方式的支付逻辑
func (p *BankPay) Pay(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用银行卡支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (p *BankPay) Refund(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用银行卡退款...")
	return
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
