package constdef

const (
	TaskState_Dormant = iota
	TaskState_In_Progress
	TaskState_Completed
	TaskState_Discard
)

var TaskState = map[int]string{
	TaskState_Dormant:     "TaskState_Dormant",
	TaskState_In_Progress: "TaskState_In_Progress",
	TaskState_Completed:   "TaskState_Completed",
	TaskState_Discard:     "TaskState_Discard",
}
