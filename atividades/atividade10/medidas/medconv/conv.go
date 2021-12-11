package medconv

func MToP(m Metros) Pés    { return Pés(m * 3.281) }
func PToM(p Pés) Metros    { return Metros(p / 3.281) }
func QToL(q Quilos) Libras { return Libras(q * 2.205) }
func LToQ(l Libras) Quilos { return Quilos(l / 2.205) }
