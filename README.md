# Go ORMS Benchmarks

it is very basic and simple benchmarks test for getting results that are close to reality so we can chose which GORM Lib have best result for usage in product project

# implementation:

we use postgress (run in docker) as database and northwind data for dataset , all is available in respository

# run:

first you should execute **.sql** files in your postgress database the edit **sqlboiler.toml** base on your configuration
then run this command:(for generate sqlboiler models)

    sqlboiler psql

# Result

> ## Testing System :
>
> goos: windows
> goarch: amd64
> cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
>
> ### Select One column benchmarks
>
> BenchmarkGromSelectInts1-12 1879 639467 ns/op 6967 B/op 154 allocs/op
>
> BenchmarkBoilerSelectInts1-12 1656 683596 ns/op 3178 B/op 66 allocs/op
>
> ### Select All columns benchmarks
>
> BenchmarkGromSelectAll-12 520 2356912 ns/op 129725 B/op 161 allocs/op
>
> BenchmarkBoilerSelectAll-12 288 9832391 ns/op 816028 B/op 23939
> allocs/op
>
> ### Select All columns with limit benchmarks
>
> BenchmarkGromSelectAll10-12 1791 650373 ns/op 8081 B/op 149 allocs/op
>
> BenchmarkBoilerSelectAll10-12 1418 738644 ns/op 12449 B/op 334
> allocs/op
>
> BenchmarkGromSelectAll100-12 1346 844146 ns/op 21420 B/op 151
> allocs/op
>
> BenchmarkBoilerSelectAll100-12 978 1233562 ns/op 100778 B/op 2932
> allocs/op
>
> BenchmarkGromSelectAll1000-12 543 2304961 ns/op 129863 B/op 165
> allocs/op
>
> BenchmarkBoilerSelectAll1000-12 100 10973199 ns/op 809278 B/op 23937
> allocs/op
>
> PASS
>
> ok ormsbench 15.940s

## Conclusion

Grom has a relatively better result, and also when we consider the simplicity of its implementation , so I do not see any reason to use Sqlboiler
