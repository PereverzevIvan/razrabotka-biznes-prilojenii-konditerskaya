package order_params

type RepoGetAllParams struct {
	CustomerID  int
	ManagerID   int
	ORManagerID int
	StatusIDs   []int
}
