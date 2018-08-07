
function createDataCellDef(results, columnWidths) {
    var datainfo = [];
    var i = 0;
    for(var prop in results) {
        if (i === columnWidths.length) break;
        datainfo.push({value: results[prop], widthPercent: columnWidths[i++], cssClassName: 'mydatacell'})
    }
    return datainfo;
}

function createHeaderDef(headers, columnWidths) {
    var headerInfo = [];

    for(var i=0; i < headers.length; i++) {
        headerInfo.push({value:headers[i], widthPercent: columnWidths[i], cssClassName:'myheadercell'})
    }

    return headerInfo;
}

function searchBooks(app) {
    app.$http.get('/searchbooks?searchInput=' + app.searchInput).then((response) => {
        app.searchOutput = JSON.parse(response.bodyText);
        app.tableHeaderInfo = createHeaderDef(app.searchOutput.Headers, app.searchOutput.ColumnWidths);
        for(let ind in app.searchOutput.Results) {
            app.searchOutput.Results[ind]['celldef'] = createDataCellDef(app.searchOutput.Results[ind], app.searchOutput.ColumnWidths);
        }
    });
}

function getEarnings(app) {
    app.$http.get('/getEarnings?symbol=' + app.searchInput).then((response) => {
        app.searchOutput = JSON.parse(response.bodyText);
        app.tableHeaderInfo = createHeaderDef(app.searchOutput.Headers, app.searchOutput.ColumnWidths);
        Results = []
        app.searchOutput['Results'] = Results;
        for(let ind in app.searchOutput.EarningsData.earnings) {
            app.searchOutput.Results.push(app.searchOutput.EarningsData.earnings[ind]);
            app.searchOutput.Results[ind]['celldef'] = createDataCellDef(app.searchOutput.EarningsData.earnings[ind], app.searchOutput.ColumnWidths);
        }
    });
}
