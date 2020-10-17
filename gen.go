// +build ignore

package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strconv"
  "text/template"
  "time"

  "github.com/MarkusFreitag/advent-of-code/util"
)

var (
  year = time.Now().Year()
  day = time.Now().Day()
)

func yearPkg() string { return fmt.Sprintf("year%d", year) }

func dayPkg() string { return fmt.Sprintf("day%d", day) }

func main() {
  if os.Args[1] != "" {
    num, err := strconv.Atoi(os.Args[1])
    if err != nil {
      fmt.Printf("YEAR must be a number: %s\n", err.Error())
      return
    }
    year = num
  }
  if os.Args[2] != "" {
    num, err := strconv.Atoi(os.Args[2])
    if err != nil {
      fmt.Printf("DAY must be a number: %s\n", err.Error())
      return
    }
    day = num
  }

  templates, err := template.ParseGlob("templates/*.tpl")
  if err != nil {
    fmt.Printf("error while loading templates: %s\n", err.Error())
    return
  }

  path := filepath.Join("puzzles", yearPkg(), dayPkg())
  // check whether this day has already been generated
  if _, err := os.Stat(path); !os.IsNotExist(err) {
    fmt.Printf("%d_%d already exists\n", year, day)
    return
  }

  /*
    1. download the days puzzle input and save it at inputs/YEAR_DAY.txt
    2. ensure the days package directory exists
    3. generate the code and test files at puzzles/YEAR/DAY
    4. regenerate the years code file
    5. regenerate the puzzles code file
  */

  input, err := util.InputFromURL(year, day)
  if err != nil {
    fmt.Printf("error while downloading puzzle input for %d:%d: %s\n", year, day, err.Error())
    return
  }
  if err := os.MkdirAll("inputs", os.ModePerm); err != nil {
    fmt.Printf("error while creating input directory: %s\n", err.Error())
    return
  }
  if err := util.InputToFile(year, day, input); err != nil {
    fmt.Printf("error while saving puzzle input into file: %s\n", err.Error())
    return
  }

  if err := os.MkdirAll(path, os.ModePerm); err != nil {
    fmt.Printf("error while creating package directories: %s\n", err.Error())
    return
  }


  dayData := struct{
    Package string
  }{
    Package: dayPkg(),
  }
  fd, err := os.OpenFile(filepath.Join(path, fmt.Sprintf("%s.go", dayPkg())), os.O_RDWR|os.O_CREATE, 0664)
  if err != nil {
    fmt.Printf("error while creating code file: %s\n", err.Error())
    return
  }
  if err := templates.ExecuteTemplate(fd, "day.tpl", dayData); err != nil {
    fmt.Printf("error while generating code: %s\n", err.Error())
    return
  }
  if err := fd.Close(); err != nil {
    fmt.Printf("error while saving code file: %s\n", err.Error())
    return
  }

  fd, err = os.OpenFile(filepath.Join(path, fmt.Sprintf("%s_test.go", dayPkg())), os.O_RDWR|os.O_CREATE, 0664)
  if err != nil {
    fmt.Printf("error while creating testcode file: %s\n", err.Error())
    return
  }
  if err := templates.ExecuteTemplate(fd, "day_test.tpl", dayData); err != nil {
    fmt.Printf("error while generating testcode: %s\n", err.Error())
    return
  }
  if err := fd.Close(); err != nil {
    fmt.Printf("error while saving testcode file: %s\n", err.Error())
    return
  }


  yearData := struct{
    Package string
    Imports []string
    Days    []string
  }{
    Package: yearPkg(),
    Imports: make([]string, 0),
    Days:    make([]string, 0),
  }
  entries, err := ioutil.ReadDir(filepath.Join("puzzles", yearPkg()))
  if err != nil {
    fmt.Print("err: %s\n", err.Error())
    return
  }
  for _, entry := range entries {
    if entry.IsDir() {
      yearData.Imports = append(
        yearData.Imports,
        fmt.Sprintf("github.com/MarkusFreitag/advent-of-code/puzzles/%s/%s", yearPkg(), entry.Name()),
      )
      yearData.Days = append(yearData.Days, entry.Name())
    }
  }
  fd, err = os.OpenFile(filepath.Join("puzzles", yearPkg(), fmt.Sprintf("%s.go", yearPkg())), os.O_RDWR|os.O_CREATE, 0664)
  if err != nil {
    fmt.Printf("error while creating code file: %s\n", err.Error())
    return
  }
  if err := templates.ExecuteTemplate(fd, "year.tpl", yearData); err != nil {
    fmt.Printf("error while generating code: %s\n", err.Error())
    return
  }
  if err := fd.Close(); err != nil {
    fmt.Printf("error while saving code file: %s\n", err.Error())
    return
  }


  puzzlesData := struct{
    Years []string
  }{
    Years: make([]string, 0),
  }
  entries, err = ioutil.ReadDir("puzzles")
  if err != nil {
    fmt.Print("err: %s\n", err.Error())
    return
  }
  for _, entry := range entries {
    if entry.IsDir() {
      puzzlesData.Years = append(puzzlesData.Years, entry.Name())
    }
  }
  fd, err = os.OpenFile(filepath.Join("puzzles", "puzzles.go"), os.O_RDWR|os.O_CREATE, 0664)
  if err != nil {
    fmt.Printf("error while creating code file: %s\n", err.Error())
    return
  }
  if err := templates.ExecuteTemplate(fd, "puzzles.tpl", puzzlesData); err != nil {
    fmt.Printf("error while generating code: %s\n", err.Error())
    return
  }
  if err := fd.Close(); err != nil {
    fmt.Printf("error while saving code file: %s\n", err.Error())
    return
  }
}
