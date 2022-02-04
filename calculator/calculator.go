package calculator

func Add(val1 MixedFraction, val2 MixedFraction) MixedFraction {
	frac1 := val1.ToImproperFraction()
	frac2 := val2.ToImproperFraction()

	result := Fraction{
		Numerator:   (frac1.Numerator * frac2.Denominator) + (frac1.Denominator * frac2.Numerator),
		Denominator: frac1.Denominator * frac2.Denominator,
	}

	return result.ToMixedFraction()
}

func Subtract(val1 MixedFraction, val2 MixedFraction) MixedFraction {
	frac1 := val1.ToImproperFraction()
	frac2 := val2.ToImproperFraction()

	result := Fraction{
		Numerator:   (frac1.Numerator * frac2.Denominator) - (frac1.Denominator * frac2.Numerator),
		Denominator: frac1.Denominator * frac2.Denominator,
	}

	return result.ToMixedFraction()
}

func Multiply(val1 MixedFraction, val2 MixedFraction) MixedFraction {
	frac1 := val1.ToImproperFraction()
	frac2 := val2.ToImproperFraction()

	result := Fraction{
		Numerator:   frac1.Numerator * frac2.Numerator,
		Denominator: frac1.Denominator * frac2.Denominator,
	}

	return result.ToMixedFraction()
}

func Divide(val1 MixedFraction, val2 MixedFraction) MixedFraction {
	frac1 := val1.ToImproperFraction()
	frac2 := val2.ToImproperFraction()

	result := Fraction{
		Numerator:   frac1.Numerator * frac2.Denominator,
		Denominator: frac1.Denominator * frac2.Numerator,
	}

	return result.ToMixedFraction()
}
