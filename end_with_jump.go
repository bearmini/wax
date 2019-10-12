package wax

type EndWithJump struct {
	LabelsExited uint32
}

func NewEndWithJump(labelsExited uint32) *EndWithJump {
	if labelsExited == 0 {
		return nil
	}
	return &EndWithJump{
		LabelsExited: labelsExited,
	}
}

func (e *EndWithJump) Error() string {
	return "end with jump"
}
