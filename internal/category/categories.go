package category

type Category string

const (
	Beauty         Category = "Beauty"
	Birthday       Category = "Birthday"
	Donation       Category = "Donation"
	Education      Category = "Education"
	Experience     Category = "Experience"
	Fee            Category = "Fee"
	Fitness        Category = "Fitness"
	Gift           Category = "Gift"
	Grocery        Category = "Grocery"
	Medical        Category = "Medical"
	NightOut       Category = "NightOut"
	Rent           Category = "Rent"
	Restaurant     Category = "Restaurant"
	Shopping       Category = "Shopping"
	Subscription   Category = "Subscription"
	Taxes          Category = "Taxes"
	Transportation Category = "Transportation"
	Travel         Category = "Travel"
	Utilities      Category = "Utilities"
	Work           Category = "Work"
	Mortgage       Category = "Mortgage"
	Family         Category = "Family"
	Oreo           Category = "Oreo"
	Unknown        Category = "unknown"
)

func (c Category) String() string {
	switch c {
	case Beauty:
		return "Beauty"
	case Birthday:
		return "Birthday"
	case Donation:
		return "Donation"
	case Education:
		return "Education"
	case Experience:
		return "Experience"
	case Fee:
		return "Fee"
	case Fitness:
		return "Fitness"
	case Gift:
		return "Gift"
	case Grocery:
		return "Grocery"
	case NightOut:
		return "Night Out"
	case Rent:
		return "Rent"
	case Restaurant:
		return "Restaurant"
	case Shopping:
		return "Shopping"
	case Subscription:
		return "Subscription"
	case Taxes:
		return "Taxes"
	case Transportation:
		return "Transportation"
	case Travel:
		return "Travel"
	case Utilities:
		return "Utilities"
	case Work:
		return "Work"
	case Mortgage:
		return "Mortgage"
	case Medical:
		return "Medical"
	case Family:
		return "Family"
	case Oreo:
		return "Oreo"
	case Unknown:
		return "unknown"
	default:
		return "Unknown"
	}
}
