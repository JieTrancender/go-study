module loadgen

go 1.17

require (
    lib v0.0.0
    testhelper v0.0.0
)

replace (
    lib => ./lib
    testhelper => ./testhelper
)
