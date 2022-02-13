package codepractice

import (
	"strings"
	"testing"
)

const msg = `Today marks seven years since we open-sourced our preliminary sketch of Go. With the help of the open source community, including more than a thousand individual contributors to the Go source repositories, Go has matured into a language used all over the world.

The most significant user-facing changes to Go over the past year are the addition of built-in support for HTTP/2 in Go 1.6 and the integration of the context package into the standard library in Go 1.7. But we’ve been making many less visible improvements. Go 1.7 changed the x86-64 compiler to use a new SSA-based back end, improving the performance of most Go programs by 10–20%. For Go 1.8, planned for release next February, we have changed the compilers for the other architectures to use the new back end too. We’ve also added new ports, to Android on 32-bit x86, Linux on 64-bit MIPS, and Linux on IBM z Systems. And we’ve developed new garbage-collection techniques that reduce typical “stop the world” pauses to under 100 microseconds. (Contrast that with Go 1.5’s big news of 10 milliseconds or less.)

This year kicked off with a global Go hackathon, the Gopher Gala, in January. Then there were Go conferences in India and Dubai in February, China and Japan in April, San Francisco in May, Denver in July, London in August, Paris last month, and Brazil this past weekend. And GothamGo in New York is next week. This year also saw more than 30 new Go user groups, eight new Women Who Go chapters, and four GoBridge workshops around the world.

We continue to be overwhelmed by and grateful for the enthusiasm and support of the Go community. Whether you participate by contributing changes, reporting bugs, sharing your expertise in design discussions, writing blog posts or books, running meetups, helping others learn or improve, open sourcing Go packages you wrote, or just being part of the Go community, the Go team thanks you for your help, your time, and your energy. Go would not be the success it is today without you.

Thank you, and here’s to another year of fun and success with Go!
Today marks seven years since we open-sourced our preliminary sketch of Go. With the help of the open source community, including more than a thousand individual contributors to the Go source repositories, Go has matured into a language used all over the world.

The most significant user-facing changes to Go over the past year are the addition of built-in support for HTTP/2 in Go 1.6 and the integration of the context package into the standard library in Go 1.7. But we’ve been making many less visible improvements. Go 1.7 changed the x86-64 compiler to use a new SSA-based back end, improving the performance of most Go programs by 10–20%. For Go 1.8, planned for release next February, we have changed the compilers for the other architectures to use the new back end too. We’ve also added new ports, to Android on 32-bit x86, Linux on 64-bit MIPS, and Linux on IBM z Systems. And we’ve developed new garbage-collection techniques that reduce typical “stop the world” pauses to under 100 microseconds. (Contrast that with Go 1.5’s big news of 10 milliseconds or less.)

This year kicked off with a global Go hackathon, the Gopher Gala, in January. Then there were Go conferences in India and Dubai in February, China and Japan in April, San Francisco in May, Denver in July, London in August, Paris last month, and Brazil this past weekend. And GothamGo in New York is next week. This year also saw more than 30 new Go user groups, eight new Women Who Go chapters, and four GoBridge workshops around the world.

We continue to be overwhelmed by and grateful for the enthusiasm and support of the Go community. Whether you participate by contributing changes, reporting bugs, sharing your expertise in design discussions, writing blog posts or books, running meetups, helping others learn or improve, open sourcing Go packages you wrote, or just being part of the Go community, the Go team thanks you for your help, your time, and your energy. Go would not be the success it is today without you.

Thank you, and here’s to another year of fun and success with Go!
Today marks seven years since we open-sourced our preliminary sketch of Go. With the help of the open source community, including more than a thousand individual contributors to the Go source repositories, Go has matured into a language used all over the world.

The most significant user-facing changes to Go over the past year are the addition of built-in support for HTTP/2 in Go 1.6 and the integration of the context package into the standard library in Go 1.7. But we’ve been making many less visible improvements. Go 1.7 changed the x86-64 compiler to use a new SSA-based back end, improving the performance of most Go programs by 10–20%. For Go 1.8, planned for release next February, we have changed the compilers for the other architectures to use the new back end too. We’ve also added new ports, to Android on 32-bit x86, Linux on 64-bit MIPS, and Linux on IBM z Systems. And we’ve developed new garbage-collection techniques that reduce typical “stop the world” pauses to under 100 microseconds. (Contrast that with Go 1.5’s big news of 10 milliseconds or less.)

This year kicked off with a global Go hackathon, the Gopher Gala, in January. Then there were Go conferences in India and Dubai in February, China and Japan in April, San Francisco in May, Denver in July, London in August, Paris last month, and Brazil this past weekend. And GothamGo in New York is next week. This year also saw more than 30 new Go user groups, eight new Women Who Go chapters, and four GoBridge workshops around the world.

We continue to be overwhelmed by and grateful for the enthusiasm and support of the Go community. Whether you participate by contributing changes, reporting bugs, sharing your expertise in design discussions, writing blog posts or books, running meetups, helping others learn or improve, open sourcing Go packages you wrote, or just being part of the Go community, the Go team thanks you for your help, your time, and your energy. Go would not be the success it is today without you.

Thank you, and here’s to another year of fun and success with Go!
Today marks seven years since we open-sourced our preliminary sketch of Go. With the help of the open source community, including more than a thousand individual contributors to the Go source repositories, Go has matured into a language used all over the world.

The most significant user-facing changes to Go over the past year are the addition of built-in support for HTTP/2 in Go 1.6 and the integration of the context package into the standard library in Go 1.7. But we’ve been making many less visible improvements. Go 1.7 changed the x86-64 compiler to use a new SSA-based back end, improving the performance of most Go programs by 10–20%. For Go 1.8, planned for release next February, we have changed the compilers for the other architectures to use the new back end too. We’ve also added new ports, to Android on 32-bit x86, Linux on 64-bit MIPS, and Linux on IBM z Systems. And we’ve developed new garbage-collection techniques that reduce typical “stop the world” pauses to under 100 microseconds. (Contrast that with Go 1.5’s big news of 10 milliseconds or less.)

This year kicked off with a global Go hackathon, the Gopher Gala, in January. Then there were Go conferences in India and Dubai in February, China and Japan in April, San Francisco in May, Denver in July, London in August, Paris last month, and Brazil this past weekend. And GothamGo in New York is next week. This year also saw more than 30 new Go user groups, eight new Women Who Go chapters, and four GoBridge workshops around the world.

We continue to be overwhelmed by and grateful for the enthusiasm and support of the Go community. Whether you participate by contributing changes, reporting bugs, sharing your expertise in design discussions, writing blog posts or books, running meetups, helping others learn or improve, open sourcing Go packages you wrote, or just being part of the Go community, the Go team thanks you for your help, your time, and your energy. Go would not be the success it is today without you.

Thank you, and here’s to another year of fun and success with Go!`

func TestCat(t *testing.T) {
	type args struct {
		str []string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{strings.Split(msg, " "), " "}, msg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cat(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("Cat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		str []string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{strings.Split(msg, " "), " "}, msg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCat(b *testing.B) {
	Cat(strings.Split(msg, " "), "")
}

func BenchmarkJoin(b *testing.B) {
	Cat(strings.Split(msg, " "), "")
}
