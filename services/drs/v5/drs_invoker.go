package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/drs/v5/model"
)

type BatchCreateJobsAsyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateJobsAsyncInvoker) Invoke() (*model.BatchCreateJobsAsyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateJobsAsyncResponse), nil
	}
}

type BatchCreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagsInvoker) Invoke() (*model.BatchCreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagsResponse), nil
	}
}

type BatchDeleteJobsByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobsByIdInvoker) Invoke() (*model.BatchDeleteJobsByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobsByIdResponse), nil
	}
}

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type BatchExecuteJobActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchExecuteJobActionsInvoker) Invoke() (*model.BatchExecuteJobActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchExecuteJobActionsResponse), nil
	}
}

type BatchStopJobsActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopJobsActionInvoker) Invoke() (*model.BatchStopJobsActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopJobsActionResponse), nil
	}
}

type BatchTagActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTagActionInvoker) Invoke() (*model.BatchTagActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTagActionResponse), nil
	}
}

type CheckDataFilterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDataFilterInvoker) Invoke() (*model.CheckDataFilterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDataFilterResponse), nil
	}
}

type CollectColumnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectColumnsInvoker) Invoke() (*model.CollectColumnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectColumnsResponse), nil
	}
}

type CollectDbObjectsAsyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectDbObjectsAsyncInvoker) Invoke() (*model.CollectDbObjectsAsyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDbObjectsAsyncResponse), nil
	}
}

type CollectDbObjectsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectDbObjectsInfoInvoker) Invoke() (*model.CollectDbObjectsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDbObjectsInfoResponse), nil
	}
}

type CollectPositionAsyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectPositionAsyncInvoker) Invoke() (*model.CollectPositionAsyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectPositionAsyncResponse), nil
	}
}

type CommitAsyncJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CommitAsyncJobInvoker) Invoke() (*model.CommitAsyncJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitAsyncJobResponse), nil
	}
}

type CopyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyJobInvoker) Invoke() (*model.CopyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyJobResponse), nil
	}
}

type CountInstanceByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountInstanceByTagsInvoker) Invoke() (*model.CountInstanceByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountInstanceByTagsResponse), nil
	}
}

type CreateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateJobInvoker) Invoke() (*model.CreateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateJobResponse), nil
	}
}

type DeleteJdbcDriverInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJdbcDriverInvoker) Invoke() (*model.DeleteJdbcDriverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJdbcDriverResponse), nil
	}
}

type DeleteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobInvoker) Invoke() (*model.DeleteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobResponse), nil
	}
}

type DownloadBatchCreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBatchCreateTemplateInvoker) Invoke() (*model.DownloadBatchCreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBatchCreateTemplateResponse), nil
	}
}

type DownloadDbObjectTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDbObjectTemplateInvoker) Invoke() (*model.DownloadDbObjectTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDbObjectTemplateResponse), nil
	}
}

type ExecuteJobActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteJobActionInvoker) Invoke() (*model.ExecuteJobActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteJobActionResponse), nil
	}
}

type ExportOperationInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportOperationInfoInvoker) Invoke() (*model.ExportOperationInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportOperationInfoResponse), nil
	}
}

type ImportBatchCreateJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportBatchCreateJobsInvoker) Invoke() (*model.ImportBatchCreateJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportBatchCreateJobsResponse), nil
	}
}

type ListAsyncJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAsyncJobDetailInvoker) Invoke() (*model.ListAsyncJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAsyncJobDetailResponse), nil
	}
}

type ListAsyncJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAsyncJobsInvoker) Invoke() (*model.ListAsyncJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAsyncJobsResponse), nil
	}
}

type ListDbObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbObjectsInvoker) Invoke() (*model.ListDbObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbObjectsResponse), nil
	}
}

type ListInstanceByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceByTagsInvoker) Invoke() (*model.ListInstanceByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceByTagsResponse), nil
	}
}

type ListInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTagsInvoker) Invoke() (*model.ListInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTagsResponse), nil
	}
}

type ListJdbcDriversInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJdbcDriversInvoker) Invoke() (*model.ListJdbcDriversResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJdbcDriversResponse), nil
	}
}

type ListJobHistoryParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobHistoryParametersInvoker) Invoke() (*model.ListJobHistoryParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobHistoryParametersResponse), nil
	}
}

type ListJobParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobParametersInvoker) Invoke() (*model.ListJobParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobParametersResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ListLinksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLinksInvoker) Invoke() (*model.ListLinksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLinksResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ListsAgencyPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListsAgencyPermissionsInvoker) Invoke() (*model.ListsAgencyPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListsAgencyPermissionsResponse), nil
	}
}

type ShowActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowActionsInvoker) Invoke() (*model.ShowActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowActionsResponse), nil
	}
}

type ShowColumnInfoResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowColumnInfoResultInvoker) Invoke() (*model.ShowColumnInfoResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowColumnInfoResultResponse), nil
	}
}

type ShowComparePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComparePolicyInvoker) Invoke() (*model.ShowComparePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComparePolicyResponse), nil
	}
}

type ShowDataFilteringResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataFilteringResultInvoker) Invoke() (*model.ShowDataFilteringResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataFilteringResultResponse), nil
	}
}

type ShowDataProcessingRulesResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataProcessingRulesResultInvoker) Invoke() (*model.ShowDataProcessingRulesResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataProcessingRulesResultResponse), nil
	}
}

type ShowDataProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataProgressInvoker) Invoke() (*model.ShowDataProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataProgressResponse), nil
	}
}

type ShowDbObjectCollectionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDbObjectCollectionStatusInvoker) Invoke() (*model.ShowDbObjectCollectionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDbObjectCollectionStatusResponse), nil
	}
}

type ShowDbObjectTemplateProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDbObjectTemplateProgressInvoker) Invoke() (*model.ShowDbObjectTemplateProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDbObjectTemplateProgressResponse), nil
	}
}

type ShowDbObjectTemplateResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDbObjectTemplateResultInvoker) Invoke() (*model.ShowDbObjectTemplateResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDbObjectTemplateResultResponse), nil
	}
}

type ShowDbObjectsListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDbObjectsListInvoker) Invoke() (*model.ShowDbObjectsListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDbObjectsListResponse), nil
	}
}

type ShowDirtyDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDirtyDataInvoker) Invoke() (*model.ShowDirtyDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDirtyDataResponse), nil
	}
}

type ShowEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnterpriseProjectInvoker) Invoke() (*model.ShowEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnterpriseProjectResponse), nil
	}
}

type ShowHealthCompareJobListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHealthCompareJobListInvoker) Invoke() (*model.ShowHealthCompareJobListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHealthCompareJobListResponse), nil
	}
}

type ShowIncrementComponentsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIncrementComponentsDetailInvoker) Invoke() (*model.ShowIncrementComponentsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIncrementComponentsDetailResponse), nil
	}
}

type ShowInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceTagsInvoker) Invoke() (*model.ShowInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceTagsResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowMeteringInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMeteringInvoker) Invoke() (*model.ShowMeteringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMeteringResponse), nil
	}
}

type ShowMonitorDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonitorDataInvoker) Invoke() (*model.ShowMonitorDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonitorDataResponse), nil
	}
}

type ShowObjectMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowObjectMappingInvoker) Invoke() (*model.ShowObjectMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowObjectMappingResponse), nil
	}
}

type ShowPositionResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPositionResultInvoker) Invoke() (*model.ShowPositionResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPositionResultResponse), nil
	}
}

type ShowProgressDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProgressDataInvoker) Invoke() (*model.ShowProgressDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProgressDataResponse), nil
	}
}

type ShowReplayResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReplayResultsInvoker) Invoke() (*model.ShowReplayResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReplayResultsResponse), nil
	}
}

type ShowSupportObjectTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSupportObjectTypeInvoker) Invoke() (*model.ShowSupportObjectTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSupportObjectTypeResponse), nil
	}
}

type ShowUpdateObjectSavingStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUpdateObjectSavingStatusInvoker) Invoke() (*model.ShowUpdateObjectSavingStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpdateObjectSavingStatusResponse), nil
	}
}

type StopJobActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobActionInvoker) Invoke() (*model.StopJobActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobActionResponse), nil
	}
}

type SyncJdbcDriverInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncJdbcDriverInvoker) Invoke() (*model.SyncJdbcDriverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncJdbcDriverResponse), nil
	}
}

type UpdateBatchAsyncJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBatchAsyncJobsInvoker) Invoke() (*model.UpdateBatchAsyncJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBatchAsyncJobsResponse), nil
	}
}

type UpdateDataProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataProgressInvoker) Invoke() (*model.UpdateDataProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataProgressResponse), nil
	}
}

type UpdateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobInvoker) Invoke() (*model.UpdateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobResponse), nil
	}
}

type UpdateJobConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobConfigurationsInvoker) Invoke() (*model.UpdateJobConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobConfigurationsResponse), nil
	}
}

type UpdateStartPositionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStartPositionInvoker) Invoke() (*model.UpdateStartPositionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStartPositionResponse), nil
	}
}

type UploadDbObjectTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadDbObjectTemplateInvoker) Invoke() (*model.UploadDbObjectTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadDbObjectTemplateResponse), nil
	}
}

type UploadJdbcDriverInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadJdbcDriverInvoker) Invoke() (*model.UploadJdbcDriverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadJdbcDriverResponse), nil
	}
}

type ValidateJobNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateJobNameInvoker) Invoke() (*model.ValidateJobNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateJobNameResponse), nil
	}
}
