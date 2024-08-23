package payroll

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	yearlyPay := d.HourlyRate * d.HoursWorkedInYear
	name := d.Individual.FirstName + " " + d.Individual.LastName
	return name, yearlyPay
}

func convertRating(s string) float64 {
	switch rating := s; rating {
	case "Excellent":
		return 5.00
	case "Good":
		return 4.00

	case "Fair":
		return 3.00

	case "Poor":
		return 2.00

	case "Unsatisfactory":
		return 1.00

	default:
		return 0.00
	}

}

func (d Developer) ReviewRating() float64 {
	totalScore := 0.00
	for _, i := range d.Review {
		switch v := i.(type) {
		case string:
			totalScore += convertRating(v)
		case int:
			totalScore += float64(v)
		default:
			totalScore += 0
		}
	}

	return totalScore / float64(len(d.Review))
}
