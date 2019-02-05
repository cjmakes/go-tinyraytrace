# go-tinyraytrace
I saw [tinyraytrace](https://github.com/ssloy/tinyraytracer) and thought, "wow
that is pretty cool and doesn't look too bad" so I decided to follow the
amazing wiki attached to the same repo and try my hand at implimenting it in
golang. I chose golang because it is a language I am trying to get better at
and this project was the perfect oppertunity for me to learn better golang.

# How to use
- Clone the repo
- Set the `RAYOUT` envriornment variable as the full path of the file you want
	to output to. Should be something like `$PWD/out.ppm`
- Compile and run with `bazel run //:ray_trace`

# Tests
There are also some tests included, they are defined in the `BUILD` file and
run with: `bazel test //:<test_name>`
