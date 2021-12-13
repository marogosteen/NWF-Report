function drawResult(drawId, observed, predicted) {
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
        title: "title"
    };

    return Plotly.newPlot(drawId, data, layout);
}

