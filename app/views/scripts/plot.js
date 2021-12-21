function drawResult(datetime, observed, predicted) {
    observed = replaceWithA(observed)
    predicted = replaceWithA(predicted)

    var observedTrace = {
        x: datetime,
        y: observed,
        type: 'scatter',
        name: "observed",
    };

    var predictedTrace = {
        x: datetime,
        y: predicted,
        type: 'scatter',
        name: "predicted",
    };

    var data = [observedTrace, predictedTrace];

    var layout = {
        title: "title",
        xaxis: {
            type: 'date'
        },
        margin: {
            t: 100,
            b: 100,
        }
    };

    return Plotly.newPlot("draw", data, layout);
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