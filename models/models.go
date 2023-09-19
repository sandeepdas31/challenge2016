package models

type DistributorDetailsInput struct {
	DistriburtorID      int
	DistributorLocation DistributionLocationInfo
}

type DistributionLocationInfo struct {
	Include []Location
	Exclude []Location
}

type Location struct {
	Country  string
	Province string
	City     string
}

type DistributorPermissions struct {
	DistriburtorID      int
	DistributorLocation Location
}

type AddSubDistributor struct {
	SubDistriburtorID int
	Distributor       DistributorDetailsInput
}
