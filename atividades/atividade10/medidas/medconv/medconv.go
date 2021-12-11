package medconv

import "fmt"

type Metros float64
type Pés float64
type Quilos float64
type Libras float64

func (m Metros) String() string { return fmt.Sprintf("%.2f m", m) }
func (p Pés) String() string    { return fmt.Sprintf("%.2f ft", p) }
func (q Quilos) String() string { return fmt.Sprintf("%.2f kg", q) }
func (l Libras) String() string { return fmt.Sprintf("%.2f lb", l) }
