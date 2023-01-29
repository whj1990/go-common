package enum

const (
	ProcessInstanceRelevanceUserApproverType = "approver"
	ProcessInstanceRelevanceUserCcType       = "cc"

	ProcessInstanceFormComponentPhoto      = "DDPhotoField"
	ProcessInstanceFormComponentAttachment = "DDAttachment"

	ProcessInstanceTaskRunning = "RUNNING"

	TextMessageType = "text"
	CardMessageType = "action_card"
	OAMessageType   = "oa"

	ProcessInstanceTotalStatusRunning = "0"
	ProcessInstanceTotalStatusSuccess = "1"
	ProcessInstanceTotalStatusFail    = "2"
	ProcessInstanceTotalStatusRefuse  = "3"
	ProcessInstanceTotalStatusReturn  = "4"
	ProcessInstanceTotalStatusFinish  = "5"

	ProcessInstanceAgreeResult  = "agree"
	ProcessInstanceRefuseResult = "refuse"

	ProcessInstanceNewStatus        = "NEW"
	ProcessInstanceRunningStatus    = "RUNNING"
	ProcessInstanceTerminatedStatus = "TERMINATED"
	ProcessInstanceCompletedStatus  = "COMPLETED"
	ProcessInstanceCanceledStatus   = "CANCELED"
)
