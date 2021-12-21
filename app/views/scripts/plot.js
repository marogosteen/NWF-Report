function drawResult(drawId, observed, predicted) {
    observed = replaceWithA(observed)
    predicted = replaceWithA(predicted)

    var observedTrace = {
        x: Array.from(Array(observed.length), (v, k) => k),
        y: observed,
        type: 'scatter',
        name: "observed",
    };

    var predictedTrace = {
        x: Array.from(Array(predicted.length), (v, k) => k),
        y: predicted,
        type: 'scatter',
        name: "predicted",
    };

    var data = [observedTrace, predictedTrace];

    var layout = {
        title: "title",
        margin: {
            t: 100,
            b: 100,
        }
    };

    return Plotly.newPlot(drawId, data, layout);
}


function replaceWithA(dim) {
    index = 0
    dim.forEach(function (element) {
        if (element == 0) {
            dim[index] = null
        }
        index += 1
    })
    return dim
}