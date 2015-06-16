package main

import (
  "os"
  "fmt"
)

func dirSize(path string, info os.FileInfo) int64 {
  size := info.Size()
   if !info.IsDir() {
     return size
   }
   dir, err := os.Open(path)

   if err != nil {
     fmt.Println(err)
     return size
   }
   defer dir.Close()

   files, err := dir.Readdir(0)

   if err != nil {
     fmt.Println(err)
     os.Exit(1)
   }

   for _, file := range files {
     if file.Name() == "." || file.Name() == ".." {
      continue
     }
     size += dirSize(path+"/"+file.Name(), file)
   }
   return size
 }

 func main() {
   if len(os.Args) != 2 {
     fmt.Printf("USAGE : %s <target_directory> \n", os.Args[0])
     os.Exit(0)
   }

   dir := os.Args[1] // get the target directory

   info, err := os.Lstat(dir)

   if err != nil {
     fmt.Println(err)
     os.Exit(1)
   }
   
   total := float64(dirSize(dir, info))

   if total >= 1000000000 {
    fmt.Printf("This directory is: [%f] GB : [%s]\n", total / 1000000000, dir)
   }
   if total >= 1000000 && total < 1000000000 {
    fmt.Printf("This directory is: [%f] MB : [%s]\n", total / 1000000, dir)
   }
   if total >= 1000 && total < 1000000 {
    fmt.Printf("This directory is: [%f] KB : [%s]\n", total / 1000000, dir)
   }
 }