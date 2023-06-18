package actrunstat

type ActionRunStatus uint8

const (
	Pending ActionRunStatus = iota
	Running
	Success
	Failed
)
