package ltkgo

type LlrpMessage interface {
    AddSpecParameter(param LlrpSpecParameter) error
}
