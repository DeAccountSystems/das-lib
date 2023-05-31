package witness

import (
	"fmt"
	"github.com/dotbitHQ/das-lib/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccountId(t *testing.T) {
	accounts := []string{"test.test.bit", "reverse.test.bit"}
	outs := make([]string, 0)
	for _, v := range accounts {
		out := common.Bytes2Hex(common.GetAccountIdByAccount(v))
		outs = append(outs, out)
	}
	t.Log(outs)
}

func TestRuleSpecialCharacters(t *testing.T) {
	rule := NewSubAccountRuleEntity("test.bit")

	price := 100000000

	err := rule.ParseFromJSON([]byte(fmt.Sprintf(`
{
    "version": 1,
    "rules": [
        {
            "name": "特殊字符账户",
            "note": "",
            "price": %d,
            "ast": {
                "type": "function",
                "name": "include_chars",
                "arguments": [
                    {
                        "type": "variable",
                        "name": "account_chars"
                    },
                    {
                        "type": "value",
                        "value_type": "string[]",
                        "value": [
                            "⚠️",
                            "❌",
                            "✅"
                        ]
                    }
                ]
            }
        }
    ]
}
`, price)))
	if err != nil {
		t.Fatal(err)
	}

	witness, err := rule.GenWitnessData(common.ActionDataTypeSubAccountPriceRules)
	assert.NoError(t, err)
	for _, v := range witness {
		t.Log(common.Bytes2Hex(v))
	}

	hit, _, err := rule.Hit("jerry.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, idx, err := rule.Hit("jerry⚠️.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 0)
	assert.EqualValues(t, rule.Rules[idx].Price, price)

	hit, idx, err = rule.Hit("jerry❌.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 0)
	assert.EqualValues(t, rule.Rules[idx].Price, price)

	hit, idx, err = rule.Hit("jerry✅.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 0)
	assert.EqualValues(t, rule.Rules[idx].Price, price)

	hit, _, err = rule.Hit("jerry💚.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	res, err := rule.GenWitnessData(common.ActionDataTypeSubAccountPriceRules)
	assert.NoError(t, err)

	parseRules := NewSubAccountRuleEntity("test.bit")
	err = parseRules.ParseFromDasActionWitnessData(res)
	assert.NoError(t, err)
	assert.EqualValues(t, len(parseRules.Rules), 1)

	assert.EqualValues(t, parseRules.Rules[0].Name, "特殊字符账户")
	assert.EqualValues(t, parseRules.Rules[0].Price, price)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Type, Function)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Name, FunctionIncludeCharts)
	assert.EqualValues(t, len(parseRules.Rules[0].Ast.Arguments), 2)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Arguments[0].Type, Variable)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Arguments[0].Name, AccountChars)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Arguments[1].Type, Value)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Arguments[1].ValueType, StringArray)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Arguments[1].Value, []string{"⚠️", "❌", "✅"})
}

