# isevengo

An unofficial Golang wrapper for https://isevenapi.xyz

## About

Go is fast. But is it fast enough? There are computations where it may struggle. One such computation is that of determining whether an integer is even or odd. This is quite a resouce intensive task. Fortunately, there exists the [isEven API](https://isevenapi.xyz) which does the task for us.

This wrapper can be used to access the API easily in Go application. The wrapper has no dependencies outside of the standard library, making it very small, easy to ship and secure (_looking at you, node modules_).

## Usage

Install Dependency:

    go get github.com/funoctis/isevengo

Check if a number is even:

    isNumEven, err  := isevengo.IsEven(num)
    if err != nil {
      // handle error
    }

    if isNumEven {
      fmt.Println("Number is even")
    }

The API also provides an ad with every result. Now if you're an advertisement enthusiast and want to see the ad, you can do:

    resp, err := isevengo.CallIsEvenAPI(num)
    if err != nil {
      // handle error
    }
    ad := resp.Ad
    fmt.Println(ad)

## Note

The free tier of the API is only limited to numbers between 0 and 999,999 (both included). If your job requires computing negative or larger numbers, I'm afraid this wrapper won't be able to help you.

## License

MIT License
