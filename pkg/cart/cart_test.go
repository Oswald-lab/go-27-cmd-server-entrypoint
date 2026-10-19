package cart

import "testing"

func TestTotal(t *testing.T) {
	c := New()
	c.Add("sku-1", 2, 1999)
	c.Add("sku-2", 1, 500)
	if got := c.TotalCents(); got != 4498 {
		t.Fatalf("total = %d, want 4498", got)
	}
}