func TestAccountLengthPrice(t *testing.T) {
	rule := NewSubAccountRuleEntity("test.bit")

	price100 := uint64(100000000)
	price10 := uint64(10000000)
	price1 := uint64(100000)

	err := rule.ParseFromJSON([]byte(fmt.Sprintf(`
{
    "version": 1,
    "rules": [
        {
            "name": "1 位账户",
            "note": "",
            "price": %d,
            "ast": {
                "type": "operator",
                "symbol": "==",
                "expressions": [
                    {
                        "type": "variable",
                        "name": "account_length"
                    },
                    {
                        "type": "value",
                        "value_type": "uint8",
                        "value": 1
                    }
                ]
            }
        },
        {
            "name": "2 位账户",
            "note": "",
            "price": %d,
            "ast": {
                "type": "operator",
                "symbol": "==",
                "expressions": [
                    {
                        "type": "variable",
                        "name": "account_length"
                    },
                    {
                        "type": "value",
                        "value_type": "uint8",
                        "value": 2
                    }
                ]
            }
        },
        {
            "name": "8 位及以上账户",
            "note": "",
            "price": %d,
            "ast": {
                "type": "operator",
                "symbol": ">=",
                "expressions": [
                    {
                        "type": "variable",
                        "name": "account_length"
                    },
                    {
                        "type": "value",
                        "value_type": "uint8",
                        "value": 8
                    }
                ]
            }
        }
    ]
}
`, price100, price10, price1)))
	if err != nil {
		t.Fatal(err)
	}

	witness, err := rule.GenWitnessData(common.ActionDataTypeSubAccountPriceRules)
	assert.NoError(t, err)
	for _, v := range witness {
		t.Log(common.Bytes2Hex(v))
	}

	hit, idx, err := rule.Hit("1.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 0)
	assert.EqualValues(t, rule.Rules[idx].Price, price100)

	hit, idx, err = rule.Hit("22.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 1)
	assert.EqualValues(t, rule.Rules[idx].Price, price10)

	hit, _, err = rule.Hit("333.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, _, err = rule.Hit("4444.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, _, err = rule.Hit("55555.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, _, err = rule.Hit("666666.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, _, err = rule.Hit("7777777.bit")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, idx, err = rule.Hit("88888888.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 2)
	assert.EqualValues(t, rule.Rules[idx].Price, price1)

	hit, idx, err = rule.Hit("999999999.bit")
	assert.NoError(t, err)
	assert.True(t, hit)
	assert.Equal(t, idx, 2)
	assert.EqualValues(t, rule.Rules[idx].Price, price1)

	res, err := rule.GenWitnessData(common.ActionDataTypeSubAccountPriceRules)
	assert.NoError(t, err)

	parseRules := NewSubAccountRuleEntity("test.bit")
	err = parseRules.ParseFromDasActionWitnessData(res)
	assert.NoError(t, err)
	assert.EqualValues(t, len(parseRules.Rules), 3)

	assert.EqualValues(t, parseRules.Rules[0].Name, "1 位账户")
	assert.EqualValues(t, parseRules.Rules[0].Price, price100)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Type, Operator)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Symbol, Equ)
	assert.EqualValues(t, len(parseRules.Rules[0].Ast.Expressions), 2)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Expressions[0].Type, Variable)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Expressions[0].Name, AccountLength)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Expressions[1].Type, Value)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Expressions[1].ValueType, Uint8)
	assert.EqualValues(t, parseRules.Rules[0].Ast.Expressions[1].Value, 1)

	assert.EqualValues(t, parseRules.Rules[1].Price, price10)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Type, Operator)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Symbol, Equ)
	assert.EqualValues(t, len(parseRules.Rules[1].Ast.Expressions), 2)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Expressions[0].Type, Variable)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Expressions[0].Name, AccountLength)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Expressions[1].Type, Value)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Expressions[1].ValueType, Uint8)
	assert.EqualValues(t, parseRules.Rules[1].Ast.Expressions[1].Value, 2)

	assert.EqualValues(t, parseRules.Rules[2].Price, price1)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Type, Operator)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Symbol, Gte)
	assert.EqualValues(t, len(parseRules.Rules[2].Ast.Expressions), 2)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Expressions[0].Type, Variable)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Expressions[0].Name, AccountLength)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Expressions[1].Type, Value)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Expressions[1].ValueType, Uint8)
	assert.EqualValues(t, parseRules.Rules[2].Ast.Expressions[1].Value, 8)

}

