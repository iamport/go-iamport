package util

// Status 결제 상태
type Status string

const (
	StatusAll      = "all"      // 전체
	StatusReady    = "ready"    // 미결제
	StatusPaid     = "paid"     // 결제완료
	StatusCanceled = "canceled" // 결제취소
	StatusFailed   = "failed"   // 결제실패

	SortDESCStarted = "-started" // 결제시작시각(결제창오픈시각) 기준 내림차순(DESC) 정렬
	SortASCStarted  = "started"  // 결제시작시각(결제창오픈시각) 기준 오름차순(ASC) 정렬
	SortDESCPaid    = "-paid"    // 결제완료시각 기준 내림차순(DESC) 정렬
	SortASCPaid     = "paid"     // 결제완료시각 기준 오름차순(ASC) 정렬
	SortDESCUpdated = "-updated" // 최종수정시각(결제건 상태변화마다 수정시각 변경됨) 기준 내림차순(DESC) 정렬
	SortASCUpdated  = "updated"  // 최종수정시각(결제건 상태변화마다 수정시각 변경됨) 기준 오름차순(ASC) 정렬
)
