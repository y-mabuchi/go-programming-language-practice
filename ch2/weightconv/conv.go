package weightconv

func KToP(k Kilogram) Pond {
	return Pond(k * 2.2046)
}

func PToK(p Pond) Kilogram {
	return Kilogram(p / 2.2046)
}