func TestRuleWhitelist(t *testing.T) {
	rule := NewSubAccountRuleEntity("test.bit")

	price := 100000000

	err := rule.ParseFromJSON([]byte(fmt.Sprintf(`
{
    "version": 1,
    "rules": [
        {
            "name": "特殊账户",
            "note": "",
            "price": %d,
            "ast": {
                "type": "function",
                "name": "in_list",
                "arguments": [
                    {
                        "type": "variable",
                        "name": "account"
                    },
                    {
                        "type": "value",
                        "value_type": "binary[]",
                        "value": [
                            "0x6ade4c435b8f3c4cf52336c9dd9dac71ed98520d",
                            "0xa84c83477c8f43670e70cef260da053818d770a5"
                        ]
                    }
                ]
            }
        }
    ]
}
`, price)))
	if err != nil {
		t.Fatal(err)
	}

	witness, err := rule.GenWitnessData(common.ActionDataTypeSubAccountPriceRules)
	assert.NoError(t, err)
	for _, v := range witness {
		t.Log(common.Bytes2Hex(v))
	}

	hit, _, err := rule.Hit("jerry")
	assert.NoError(t, err)
	assert.False(t, hit)

	hit, _, err = rule.Hit("test")
	assert.NoError(t, err)
	assert.True(t, hit)

	hit, _, err = rule.Hit("reverse")
	assert.NoError(t, err)
	assert.True(t, hit)

	res, err := rule.GenWitnessData(common.ActionDataTypeSubAccountPriceRules)
	assert.NoError(t, err)

	parseRules := NewSubAccountRuleEntity("test.bit")
	err = parseRules.ParseFromDasActionWitnessData(res)
	assert.NoError(t, err)
	assert.EqualValues(t, len(parseRules.Rules), 1)

	parseRule := parseRules.Rules[0]
	assert.EqualValues(t, parseRule.Name, "特殊账户")
	assert.EqualValues(t, parseRule.Note, "")
	assert.EqualValues(t, parseRule.Price, price)
	assert.EqualValues(t, parseRule.Ast.Type, "function")
	assert.EqualValues(t, parseRule.Ast.Name, "in_list")
	assert.EqualValues(t, len(parseRule.Ast.Arguments), 2)
	assert.EqualValues(t, parseRule.Ast.Arguments[0].Type, "variable")
	assert.EqualValues(t, parseRule.Ast.Arguments[0].Name, "account")
	assert.EqualValues(t, parseRule.Ast.Arguments[1].Type, "value")
	assert.EqualValues(t, parseRule.Ast.Arguments[1].ValueType, "binary[]")
	assert.EqualValues(t, len(parseRule.Ast.Arguments[1].Value.([]string)), 2)
	assert.EqualValues(t, parseRule.Ast.Arguments[1].Value.([]string)[0], "0x6ade4c435b8f3c4cf52336c9dd9dac71ed98520d")
	assert.EqualValues(t, parseRule.Ast.Arguments[1].Value.([]string)[1], "0xa84c83477c8f43670e70cef260da053818d770a5")
}

func TestSubAccountRule_Parser(t *testing.T) {
	data := common.Hex2Bytes("0x0400000001000000f3010000f301000010000000ae0000004c0100009e000000180000001c0000002b0000002f00000037000000000000000b0000003120e4bd8de8b4a6e688b70000000000e1f50500000000670000000c0000000d0000000056000000560000000c0000000d00000007490000000c000000260000001a0000000c0000000d0000000209000000090000000800000002230000000c0000000d0000000312000000120000000c0000000d0000000101000000019e000000180000001c0000002b0000002f00000037000000010000000b0000003220e4bd8de8b4a6e688b7000000008096980000000000670000000c0000000d0000000056000000560000000c0000000d00000007490000000c000000260000001a0000000c0000000d0000000209000000090000000800000002230000000c0000000d0000000312000000120000000c0000000d000000010100000002a7000000180000001c00000034000000380000004000000002000000140000003820e4bd8de58f8ae4bba5e4b88ae8b4a6e688b700000000a086010000000000670000000c0000000d0000000056000000560000000c0000000d00000004490000000c000000260000001a0000000c0000000d0000000209000000090000000800000002230000000c0000000d0000000312000000120000000c0000000d000000010100000008")
	rule := NewSubAccountRuleEntity("sub-account-test.bit")
	assert.NoError(t, rule.ParseFromWitnessData([][]byte{data}))
	assert.NoError(t, rule.Check())
}
