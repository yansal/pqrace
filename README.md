# pqrace

This repository is a short program to reproduce a race in the github.com/lib/pq package. To reproduce, run the following commands:
```bash
go build -race -trimpath
DATABASE_URL="host=/path/to/postgresql/unix/socket/directory" ./pqrace
```
```
==================
WARNING: DATA RACE
Write at 0x00c00008c720 by goroutine 9:
  runtime.mapassign_faststr()
      runtime/map_faststr.go:202 +0x0
  github.com/lib/pq.dial()
      github.com/lib/pq@v1.2.0/conn.go:334 +0x6c9
  github.com/lib/pq.(*Connector).open()
      github.com/lib/pq@v1.2.0/conn.go:298 +0x295
  github.com/lib/pq.(*Connector).Connect()
      github.com/lib/pq@v1.2.0/connector.go:27 +0x50
  database/sql.(*DB).conn()
      database/sql/sql.go:1228 +0xa15
  database/sql.(*DB).query()
      database/sql/sql.go:1565 +0x76
  database/sql.(*DB).QueryContext()
      database/sql/sql.go:1547 +0xe4
  main.main.func1()
      database/sql/sql.go:1648 +0x12a

Previous read at 0x00c00008c720 by goroutine 10:
  runtime.mapaccess1_faststr()
      runtime/map_faststr.go:12 +0x0
  github.com/lib/pq.(*conn).handlePgpass()
      github.com/lib/pq@v1.2.0/conn.go:211 +0x21c
  github.com/lib/pq.(*Connector).open()
      github.com/lib/pq@v1.2.0/conn.go:296 +0x23f
  github.com/lib/pq.(*Connector).Connect()
      github.com/lib/pq@v1.2.0/connector.go:27 +0x50
  database/sql.(*DB).conn()
      database/sql/sql.go:1228 +0xa15
  database/sql.(*DB).query()
      database/sql/sql.go:1565 +0x76
  database/sql.(*DB).QueryContext()
      database/sql/sql.go:1547 +0xe4
  main.main.func1()
      database/sql/sql.go:1648 +0x12a

Goroutine 9 (running) created at:
  main.main()
      pqrace@/main.go:26 +0x15b

Goroutine 10 (running) created at:
  main.main()
      pqrace@/main.go:26 +0x15b
==================
==================
WARNING: DATA RACE
Read at 0x00c00008c720 by goroutine 10:
  runtime.mapaccess1_faststr()
      runtime/map_faststr.go:12 +0x0
  github.com/lib/pq.network()
      github.com/lib/pq@v1.2.0/conn.go:371 +0x71
  github.com/lib/pq.dial()
      github.com/lib/pq@v1.2.0/conn.go:331 +0x66
  github.com/lib/pq.(*Connector).open()
      github.com/lib/pq@v1.2.0/conn.go:298 +0x295
  github.com/lib/pq.(*Connector).Connect()
      github.com/lib/pq@v1.2.0/connector.go:27 +0x50
  database/sql.(*DB).conn()
      database/sql/sql.go:1228 +0xa15
  database/sql.(*DB).query()
      database/sql/sql.go:1565 +0x76
  database/sql.(*DB).QueryContext()
      database/sql/sql.go:1547 +0xe4
  main.main.func1()
      database/sql/sql.go:1648 +0x12a

Previous write at 0x00c00008c720 by goroutine 9:
  runtime.mapassign_faststr()
      runtime/map_faststr.go:202 +0x0
  github.com/lib/pq.dial()
      github.com/lib/pq@v1.2.0/conn.go:334 +0x6c9
  github.com/lib/pq.(*Connector).open()
      github.com/lib/pq@v1.2.0/conn.go:298 +0x295
  github.com/lib/pq.(*Connector).Connect()
      github.com/lib/pq@v1.2.0/connector.go:27 +0x50
  database/sql.(*DB).conn()
      database/sql/sql.go:1228 +0xa15
  database/sql.(*DB).query()
      database/sql/sql.go:1565 +0x76
  database/sql.(*DB).QueryContext()
      database/sql/sql.go:1547 +0xe4
  main.main.func1()
      database/sql/sql.go:1648 +0x12a

Goroutine 10 (running) created at:
  main.main()
      pqrace@/main.go:26 +0x15b

Goroutine 9 (running) created at:
  main.main()
      pqrace@/main.go:26 +0x15b
==================
==================
WARNING: DATA RACE
Write at 0x00c0000ba328 by goroutine 10:
  github.com/lib/pq.dial()
      github.com/lib/pq@v1.2.0/conn.go:334 +0x6e1
  github.com/lib/pq.(*Connector).open()
      github.com/lib/pq@v1.2.0/conn.go:298 +0x295
  github.com/lib/pq.(*Connector).Connect()
      github.com/lib/pq@v1.2.0/connector.go:27 +0x50
  database/sql.(*DB).conn()
      database/sql/sql.go:1228 +0xa15
  database/sql.(*DB).query()
      database/sql/sql.go:1565 +0x76
  database/sql.(*DB).QueryContext()
      database/sql/sql.go:1547 +0xe4
  main.main.func1()
      database/sql/sql.go:1648 +0x12a

Previous write at 0x00c0000ba328 by goroutine 9:
  github.com/lib/pq.dial()
      github.com/lib/pq@v1.2.0/conn.go:334 +0x6e1
  github.com/lib/pq.(*Connector).open()
      github.com/lib/pq@v1.2.0/conn.go:298 +0x295
  github.com/lib/pq.(*Connector).Connect()
      github.com/lib/pq@v1.2.0/connector.go:27 +0x50
  database/sql.(*DB).conn()
      database/sql/sql.go:1228 +0xa15
  database/sql.(*DB).query()
      database/sql/sql.go:1565 +0x76
  database/sql.(*DB).QueryContext()
      database/sql/sql.go:1547 +0xe4
  main.main.func1()
      database/sql/sql.go:1648 +0x12a

Goroutine 10 (running) created at:
  main.main()
      pqrace@/main.go:26 +0x15b

Goroutine 9 (running) created at:
  main.main()
      pqrace@/main.go:26 +0x15b
==================
Found 3 data race(s)
```
