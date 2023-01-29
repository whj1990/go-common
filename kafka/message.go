package kafka

const (
	ProcessSubscriptionBusiness                            = "process_subscription"
	ProcessInstanceBusiness                                = "process_instance"
	ProcessFinanceTransferBusiness                         = "process_finance_transfer"
	ProcessFinanceTransferResultBusiness                   = "process_finance_transfer_result"
	ProcessFinanceTransactionDetailProcessInstanceBusiness = "process_finance_transaction_detail_process_instance"
	ProcessFinanceTransferElectronicReceiptBusiness        = "process_finance_transfer_electronic_receipt"

	DingtalkType                       = "dingtalk"
	ProcessInstanceType                = "process_instance"
	ProcessInstanceTaskType            = "process_instance_task"
	ProcessInstanceOperationRecordType = "process_instance_operation_record"
	FinanceIntrabankTransferType       = "intrabank_transfer"

	TaskBusiness = "task"

	DataSubmitTaskBusiness = "data_submit_task"
	DataCheckTaskBusiness  = "data_check_task"
	DataKillTaskBusiness   = "data_kill_task"

	IndexSubmitTaskBusiness = "index_submit_task"
	IndexCheckTaskBusiness  = "index_check_task"
	IndexKillTaskBusiness   = "index_kill_task"

	TagSubmitTaskBusiness = "tag_submit_task"
	TagCheckTaskBusiness  = "tag_check_task"
	TagKillTaskBusiness   = "tag_kill_task"

	SyncMetadataBusiness = "sync_metadata_business"

	K8SCleanTaskBusiness = "k8s_clean_task"

	SyncLakehouseMetadataBusiness = "sync_lakehouse_metadata_business"
)

type TypeMessage struct {
	Business string      `json:"business"`
	Type     string      `json:"type"`
	Id       int64       `json:"id"`
	Name     string      `json:"name"`
	Obj      interface{} `json:"obj"`
}
