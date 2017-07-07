# Concurrency in Guardian

Guardian was originally designed to have a similar concurrency model to Golang (one of the chief inspirations for many of the features of the language), but the two have diverged over time, and Guardian now has its own unique combination of concurrency components.

## Threading

Go correctly distinguishes between OS threads and artificial threads which can be dynamically created and adjusted (known as goroutines). This has three crucial benefits:

1) Increasing the number of available threads
2) Allowing for total control over mechanisms of communication between threads
3) Something

However, goroutines fail to address three desirable characteristics of pseudo-threads.

Guardian therefore proposes and implements a new pseudo-thread protocol - "".

## Messaging

Guardian keeps the channel syntax from Go, and there is only one consequential difference between the two protocols:
