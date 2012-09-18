package main

import (
    "fmt"
)

func createTrainingDataElement(input, output []float64) TrainingDataElement {
    return TrainingDataElement{input: input, output: output}
}

func createTestDataElement(input []float64) DataElement {
    return DataElement{input: input}
}

func main() {
    trainingdata := make([]TrainingDataElement, 4)
    trainingdata[0] = createTrainingDataElement([]float64{0.0, 0.0}, []float64{0.0})
    trainingdata[1] = createTrainingDataElement([]float64{0.0, 1.0}, []float64{1.0})
    trainingdata[2] = createTrainingDataElement([]float64{1.0, 0.0}, []float64{1.0})
    trainingdata[3] = createTrainingDataElement([]float64{1.0, 1.0}, []float64{0.0})

    testdata := make([]DataElement, 4)
    testdata[0] = createTestDataElement([]float64{0.0, 0.0})
    testdata[1] = createTestDataElement([]float64{0.0, 1.0})
    testdata[2] = createTestDataElement([]float64{1.0, 0.0})
    testdata[3] = createTestDataElement([]float64{1.0, 1.0})

    network := Create(2, 2, 1)

    network.Train(trainingdata, 1000, 0.5, 0.1)
    result := network.Test(testdata)

    fmt.Println(result)
}
