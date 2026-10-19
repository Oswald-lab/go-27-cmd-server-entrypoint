package cart

// Item is one line in a cart.
type Item struct {
	SKU        string
	Qty        int
	PriceCents int
}

// Cart holds items.
type Cart struct {
	items []Item
}

// New returns an empty cart.
func New() *Cart { return &Cart{} }

// Add appends an item.
func (c *Cart) Add(sku string, qty, priceCents int) {
	c.items = append(c.items, Item{SKU: sku, Qty: qty, PriceCents: priceCents})
}

// Count returns the number of lines.
func (c *Cart) Count() int { return len(c.items) }

// TotalCents returns the cart total in cents.
func (c *Cart) TotalCents() int {
	total := 0
	for _, it := range c.items {
		total += it.Qty * it.PriceCents
	}
	return total
}
