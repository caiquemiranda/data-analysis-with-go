package main

import (
    "encoding/csv"
    "fmt"
    "io/ioutil"
    "strconv"
)

func main() {
    file := "datasets/dados.csv"

    file, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo CSV:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(bytes.NewReader(file))
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Erro ao ler o arquivo CSV:", err)
        return
    }

    coluna := 1
    soma := 0.0
    contagem := 0

    for _, record := range records {
        valor, err := strconv.ParseFloat(record[coluna], 64)
        if err != nil {
            fmt.Println("Erro ao converter valor:", err)
            continue
        }

        soma += valor
        contagem++
    }

    if contagem > 0 {
        media := soma / float64(contagem)
        fmt.Printf("Média da coluna %d: %.2f\n", coluna+1, media)
    } else {
        fmt.Println("Nenhum dado encontrado na coluna especificada.")
    }

    sort.Slice(records, func(i, j int) bool {
        valorI, _ := strconv.ParseFloat(records[i][coluna], 64)
        valorJ, _ := strconv.ParseFloat(records[j][coluna], 64)
        return valorI < valorJ
    })

    if contagem%2 == 1 {
        medianaIndex := contagem / 2
        medianaValor, _ := strconv.ParseFloat(records[medianaIndex][coluna], 64)
        fmt.Printf("Mediana da coluna %d: %.2f\n", coluna+1, medianaValor)
    } else {
        medianaIndex1 := (contagem - 1) / 2
        medianaIndex2 := medianaIndex1 + 1
        medianaValor1, _ := strconv.ParseFloat(records[medianaIndex1][coluna], 64)
        medianaValor2, _ := strconv.ParseFloat(records[medianaIndex2][coluna], 64)
        mediana := (medianaValor1 + medianaValor2) / 2.0
        fmt.Printf("Mediana da coluna %d: %.2f\n", coluna+1, mediana)
    }
}
