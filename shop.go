package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/dunhamsteve/plist"

	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// var filename = "shop.shopshop"

var fn = "$HOME/Dropbox/ShopShop/Shopping List.shopshop"

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

type Item struct {
	Done  bool
	Count string
	Name  string
}

type ShopFile struct {
	Color        []float64
	ShoppingList []Item
}

func save(list *ShopFile) {
	out, err := plist.Marshal(list)
	must(err)
	ioutil.WriteFile(os.ExpandEnv(fn), out, 0644)
}

func process(line []string) {
	cmd := line[0]
	switch cmd {
	case "rm", "remove":
		idx, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		item := list.ShoppingList[idx]
		fmt.Println("Removing:", item.Count, item.Name)
		list.ShoppingList = append(list.ShoppingList[:idx], list.ShoppingList[idx+1:]...)
		save(list)
	case "add", "buy":
		count := ""
		buf := bytes.NewBuffer(nil)
		i := 1
		if _, err := strconv.Atoi(line[1]); err == nil {
			i = 2
			count = line[1]
		}
		for _, word := range line[i:] {
			if buf.Len() > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(word)
		}
		name := buf.String()
		list.ShoppingList = append(list.ShoppingList, Item{false, count, name})
		fmt.Println("Adding:", count, name)
		save(list)
	case "checkout", "co":
		var newList []Item
		for _, item := range list.ShoppingList {
			if !item.Done {
				newList = append(newList, item)
			}
		}
		list.ShoppingList = newList
		save(list)
	case "help":
		fmt.Println(`Commands:
  add ...  add item
  rm #     remove item at index
  co       checkout (remove done items)`)
	case "list", "ls":
		fmt.Println("Items:")
		for i, item := range list.ShoppingList {
			check := " "
			if item.Done {
				check = "@done"
			}
			fmt.Printf("%2d: %s %s %s\n", i, item.Count, item.Name, check)
		}
		fmt.Println()
	default:
		fmt.Println("Unknown command:", cmd)
	}
	fmt.Println()

}

func interact() {
	out := os.Stdout
	reader := bufio.NewReader(os.Stdin)
	for {
		out.WriteString("> ")
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			if len(line) < 2 {
				os.Exit(0)
			}
			l := strings.Split(line[:len(line)-1], " ")
			if len(l) > 0 {
				process(l)
			}
		case io.EOF:
			fmt.Println()
			os.Exit(0)
		default:
			panic(err)
		}
	}
}

var list *ShopFile

func main() {
	var interactive bool
	flag.BoolVar(&interactive, "i", false, "interactive mode")
	flag.Parse()
	f, err := os.Open(os.ExpandEnv(fn))
	must(err)

	list = new(ShopFile)

	err = plist.Unmarshal(f, list)

	args := flag.Args()

	if len(args) > 0 {
		process(args)
		process([]string{"ls"})
	} else if interactive {
		interact()
	} else {
		process([]string{"ls"})
		// interactive()
	}

}
