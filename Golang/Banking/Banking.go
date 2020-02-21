package main

import (
	"fmt"
	"math"
)

type Banking_Payments interface {
	Compound_Amount() float64         // 如果我有X元，利息是Y%，那么我如果把钱完全存到银行，并且不做任何改动的情况下，Z年之后有多少钱
	Present_Worth() float64           // 如果你想在Z年后，Y%的利息下，得到U的钱，那么你现在要一下子存多少钱
	Sinking_Fund() float64            // 希望5年内能有100万，那么如果一年interest是10%, 你每年要往银行里存多少钱
	Uniform_Series_Compound() float64 // 如果每年往银行里存10000元，interest是10%，那么10年后，你银行里有多少钱
	Capital_Recovery() float64        // 我现在最多可以花出100万买房，那么如果我坚持24个月，每年利息6%，我每个月要换多少钱？
	Series_Present_Worth() float64    // 我最大的忍受程度是每个月500元的债务，一共坚持24个月，每年利息6%。那么我现在做多可以花多少钱买房？
}

type Uniform_Payments struct {
	target_value     float64
	current_value    float64
	annual_interest  float64
	years            float64
	annual_pay       float64
	monthly_interest float64
	months           float64
	monthly_pay      float64
}

func (u_p *Uniform_Payments) Compound_Amount(current_value float64, annual_interest float64, years float64) float64 {
	total_value := current_value * math.Pow(1.0+annual_interest, years)
	fmt.Println("Original amount is ", current_value, " after years ", years, " with annual_interest ", annual_interest, " is ", total_value)
	return total_value
}

func (u_p *Uniform_Payments) Present_Worth(target_value float64, annual_interest float64, years float64) float64 {
	current_value := target_value * (1 / math.Pow(1.0+annual_interest, years))
	fmt.Println("If you want to have ", target_value, " after years ", years, " with annual_interest ", annual_interest, " then you need ", current_value, "right now")
	return current_value
}

func (u_p *Uniform_Payments) Sinking_Fund(target_value float64, annual_interest float64, years float64) float64 {
	annual_pay := target_value * (annual_interest / (math.Pow(1.0+annual_interest, years) - 1))
	fmt.Println("If you want to have ", target_value, " in ", years, " years, and annual_interest is ", annual_interest, " . You need to deposit ", annual_pay, " each month")
	return annual_pay
}

func (u_p *Uniform_Payments) Uniform_Series_Compound(annual_pay float64, annual_interest float64, years float64) float64 {
	target_value := annual_pay * ((math.Pow(1.0+annual_interest, years) - 1) / annual_interest)
	fmt.Println("If you deposit ", annual_pay, " every year, and annual_interest is ", annual_interest, " for ", years, " years. Then you will have ", target_value, " eventually")
	return target_value
}

func (u_p *Uniform_Payments) Capital_Recovery(target_value float64, monthly_interest float64, months float64) float64 {
	monthly_pay := target_value * (monthly_interest * math.Pow((1.0+monthly_interest), months)) / (math.Pow((1.0+monthly_interest), months) - 1)
	fmt.Println("If you plan to spend ", target_value, ", and if monthly interest is ", monthly_interest, " and the entire loan period is ", months, " months. You have to return ", monthly_pay, " per month")
	return monthly_pay
}

func (u_p *Uniform_Payments) Series_Present_Worth(monthly_pay float64, monthly_interest float64, months float64) float64 {
	total_value := monthly_pay * ((math.Pow(1.0+monthly_interest, months) - 1) / (monthly_interest * math.Pow(1.0+monthly_interest, months)))
	fmt.Println("If you can have monthly loan ", monthly_pay, " with interest ", monthly_interest, " for months ", months, " then you can spend ", total_value)
	return total_value
}

func main() {
	curr_money := &Uniform_Payments{current_value: 1000000.0, annual_interest: 0.1, years: 15.0}
	curr_money.Compound_Amount(curr_money.current_value, curr_money.annual_interest, curr_money.years)

	curr_money2 := &Uniform_Payments{target_value: 4000000.0, annual_interest: 0.1, years: 15.0}
	curr_money2.Present_Worth(curr_money2.target_value, curr_money2.annual_interest, curr_money2.years)

	curr_money5 := &Uniform_Payments{target_value: 50000.0, annual_interest: 0.1, years: 5.0}
	curr_money5.Sinking_Fund(curr_money5.target_value, curr_money5.annual_interest, curr_money5.years)

	curr_money6 := &Uniform_Payments{annual_pay: 8190.0, annual_interest: 0.1, years: 5.0}
	curr_money6.Uniform_Series_Compound(curr_money6.annual_pay, curr_money6.annual_interest, curr_money6.years)

	curr_money3 := &Uniform_Payments{target_value: 21290.0, monthly_interest: 0.005, months: 4.0 * 12.0}
	curr_money3.Capital_Recovery(curr_money3.target_value, curr_money3.monthly_interest, curr_money3.months)

	curr_money4 := &Uniform_Payments{monthly_pay: 500.0, monthly_interest: 0.005, months: 4.0 * 12.0}
	curr_money4.Series_Present_Worth(curr_money4.monthly_pay, curr_money4.monthly_interest, curr_money4.months)
}
