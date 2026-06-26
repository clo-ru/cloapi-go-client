package cloapi

import (
	"context"
	"errors"
	"testing"
)

// makeFetch returns a PageFunc serving `total` sequential ints in pages, recording
// the (limit, offset) of each call.
func makeFetch(total int) (PageFunc[int], *[][2]int) {
	var calls [][2]int
	fn := func(_ context.Context, limit, offset int) ([]int, int, error) {
		calls = append(calls, [2]int{limit, offset})
		var items []int
		for i := offset; i < offset+limit && i < total; i++ {
			items = append(items, i)
		}
		return items, total, nil
	}
	return fn, &calls
}

func TestPaginatorAll(t *testing.T) {
	fetch, calls := makeFetch(25)
	p := NewPaginator(10, fetch)

	got, err := p.All(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 25 {
		t.Fatalf("got %d items, want 25", len(got))
	}
	if p.Count() != 25 {
		t.Errorf("Count = %d, want 25", p.Count())
	}
	// 25 items / page 10 => offsets 0, 10, 20 => 3 calls.
	if len(*calls) != 3 {
		t.Fatalf("made %d calls, want 3: %v", len(*calls), *calls)
	}
	if p.HasNext() {
		t.Errorf("HasNext should be false after All")
	}
}

func TestPaginatorExactMultiple(t *testing.T) {
	fetch, calls := makeFetch(20)
	p := NewPaginator(10, fetch)
	if _, err := p.All(context.Background()); err != nil {
		t.Fatal(err)
	}
	// 20/10 => offsets 0, 10 ; after offset reaches 20 == count, stop. 2 calls.
	if len(*calls) != 2 {
		t.Fatalf("made %d calls, want 2: %v", len(*calls), *calls)
	}
}

func TestPaginatorNextThenExhausted(t *testing.T) {
	fetch, _ := makeFetch(5)
	p := NewPaginator(10, fetch)

	page, err := p.Next(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(page) != 5 {
		t.Fatalf("first page %d items, want 5", len(page))
	}
	if p.HasNext() {
		t.Errorf("HasNext should be false")
	}
	if _, err := p.Next(context.Background()); err == nil {
		t.Errorf("Next past end should error")
	}
}

func TestPaginatorPropagatesError(t *testing.T) {
	sentinel := errors.New("boom")
	p := NewPaginator(10, func(_ context.Context, _, _ int) ([]int, int, error) {
		return nil, 0, sentinel
	})
	if _, err := p.Next(context.Background()); !errors.Is(err, sentinel) {
		t.Fatalf("got %v, want sentinel", err)
	}
}

func TestPaginatorDefaultLimit(t *testing.T) {
	fetch, calls := makeFetch(3)
	p := NewPaginator(0, fetch) // invalid -> default 50
	if _, err := p.Next(context.Background()); err != nil {
		t.Fatal(err)
	}
	if (*calls)[0][0] != 50 {
		t.Fatalf("limit = %d, want default 50", (*calls)[0][0])
	}
}
