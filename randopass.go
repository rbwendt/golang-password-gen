package main

import (
    "io/ioutil"
    "fmt"
    "strings"
    "flag"
    "math/rand"
    "time"
)
var num_passwords = flag.Int("passwords", 1, "The number of passwords you would like to see")
var num_words = flag.Int("words", 4, "The number of words you would like to use")
var separators_string = flag.String("separators", " ", "The word separators you would like to use")

func main() {
    flag.Parse()
    if *num_passwords == 1 {
        fmt.Println("Your generated password is:")
    } else if *num_passwords > 1 {
        fmt.Println("Your generated passwords are:")
    } else {
    	return
    }        

    rand.Seed( time.Now().UTC().UnixNano())

    separators := strings.Split(*separators_string, "")
    words_bytes, err := ioutil.ReadFile("/usr/share/dict/words")
    if err != nil {
    	panic(err)
    }
    words := strings.Split(string(words_bytes), "\n")
    
    out_string := ""
    for j := 0; j < *num_passwords; j++ {
        out_string = ""
        for i := 0; i < *num_words; i++ {
            out_string += words[rand.Intn(len(words))]
            if i < *num_words -1 {
	        out_string += separators[rand.Intn(len(separators))]
    	    }
        }
        
        fmt.Println(out_string)
    }
}
