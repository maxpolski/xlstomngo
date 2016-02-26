package main

import (
  "fmt"
  "github.com/tealeg/xlsx"
)

func main() {
  excelFileName := "/home/maxpolski/Рабочий стол/Report.xls"

  xlFile, err := xlsx.OpenFile(excelFileName)
  if err != nil {
    panic(err)
  }
  for _, sheet := range xlFile.Sheets {
    // fmt.Printf("%s\n", sheet.Name)
    if sheet.Name == "pairs isbn -> ASIN" || sheet.Name == "пары ISBN13 -> ASIN" {
      for _, row := range sheet.Rows {
        for _, cell := range row.Cells {
          if cell.String() != "" {
            // u, _ := strconv.ParseUint(cell.Value, 10)
            // strconv.FormatUint(u, 10)
            fmt.Printf("%s\n", cell.String())
          }
        }
      }
    }
  }
}
