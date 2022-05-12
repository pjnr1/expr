package node

import (
	"fmt"
)

type FloatNode struct {
	Value     float64
	TextValue string
}

func (iNode FloatNode) GetTextValue() string {
	return iNode.TextValue
}

//func (iNode IntNode) GetValue() int64 {
//	return iNode.Value
//}

func (iNode FloatNode) GetType() Type {
	return TypeFloat64
}

//func (iNode IntNode) SetValue(str string) {
//	iNode.TextValue = str
//}

func NewFloatNode(value float64) ValueNode {
	textValue := fmt.Sprintf("%f", value)
	return FloatNode{Value: value, TextValue: textValue}
}
