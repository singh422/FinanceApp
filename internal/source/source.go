package source

type Source string

const (
	AMEX     Source = "amex"
	Apple    Source = "apple"
	Chase    Source = "chase"
	Discover Source = "discover"
	Paypal   Source = "paypal"
	Venmo    Source = "venmo"
	Unknown  Source = "unknown"
)

func (s Source) String() string {
	switch s {
	case AMEX:
		return "AMEX"
	case Apple:
		return "Apple"
	case Chase:
		return "Chase"
	case Discover:
		return "Discover"
	case Paypal:
		return "Paypal"
	case Venmo:
		return "Venmo"
	case Unknown:
		return "unknown"
	default:
		return "Unknown"
	}
}

func GetAllSources() []Source {
	return []Source{
		AMEX,
		Apple,
		Chase,
		Discover,
		Paypal,
		Venmo,
	}
}
