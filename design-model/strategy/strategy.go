package main

import "fmt"

/*
策略模式的示例代码:
	我们定义了一个 PaymentStrategy 接口，其中包含一个 Pay() 方法，
	以及两个具体的支付方式 CreditCardStrategy 和 AlipayStrategy，
	它们分别实现了 Pay() 方法。我们还定义了一个 PaymentContext 结构体，
	其中包含一个 PaymentStrategy 类型的属性，用于选择不同的支付方式。
	在实际使用时，我们可以通过设置 PaymentContext 的 strategy 属性
	来选择不同的支付方式，并调用 Pay() 方法进行支付操作。
*/

// PaymentStrategy 定义策略接口
type PaymentStrategy interface {
	Pay(amount float64) error
}

// CreditCardStrategy 实现策略接口的支付方式：信用卡支付
type CreditCardStrategy struct {
	Name         string
	CardNo       string
	ExpMonth     string
	ExpYear      string
	SecurityCode string
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	fmt.Printf("支付 %.2f 元（信用卡 %s）\n", amount, c.CardNo)
	return nil
}

// AlipayStrategy 实现策略接口的支付方式：支付宝支付
type AlipayStrategy struct {
	Account  string
	Password string
}

func (a *AlipayStrategy) Pay(amount float64) error {
	fmt.Printf("支付 %.2f 元（支付宝 %s）\n", amount, a.Account)
	return nil
}

// PaymentContext 定义上下文结构体，用于根据不同的支付方式选择不同的策略
type PaymentContext struct {
	strategy PaymentStrategy
}

// SetStrategy 设置上下文结构体的策略
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Pay 执行支付操作
func (p *PaymentContext) Pay(amount float64) error {
	return p.strategy.Pay(amount)
}

func main() {
	// 创建信用卡支付策略实例
	creditCard := &CreditCardStrategy{
		Name:         "张三",
		CardNo:       "1234 5678 9012 3456",
		ExpMonth:     "12",
		ExpYear:      "25",
		SecurityCode: "123",
	}

	// 创建支付宝支付策略实例
	alipay := &AlipayStrategy{
		Account:  "zhangsan@qq.com",
		Password: "123456",
	}

	// 创建上下文结构体实例
	context := &PaymentContext{}

	// 设置信用卡支付策略并执行支付
	context.SetStrategy(creditCard)
	context.Pay(100.00)

	// 设置支付宝支付策略并执行支付
	context.SetStrategy(alipay)
	context.Pay(50.00)
}
