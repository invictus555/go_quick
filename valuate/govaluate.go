package valuate

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func TestifyGoValuateSimpleExp() {
	expression, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func TestifyGoValuateComplexExp() {
	condition := "1>2?'sheng':'yang'"
	expression, err := govaluate.NewEvaluableExpression(condition)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func TestifyGoValuateComplexExpV2() {
	condition := "1 < 2? var1:var2"
	expression, err := govaluate.NewEvaluableExpression(condition)
	if err != nil {
		fmt.Println(err)
		return
	}

	parameters := make(map[string]interface{}, 2)
	parameters["var1"] = "12 < 12? 1:13"
	parameters["var2"] = 2048

	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
