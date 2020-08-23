package strategy

import "testing"

func TestName(t *testing.T) {
	// 相对于被调用的支付策略 这里就是支付策略的客户端

	// 业务上下文
	ctx := &Context{
		PayType: "wechat_pay",
	}

	// 获取支付方式
	var instance PaymentInterface
	switch ctx.PayType {
	case ConstWechatPay:
		instance = &WechatPay{}
	case ConstAliPayWap:
		instance = &AliPayWap{}
	case ConstBankPay:
		instance = &BankPay{}
	default:
		panic("无效的支付方式")
	}

	// 支付
	instance.Pay(ctx)
}