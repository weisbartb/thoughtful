package thoughtful

const (
	PackageLabelStandard = "STANDARD"
	PackageLabelSpecial  = "SPECIAL"
	PackageLabelRejected = "REJECTED"
)

const (
	MaxWeight    float64 = 20
	MaxArea      float64 = 1_000_000
	MaxDimension         = 150
)

func Sort(width, height, length, mass float64) (label string) {
	if width <= 0 || height <= 0 || length <= 0 || mass <= 0 {
		// Bounds check
		return PackageLabelRejected
	}
	var isOverweight, isBulky bool
	if mass >= MaxWeight {
		isOverweight = true
	}
	if width >= MaxDimension || height >= MaxDimension || length >= MaxDimension {
		isBulky = true
	} else {
		isBulky = width*height*length >= MaxArea
	}
	if isBulky || isOverweight {
		if isBulky && isOverweight {
			return PackageLabelRejected
		}
		return PackageLabelSpecial
	}
	return PackageLabelStandard
}
