// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"atc_thrift"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  AtcdInfo get_atcd_info()")
	fmt.Fprintln(os.Stderr, "  ShapingGroup create_group(string member)")
	fmt.Fprintln(os.Stderr, "  ShapingGroup get_group(i64 id)")
	fmt.Fprintln(os.Stderr, "  ShapingGroup get_group_with(string member)")
	fmt.Fprintln(os.Stderr, "  string get_group_token(i64 id)")
	fmt.Fprintln(os.Stderr, "  void leave_group(i64 id, string to_remove, string token)")
	fmt.Fprintln(os.Stderr, "  void join_group(i64 id, string to_add, string token)")
	fmt.Fprintln(os.Stderr, "  Shaping shape_group(i64 id, Shaping settings, string token)")
	fmt.Fprintln(os.Stderr, "  void unshape_group(i64 id, string token)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := atc_thrift.NewAtcdClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "get_atcd_info":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetAtcdInfo requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetAtcdInfo())
		fmt.Print("\n")
		break
	case "create_group":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateGroup requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.CreateGroup(value0))
		fmt.Print("\n")
		break
	case "get_group":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetGroup requires 1 args")
			flag.Usage()
		}
		argvalue0, err23 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err23 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetGroup(value0))
		fmt.Print("\n")
		break
	case "get_group_with":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetGroupWith requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetGroupWith(value0))
		fmt.Print("\n")
		break
	case "get_group_token":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetGroupToken requires 1 args")
			flag.Usage()
		}
		argvalue0, err25 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err25 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetGroupToken(value0))
		fmt.Print("\n")
		break
	case "leave_group":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "LeaveGroup requires 3 args")
			flag.Usage()
		}
		argvalue0, err26 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err26 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.LeaveGroup(value0, value1, value2))
		fmt.Print("\n")
		break
	case "join_group":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "JoinGroup requires 3 args")
			flag.Usage()
		}
		argvalue0, err29 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err29 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.JoinGroup(value0, value1, value2))
		fmt.Print("\n")
		break
	case "shape_group":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "ShapeGroup requires 3 args")
			flag.Usage()
		}
		argvalue0, err32 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err32 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg33 := flag.Arg(2)
		mbTrans34 := thrift.NewTMemoryBufferLen(len(arg33))
		defer mbTrans34.Close()
		_, err35 := mbTrans34.WriteString(arg33)
		if err35 != nil {
			Usage()
			return
		}
		factory36 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt37 := factory36.GetProtocol(mbTrans34)
		argvalue1 := atc_thrift.NewShaping()
		err38 := argvalue1.Read(jsProt37)
		if err38 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.ShapeGroup(value0, value1, value2))
		fmt.Print("\n")
		break
	case "unshape_group":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UnshapeGroup requires 2 args")
			flag.Usage()
		}
		argvalue0, err40 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err40 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.UnshapeGroup(value0, value1))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
