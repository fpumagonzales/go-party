package party

type contactInfo struct {
	email string
	phone string
}

type party struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func NewParty(firstName, lastName, email string) *party {
	c := new(contactInfo)
	c.email = email

	p := new(party)
	p.firstName = firstName
	p.lastName = lastName
	p.contactInfo = *c
	return p
}
