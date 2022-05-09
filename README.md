# postpone
simple golang module to execute func() on timeout if not cancelled before that

## synopsis
```golang	
  var ppe postpone.PostponeExecutor
  ppe.Postpone(func() {println("will not be printed")}, time.Millisecond*100)
  time.Sleep(time.Millisecond)
  ppe.Cancel()
  ppe.Postpone(func() {println("THIS will BE printed")}, time.Millisecond*100)
  time.Sleep(time.Millisecond*200)
  ppe.Cancel()
```
