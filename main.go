package main

import (
  "fmt"
  "github.com/tealeg/xlsx"
  // "reflect"
)

func main() {
  excelFileName := "/home/maxpolski/Рабочий стол/xls.xlsx"

  xlFile, err := xlsx.OpenFile(excelFileName)
  if err != nil {
    panic(err)
  }
  for _, sheet := range xlFile.Sheets {
    // fmt.Printf("%s\n", sheet.Name)
    for _, row := range sheet.Rows {
      for cellNum, cell := range row.Cells {
        if cell.Value != "" {
          fmt.Printf("Row %d value %s\n", cellNum, cell.NumFmt)
        }
      }
    }
  }
}
