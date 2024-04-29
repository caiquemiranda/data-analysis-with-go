package main

import (
    "encoding/csv"
    "fmt"
    "io/ioutil"
    "strconv"
)

func main() {
    // Definir o nome do arquivo CSV
    fileName := "datasets/dados.csv"

    // Abrir o arquivo CSV
    file, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo CSV:", err)
        return
    }
    defer file.Close()

    // Ler o conteúdo do arquivo CSV
    reader := csv.NewReader(bytes.NewReader(file))
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Erro ao ler o arquivo CSV:", err)
        return
    }

    // Definir a coluna para análise (começa em 0)
    coluna := 1

    // Inicializar variáveis para soma e contagem
    soma := 0.0
    contagem := 0

    // Calcular a soma dos valores na coluna especificada
    for _, record := range records {
        valor, err := strconv.ParseFloat(record[coluna], 64)
        if err != nil {
            fmt.Println("Erro ao converter valor:", err)
            continue
        }

        soma += valor
        contagem++
    }

    // Calcular a média
    if contagem > 0 {
        media := soma / float64(contagem)
        fmt.Printf("Média da coluna %d: %.2f\n", coluna+1, media)
    } else {
        fmt.Println("Nenhum dado encontrado na coluna especificada.")
    }

    // Ordenar os registros por valor na coluna especificada
    sort.Slice(records, func(i, j int) bool {
        valorI, _ := strconv.ParseFloat(records[i][coluna], 64)
        valorJ, _ := strconv.ParseFloat(records[j][coluna], 64)
        return valorI < valorJ
    })

    // Calcular a mediana (se houver um número ímpar de registros)
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
