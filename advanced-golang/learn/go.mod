module learn.com/golang

go 1.20

require (
	github.com/marviigrey/learning-golang/projects/cryptit v0.0.0-20231005004335-f0fb98b787f0
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/marviigrey/learning-golang/projects/cryptit/decrypt => ../../projects/cryptit/decrypt
