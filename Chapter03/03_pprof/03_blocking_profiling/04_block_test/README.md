To run this test with the block profile use:

`go test -blockprofile=block.pprof ./Chapter03/03_pprof/03_blocking_profiling/04_block_test/`

Then you can access the profile using:

`go tool pprof block.pprof`