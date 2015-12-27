package opendmm

import (
  "testing"
  "sync"
)

func testEngine(t *testing.T, queries []string, search func(string, chan MovieMeta) *sync.WaitGroup) {
  for _, query := range queries {
    metach := make(chan MovieMeta)
    wg := search(query, metach)
    go func() {
      wg.Wait()
      close(metach)
    }()
    finalch := validateFields(trimSpaces(deduplicate(metach)))
    meta, ok := <-finalch
    if !ok {
      t.Error("Not found")
    } else {
      t.Logf("%+v", meta)
    }
  }
}
