package contract

type Store interface {
	UserStore
	DriverStore
	CabStore
	TripStore
}
