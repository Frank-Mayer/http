package internal

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	colorInfo        = color.New(color.FgBlack).Add(color.BgWhite)
	colorOk          = color.New(color.FgBlack).Add(color.BgGreen)
	colorRedirect    = color.New(color.FgBlack).Add(color.BgBlue)
	colorClientError = color.New(color.FgBlack).Add(color.BgRed)
	colorServerError = color.New(color.FgBlack).Add(color.BgRed)
)

func colorBool(val bool, expected bool) string {
	if val == expected {
		return color.GreenString(fmt.Sprint(val))
	} else {
		return color.RedString(fmt.Sprint(val))
	}
}

func PrintResponse(resp *http.Response, t time.Duration) {
	defer resp.Body.Close()

	fmt.Println("Took " + t.String())

	fmt.Printf("%s %s %s\n", resp.Request.Proto, resp.Request.Method, resp.Request.URL.String())

	{
		var c *color.Color
		if resp.StatusCode < 200 {
			// 0 - 199
			c = colorInfo
		} else if resp.StatusCode < 300 {
			// 200 - 299
			c = colorOk
		} else if resp.StatusCode < 400 {
			// 300 - 399
			c = colorRedirect
		} else if resp.StatusCode < 500 {
			// 400 - 499
			c = colorClientError
		} else {
			// 500 +
			c = colorServerError
		}
		fmt.Println(c.Sprintf(" %s ", resp.Status))
	}

	// Print Headers
	header := color.New(color.FgCyan).SprintfFunc()
	fmt.Println("Headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", header(key), value)
		}
	}

	// Print Transfer-Encoding
	fmt.Print("Transfer-Encoding: ")
	if resp.TransferEncoding == nil || len(resp.TransferEncoding) == 0 {
		fmt.Println("None")
	} else {
		fmt.Println(strings.Join(resp.TransferEncoding, "; "))
	}

	// Print Uncompressed
	fmt.Println("Uncompressed:", colorBool(resp.Uncompressed, false))

	// Print Trailer Header
	fmt.Println("Trailer Headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", header(key), value)
		}
	}

	// Print TLS ConnectionState
	fmt.Println("TLS ConnectionState:")
	printTLSConnectionState(resp.TLS)

	// Print Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	} else if len(body) != 0 {
		fmt.Println("Body:")
		fmt.Printf("%s\n", body)
	}
}

func printTLSConnectionState(tlsState *tls.ConnectionState) {
	if tlsState == nil {
		fmt.Println("  TLS not available")
		return
	}
	fmt.Println("  Version:", tlsState.Version)
	fmt.Println("  HandshakeComplete:", tlsState.HandshakeComplete)
	fmt.Println("  CipherSuite:", tlsState.CipherSuite)
	fmt.Println("  NegotiatedProtocol:", tlsState.NegotiatedProtocol)
}
