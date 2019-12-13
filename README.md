# FIZZBUZZ

Ye ole `fizzbuzz`.

Normal `go build`, then `./fizzbuzz`.

This will run forever, you're welcome.

If you're feeling impatient, you can stop it earlier than forever.

`./fizzbuzz impatient`

If you want to slow things down and count along

`./fizzbuzz gentle`

You can chain any command-line arguments, order doesn't matter.

If you're in a weird mood and want to fizzbuzz with a randomly generated sequence

`./fizzbuzz scatter-brained`

If you combine that with `impatient` it'll only ever stop if it lands on the number 100. **This is a feature, not a bug.**

## Next Steps

When fizzbuzz finds a prime number, it screams `PRIMO FIZZ!!!!!!`

I created a `Maybe` interface to basically simulate a Maybe monad that I don't use anywhere. I probably never will if we're being honest.
